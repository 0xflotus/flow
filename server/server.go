package server

import (
	"errors"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/fnproject/flow/cluster"
	"github.com/fnproject/flow/model"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go-opentracing"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var log = logrus.WithField("logger", "server")

func (s *Server) handlePrometheusMetrics(c *gin.Context) {
	s.promHandler.ServeHTTP(c.Writer, c.Request)
}

func (s *Server) handleSwagger(c *gin.Context) {
	file, err := model.Asset("model/model.swagger.json")
	if err != nil {
		c.AbortWithError(500, errors.New("error reading swagger content"))
		return
	}

	c.Data(200, "application/json", file)

}

func setTracer(ownURL string, zipkinURL string) {
	var (
		debugMode          = false
		serviceName        = "flow-service"
		serviceHostPort    = ownURL
		zipkinHTTPEndpoint = zipkinURL
		// ex: "http://zipkin:9411/api/v1/spans"
	)

	var collector zipkintracer.Collector

	// custom Zipkin collector to send tracing spans to Prometheus
	promCollector, promErr := NewPrometheusCollector()
	if promErr != nil {
		logrus.WithError(promErr).Fatalln("couldn't start Prometheus trace collector")
	}

	logger := zipkintracer.LoggerFunc(func(i ...interface{}) error { logrus.Error(i...); return nil })

	if zipkinHTTPEndpoint != "" {
		// Custom PrometheusCollector and Zipkin HTTPCollector
		httpCollector, zipErr := zipkintracer.NewHTTPCollector(zipkinHTTPEndpoint, zipkintracer.HTTPLogger(logger))
		if zipErr != nil {
			logrus.WithError(zipErr).Fatalln("couldn't start Zipkin trace collector")
		}
		collector = zipkintracer.MultiCollector{httpCollector, promCollector}
	} else {
		// Custom PrometheusCollector only
		collector = promCollector
	}

	ziptracer, err := zipkintracer.NewTracer(zipkintracer.NewRecorder(collector, debugMode, serviceHostPort, serviceName),
		zipkintracer.ClientServerSameSpan(true),
		zipkintracer.TraceID128Bit(true),
	)
	if err != nil {
		logrus.WithError(err).Fatalln("couldn't start tracer")
	}

	// wrap the Zipkin tracer in a FnTracer which will also send spans to Prometheus
	fntracer := NewFnTracer(ziptracer)

	opentracing.SetGlobalTracer(fntracer)
	logrus.WithFields(logrus.Fields{"url": zipkinHTTPEndpoint}).Info("started tracer")
}

// Server is the flow service root
type Server struct {
	Engine         *gin.Engine
	listen         string
	grpcServer     *grpc.Server
	clusterManager *cluster.Manager
	promHandler    http.Handler
}

func serverGrpc(server *grpc.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		log.Infof(" proto: %s %d %s %s", r.Method, r.ProtoMajor, r.Header.Get("Content-type"), r.RequestURI)

		if r.ProtoMajor == 2 && (r.Method == "PRI" || strings.Contains(r.Header.Get("Content-Type"), "application/grpc")) {
			log.Info("Serving GRPC ")
			server.ServeHTTP(w, r)
		} else {
			c.Next()
		}
	}
}

// NewAPIServer creates a new API server - params are injected dependencies
func NewAPIServer(clusterManager *cluster.Manager, restListen string, grpcListen string, zipkinURL string) (*Server, error) {

	setTracer(restListen, zipkinURL)

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpcUnaryTimeout(2*time.Minute),
			unaryMetricsInterceptor(),
			grpc_validator.UnaryServerInterceptor())),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamMetricsInterceptor(),
			grpc_validator.StreamServerInterceptor())))
	proxySvc := cluster.NewClusterProxy(clusterManager)
	model.RegisterFlowServiceServer(gRPCServer, proxySvc)

	log.WithField("listen", grpcListen).Info("Starting API gRPC service")

	l, err := net.Listen("tcp", grpcListen)
	if err != nil {
		return nil, err
	}
	go func() {
		gRPCServer.Serve(l)
	}()

	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery(), engineMetrics(), serverGrpc(gRPCServer))

	log.WithField("listen", restListen).Info("Starting API REST service")

	s := &Server{
		Engine:         engine,
		listen:         restListen,
		grpcServer:     gRPCServer,
		clusterManager: clusterManager,
		promHandler:    promhttp.Handler(),
	}

	s.Engine.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	model.RegisterFlowServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:9999", []grpc.DialOption{grpc.WithInsecure()})

	s.Engine.Any("/v1/*path", func(c *gin.Context) {
		log.Debugf("Serving HTTP, request is %+v", c.Request)

		gwmux.ServeHTTP(c.Writer, c.Request)
	})

	s.Engine.GET("/metrics", s.handlePrometheusMetrics)
	s.Engine.GET("/swagger.json", s.handleSwagger)
	return s, nil

}

// Run starts the server
func (s *Server) Run() error {
	srv := &http.Server{
		Addr:           s.listen,
		Handler:        s.Engine,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return srv.ListenAndServe()
}

func grpcUnaryTimeout(max time.Duration) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, cancel := context.WithTimeout(ctx, max)
		defer cancel()
		// log.Debugf("Original context: %+v", ctx, "; new context: %+v", newCtx)
		return handler(newCtx, req)
	}
}
