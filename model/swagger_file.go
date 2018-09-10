// Code generated by go-bindata.
// sources:
// model/model.swagger.json
// DO NOT EDIT!

package model

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _modelModelSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xdd\x6f\xdb\x38\x12\x7f\xcf\x5f\x31\xd0\xdd\xc3\x2e\xd0\x36\xbd\xde\xe1\x1e\xfa\x96\xdd\xa6\x1f\xc0\x2d\xb6\xd7\xe4\xee\xe5\xb6\x30\x68\x72\x6c\x71\x43\x91\x0a\x49\xd9\x6b\x14\xf9\xdf\x0f\x24\xf5\x6d\x49\x96\x3f\x62\x3b\xdb\x2c\xb0\x68\x42\x71\x28\x72\xe6\x37\xc3\xf9\x52\xbe\x5d\x00\x44\x66\x49\xe6\x73\xd4\xd1\x5b\x88\xde\xbc\x7a\x1d\xbd\x70\x63\x5c\xce\x54\xf4\x16\xdc\x73\x80\xc8\x72\x2b\xd0\x3d\x4f\x14\x43\xf1\x2a\xd5\xca\x2a\x3f\x0f\x20\x5a\xa0\x36\x5c\x49\xf7\x34\xff\x11\xa4\xb2\x60\xd0\x46\x17\x00\x0f\x7e\x35\x43\x63\x4c\xd0\x44\x6f\xe1\x7f\x81\x28\xb6\x36\x2d\x16\x70\x3f\x1b\x37\xf7\xab\x9f\x4b\x95\x34\x59\x63\x32\x49\x53\xc1\x29\xb1\x5c\xc9\xcb\xdf\x8d\x92\xd5\xdc\x54\x2b\x96\xd1\x91\x73\x89\x8d\x4d\x75\xa4\xcb\xc5\xdf\x2e\x67\x42\x2d\xab\x21\x37\x47\x19\x5b\xfb\x1d\x20\x52\x29\x6a\xbf\xda\x27\xe6\x4e\xf8\xb3\x46\x62\xf1\x83\x26\x69\x9c\x6f\xdf\xcf\xd2\x68\x52\x25\x0d\x9a\x06\x31\x40\xf4\xe6\xf5\xeb\xd6\x10\x40\xc4\xd0\x50\xcd\x53\x9b\x73\xad\xb6\x90\x7f\xec\x99\x45\xd6\xc8\x00\xa2\xbf\x6a\x9c\x39\x8a\xbf\x5c\x32\x9c\x71\xc9\xdd\x0a\xe6\xd2\xcb\xa4\xb6\xb1\x2f\xf9\x66\xa2\x06\xf9\xc3\x45\xd7\xcf\x0f\xb5\x43\xa4\x44\x93\x04\x2d\xea\x8a\x9d\xe1\xbf\xd6\xf6\x25\x49\x3c\x16\xa6\x8a\xad\xda\x7b\xe7\xb2\xef\x89\xc6\xfb\x8c\x6b\x74\x5c\xb4\x3a\xc3\x83\x9f\xf9\x3e\x43\x63\xc7\x1c\xf9\x6b\xed\xc8\x96\xcc\xdb\x87\x8d\xde\x0b\xb5\xbc\x41\xbd\xe0\xb4\xc6\xc1\xaf\x17\xf5\x65\x72\xae\x55\x18\xba\xfc\xe6\xfe\x99\x70\xf6\x50\x47\xd3\x1c\x87\xc1\xf4\x01\xad\xdf\xfc\x8d\x25\x16\xcf\x0b\x4e\x8d\xad\x1d\x0b\x50\x39\x0f\xbb\x31\xe5\xd4\x77\x3b\x4c\xd9\x55\xea\x97\x35\x56\x73\x39\x8f\x4e\x00\x86\x4b\xaa\x92\x84\xdb\xed\x2c\x4c\x20\x39\x2f\x34\xd4\x54\xec\xb3\x56\x14\x8d\x41\xf6\x8c\x8a\x5d\x51\xc1\x50\x90\xd5\x56\xa0\xb8\x62\xec\x9d\x27\x3a\x2b\x58\x5c\x31\x76\x63\xc9\xfc\x4f\x68\x1f\x5e\x3c\x81\xdb\xaf\xc0\x44\x2e\x82\x33\xb9\xff\x2e\xb9\x5c\xa8\x3b\xdc\x16\xdd\x9f\x3c\xd5\xfb\x4c\x52\x8f\xd0\x67\x98\x3f\xc3\xbc\xe2\x7e\x13\x1c\x67\x86\x77\xe3\xb6\xb3\x2d\xdc\xfd\x19\x9e\x51\xfe\x8c\xf2\x75\xee\x9f\x15\xae\xcd\xe5\x37\xff\xaf\x1f\x21\x4b\xd2\xf4\x65\x37\xc5\x37\x57\x8e\xa0\x00\x55\x26\xce\xcc\xa9\x6d\xef\xee\xbb\x04\x7e\x21\xdd\xf3\xd8\x8d\xe5\x09\xaa\xcc\x4e\x12\xd3\xbd\x9f\xfb\x0c\xf5\x90\x36\xce\x88\x30\x7d\x3b\xe2\xd2\xe2\x1c\x75\x9b\x7a\xa6\x74\x42\x6c\x3e\xe1\xef\x6f\x4e\x12\x0f\xac\xab\x1a\x55\x49\x2a\xd0\x6e\x77\xaf\xfc\x9c\x13\x79\x4c\x5f\xff\x61\x51\x4b\x22\xc4\x99\xc5\x0c\x3d\x9b\x7c\x56\xbe\x93\xef\xe6\xe4\x77\x60\x2f\x34\xce\xe7\x4a\xd4\x48\x92\x6d\x6e\xc0\x1b\x4f\x71\xbd\x40\x69\xcd\xc1\xf4\xf0\x87\xb0\x0f\x2e\xe7\x50\xae\xf3\xe3\xe1\xf2\x3c\xb5\x3d\x7f\x57\xda\x38\xd3\x2a\x99\x18\xbc\x3f\xfc\xd5\x93\xef\xa7\xf7\xe6\xc9\xb8\xb4\xff\xfc\xc7\x49\xae\x9e\x05\x11\xd9\xd6\xd1\xcb\x7f\x1d\xd1\x73\x08\x73\x5e\xf0\x3d\xb9\xf9\xbe\x62\xac\xb0\xe0\x35\x84\x9c\xce\x7a\xef\x6a\xad\xff\xc5\x67\x48\x57\x54\x1c\x0e\xdc\x8f\x6c\xb0\xcb\x0d\xef\x64\xb3\x77\xe5\x74\x59\xdb\xad\x6d\xaa\xaa\xae\x76\xd6\x8f\x72\x40\xa4\x1a\x0d\x4a\x4b\x72\xf6\x94\xe2\x29\xf0\xaf\xa6\xbf\x23\xad\x82\xc5\x28\xd5\x4e\x4e\x96\xb7\x98\x5f\xcc\x6f\x88\xa3\x4f\x87\xea\x27\x36\x96\xd8\x6c\x4d\x90\x63\x28\x19\xa6\x28\x19\x4a\xda\xde\x4b\x8d\x9e\x68\x4d\x9a\xea\x16\x71\x8b\x49\x7b\xfe\xc8\x4a\xc4\x43\x27\xbc\x47\xe8\xdb\x1e\x7c\x2d\xcc\xda\x0e\x0c\x6a\x5f\x27\x30\xce\xe5\xe3\x4a\xe6\x19\x82\xce\x55\xa9\x62\x38\x11\x8a\xb6\x11\x33\x7a\x5b\x94\x08\x81\x7a\x9b\x33\x5d\xb4\x16\xa9\x3a\x20\x06\xd9\x0e\xd4\x97\xa4\x0d\x10\x90\xb8\x04\xef\xe0\xc3\x92\xdb\x18\x08\x98\x14\x29\x9f\x71\x0a\x81\x49\x9d\x02\x5d\x4f\xe8\x9f\x46\x8e\xbe\x42\x35\x59\x03\xed\x90\x47\xd3\x8c\xa4\xeb\xee\xcc\x39\x8b\x72\x8d\xe1\x35\x09\x7a\x26\x04\x19\x76\x4b\x6b\x20\x2f\x7d\x1a\xb1\xcd\xf2\xad\xec\x48\x4e\xf4\x7c\x0b\xdd\xfd\x78\x7b\xfb\xf9\x0b\xde\xbf\x23\x36\x4b\x9e\x86\xb0\xfb\xe5\x05\x84\x31\x27\xf2\x82\x81\xc0\xe5\xa2\xd8\x77\xa7\xe8\xcf\x40\xd8\xa5\xf3\xb2\x93\xbd\xfd\xb5\xa4\xee\xe6\xbc\x50\x26\xd3\xdb\x98\xf2\x9f\x84\x9a\x0e\x60\x81\x61\x7a\xac\x2b\xf3\x8c\x11\xd8\x85\xb9\xd6\x45\x51\xf8\x18\xc1\xcb\x00\xab\xc0\xc6\x08\x73\xdf\x71\x36\x04\xc5\x3c\xe4\x39\x0d\x16\xcb\x4c\xd6\x21\xb9\x14\x4e\x04\x1a\x6d\xa6\xa5\xf1\x6c\x08\x7c\xfa\xf4\x0e\xd4\xcc\xff\x1e\x8c\x35\xeb\x37\xd2\x7d\x39\xfe\x27\xc2\xa6\x66\x04\xe2\x1c\xa4\x83\xf8\x56\xfd\x4e\x65\xa5\xc5\x7b\x70\x68\x2a\xd4\x74\xc7\x53\x52\x25\x2d\x4a\x3b\xd9\xd5\xb7\x17\x28\xe7\x36\x3e\x8c\xdb\xd2\x8f\xd1\x92\x4b\x10\x2b\xe1\x95\x58\xe3\x0c\x35\x4a\x8a\x4e\x61\x09\x38\x0e\x78\x80\x12\x63\x14\xe5\x1e\xa3\x7e\xcc\x58\xa5\xbb\x80\xba\x21\xf7\xf9\xf4\xe0\xfa\x1d\x44\x02\xc3\x32\x83\x84\xe8\x3b\x03\x44\x02\xfe\xc1\x8d\x75\xc1\x7f\xb0\x5e\xc4\x40\x51\xde\x19\x1d\x1a\x6c\xaa\x9a\x3c\x3d\x7c\x98\x8c\x52\x34\x66\x96\x89\x3e\xea\xa9\x52\x02\x89\xec\xd3\xd6\xe2\xf1\x18\xab\xd6\xe5\xf6\x74\xb0\xac\x65\x1e\x22\x94\xde\x0e\x56\xa9\x91\x28\x93\x77\x52\x2d\xe5\xa4\xf2\xbd\xea\xee\x33\xa5\x98\xda\x6b\x6e\xe3\x46\x99\xd1\xb7\xa4\xaf\x6e\xd5\xfa\x03\x1b\xa3\xbc\xf2\x44\x3f\xa9\x46\xb2\x30\x3c\x71\x64\xed\xc1\x2f\x99\xec\x5e\xa1\x3d\xea\x8e\xac\x0c\x76\x0c\x4f\xb9\x6c\x0c\x2f\xf3\xd9\xbe\xde\x58\x1b\x8f\x89\x64\xcd\x24\x98\xc9\xda\x5b\xe2\xbd\x6d\x5b\xb4\x11\x23\x47\xed\xd8\xb2\xc1\x1f\x21\x7e\x9d\x35\x06\xe4\xaa\x39\x80\x39\xe2\x2b\x39\x36\x9f\x3a\x06\x70\xb5\x56\xed\xb4\xa8\x13\x2e\xbd\x9c\x3e\x2a\x75\xd7\x4b\x53\xf0\x2a\x7f\x5c\xe6\x21\x23\x86\x33\x12\x2e\xdd\x01\xc9\xaf\x19\x84\x3a\xcc\x20\xa4\x02\xa7\x18\xfc\x17\x87\x34\x20\x92\xc1\x14\x63\xb2\xe0\x2a\xd3\xce\x93\x21\xb9\x65\xc8\x9d\x9a\x3e\x5f\x6f\xcd\x44\xee\xa1\xf5\x8f\xa0\x7e\x75\x11\xb7\x3c\x08\xd8\x60\xfc\xdb\x71\xc3\x46\x9b\xcb\x65\xc1\x84\xf2\x06\xf6\xa6\x13\xc2\xad\xec\x38\xec\xee\x64\x3f\x21\xa4\xfb\xe0\x87\xea\xc8\x97\x33\xc2\x05\xb2\x1f\xbb\x78\xbc\xfe\x49\xc1\x3e\xb6\x75\xbf\x78\x7c\x5b\xd3\x3c\x60\x00\x3b\xbe\x0e\x39\xe2\x9d\xd1\xbf\xb1\xbd\x7d\x4d\x4c\x52\xbb\xda\x02\x6a\xd7\x6e\xfe\x40\x9c\xea\xbc\xb4\xc3\x45\xbd\xa8\xb5\xd2\xdb\xec\xce\xcd\x1f\x58\x2f\xdc\xb8\x81\x7e\xec\x9a\x79\xf8\x33\x1b\x58\x36\xb6\x36\x9d\x68\xbc\x3f\x68\xf6\x27\x5f\xd4\xa4\x5b\xaf\x6a\xd2\x61\x16\x74\x64\xef\x87\xcf\x6f\x33\x33\xde\xc4\x04\x13\x82\x92\x92\x34\x33\x99\x08\xc9\x40\x21\x20\x55\xc6\xf0\xa9\xc0\x60\x69\x9c\x3d\x27\xce\xc9\x5b\xc1\xb4\xe1\xe6\x17\x2e\x9d\xb7\xe8\x5c\x7a\x8b\xee\xf4\xa5\xd7\xac\x87\x14\x24\x8d\x91\x65\x02\x59\xa8\xe1\xec\x63\xd9\xf7\x70\xca\x2c\x4f\xf0\x31\xb2\xbe\x76\xc7\x25\x19\xb1\xf8\xd2\x6d\xea\xb0\xf6\xb1\x23\xe9\x90\x27\x7b\x63\x62\x9c\xe4\xb4\x13\xe4\x4b\xb0\x31\x37\xb9\x1b\xef\xa4\xa8\x51\x10\xcb\x17\x18\x66\xf8\x8b\x5b\x02\x3a\x79\x81\xf3\xa1\xca\x8c\xb1\x54\x0c\x81\x1b\xd0\x48\xd5\x02\x35\xb2\x0e\xa1\xd7\xcc\xd0\x08\x59\x97\xfb\xac\xc8\x40\x17\x75\x34\x9f\x43\xca\x84\x00\xa5\xc1\x5b\xc3\xde\x28\xa2\x66\x5d\x0e\x5c\x76\x1b\x65\xd2\x6e\x1d\x61\xa7\x20\x13\x34\xa6\xd9\x45\xbd\xab\x20\xab\xb7\x35\x18\x24\xc1\x5b\xe2\xa0\xb3\x4e\xc8\x8a\xd2\x4c\xe7\xda\xca\x9d\xe8\x2a\x15\x7d\x51\x79\x6a\x5c\x32\x4e\xbd\x01\xf0\xc9\x26\x92\x99\xd2\x49\x0b\xeb\x39\x47\xc3\xfd\x96\x1f\xc0\x49\x9d\xcb\x80\x5f\xef\x5b\x0e\x0a\xe1\xb6\xc9\xc8\x1d\xc2\x90\x70\xbd\xac\xdd\x0f\x79\x1b\xe5\xfa\x83\xe0\xf7\x44\x5d\x15\x83\x0e\x9a\xca\x7b\xf1\x5e\x7e\x07\x75\x58\x55\x28\x63\x5b\x41\x01\x11\x9c\x4d\x8a\xcb\xaa\xd9\x8c\x31\xe8\x5d\x37\x0f\xd4\x25\x56\xc7\x34\xf0\x20\x43\x2f\x57\x2e\x43\x68\x90\x0b\x84\x3a\xdd\x9b\xa3\x74\xee\x77\x5b\xbc\x1d\xc2\x78\x4f\x88\xf9\x54\xa6\xfa\xcb\xba\xde\x49\x8d\xf0\x21\x13\x7d\x8f\x6f\x83\x29\x11\xe2\xc8\xee\x6d\x97\xf9\x76\xdb\x28\x72\xe5\xef\x09\xb9\x29\x53\x2c\x5d\xc6\xb7\x29\xf5\x9b\x60\xef\x4f\x7b\xf1\x3e\xd2\x05\x79\x9a\xf8\x63\x94\x80\x6a\x57\x6d\x87\x88\xba\xbf\x3f\xde\x57\x3a\xbd\x5c\x6e\xad\xe2\x9f\x11\xc6\x78\xb0\xe2\x9f\xbb\xd7\x84\x7e\xbd\x1c\xdd\xfe\x52\x5b\xac\xbb\x82\x75\x36\x21\xa4\x3f\x4f\xf8\x38\xf9\x10\xda\xb2\x47\xda\xf1\x90\xca\xb2\x06\xd9\x56\xeb\xd6\x6d\x91\x8f\x71\x57\x3b\x2d\x4e\x5f\xf8\x86\x76\x95\x72\x07\xec\x55\xc3\x53\x20\x36\x38\xfd\x45\x2d\xd9\x8f\x14\x95\xaa\x32\x1c\x88\x6b\x79\x60\xf6\x9b\x54\xbe\x76\xe0\x16\x0d\x1e\xa5\x7b\x3c\x45\x94\x2e\xec\x28\xe8\xc2\x46\x96\x5c\x08\x70\x50\x33\x71\xee\x77\x0a\x01\x84\x7a\xcf\x54\x69\x48\x51\x32\x2e\xe7\xde\x0b\x35\x10\x93\x05\x56\xaf\x79\xd5\xa5\x69\xb9\x5c\x0f\x74\xf3\x9d\xae\x7b\xe1\x51\x61\x51\x5a\xb2\x0f\x35\x30\x84\x1c\xbe\x77\x02\xbd\x50\xa4\x02\xa1\xe4\x1c\x35\xd0\x98\xc8\xce\x82\x64\xe0\x76\xc0\xc2\x33\xaf\x47\xf1\x3a\xd7\x9c\x3e\x66\xb6\xba\x1f\xf7\xb9\x22\xd6\x13\x20\xe3\x4e\xd5\xee\xdf\xde\xc7\x00\xd7\x48\xbd\xba\x4f\x8a\xe3\x8f\xf7\x07\xd7\x31\x36\xb4\x7c\xe9\x27\x6d\xfb\x82\xa6\xc9\x18\x7d\x85\xf4\xfe\x19\x89\xb3\xc8\x47\xae\x7d\x01\xf1\x27\xc2\x93\xf1\x85\xf5\xc7\xf1\x34\x8f\x82\xd5\xb2\x9e\x23\xb7\xe9\x53\xf3\xaf\xb8\xad\x48\x4f\xa0\x12\xed\xea\xd7\xc4\x14\xe9\xbe\x6d\x8a\x24\x1d\x79\xc2\x6e\x39\xfb\x40\x84\x30\xb6\xd5\xf2\xde\x3b\xbd\x72\x44\x1b\x97\xde\x85\x41\x7e\xf9\x31\x0c\xaa\x5e\xa1\xcc\x6e\x6f\x70\x74\x03\x2f\x98\x11\x62\x26\x55\x7f\xdf\xa4\x08\x42\xc6\xbf\x69\x20\x7e\x1c\xf5\xc6\x5d\xf8\x37\x98\xa9\x18\x06\x73\xf0\x59\x77\x00\x73\xdd\xd5\x1f\x6d\x3c\xd7\x34\xed\x49\x3a\x38\xfb\x54\x19\xfa\xf3\x9c\xc7\x74\x9b\x6a\xc5\xef\x56\x5a\x5b\x2a\x48\x94\xc6\xdc\x49\x35\x40\x89\x84\x29\x42\x42\x18\x86\x08\x9d\x9b\x10\x6a\xfc\x26\xfd\xcf\xde\xb7\x9d\x22\xcc\xb8\x10\x6a\x89\x0c\xa6\x2b\x20\x85\xff\xeb\x96\xaf\xe7\xc0\xe5\xaa\xf9\x66\xa5\xee\xf2\x28\x44\x67\x5d\xbd\xb4\x1f\x6f\x6f\x3f\x7f\x44\xc2\x50\xef\x03\x94\x3b\x5c\xab\x01\xee\xdc\xa4\xb4\x43\x7a\xa3\x3a\x04\x2c\x35\x49\x0d\x10\x30\x5c\xce\x05\x42\x1c\x46\xef\x70\x75\xd9\x97\x95\x77\xc4\xbf\xa0\x8d\x15\xdb\x2f\x19\x9c\x84\x35\xea\x16\x00\x1b\x79\x59\xb7\x97\xfa\xef\x69\x2b\x6f\x9b\x36\xd3\xbf\x0c\xdb\x1d\x21\x2a\x0d\x5f\xf9\xd4\x69\x88\xa5\xf1\xa8\xd4\x6e\x6b\x7b\x4d\xe6\x85\xf3\x57\x79\xdd\xaa\x23\xca\x3d\x85\x9c\xb6\x9b\x77\x65\x15\x72\x9f\x8e\x45\xc5\xb6\x29\x22\x6f\xa8\xfa\x06\xa1\x1f\xa0\xdd\x79\x43\xa5\x34\xd7\x9b\x1a\x4d\x77\x0e\x29\x69\xa3\x6b\xcc\xe2\xbf\xd4\x78\x0e\x9b\xe1\x5f\x48\xa1\xd6\xf0\x42\x24\xa8\xcc\xce\x15\x97\x73\x50\x1a\xb8\xa4\xca\x7f\x03\xe7\x45\x9a\x87\x01\x2f\x80\xdb\x5c\x67\x6c\x5c\xa8\x8b\x79\x91\x4b\xdc\x87\xd6\x2a\xef\xd3\x01\x2f\xa4\x3e\x0c\x14\x35\xe3\x67\x10\xf4\x80\x20\xdc\x68\x13\xaa\x58\xaf\xc5\xeb\xfa\x13\x19\xfd\x7f\x20\x63\x23\x20\x72\x91\x14\x26\x71\x10\x0d\xd7\xff\xfe\xcf\xf5\xcd\x6d\x1f\x1a\xf2\x16\x1e\x5f\xdc\x29\x20\x21\x55\x3f\x24\xda\xbe\xec\x89\xb2\xfa\x6a\x9b\xd6\x87\x93\x7e\x9a\x71\xec\xaf\x1a\x8f\x50\x97\x3a\x7e\xa7\xf0\x23\x55\x4e\xf2\x6f\x43\x88\x01\x1f\xd2\x6d\xfa\x22\xa4\x2b\xd0\x7a\xae\x64\x1e\xb9\x9b\x24\xc8\xac\x8c\xf3\x00\x5e\x02\x53\x4b\x19\xbe\x03\x0f\x4f\x4d\xd1\x3b\x64\x35\x9f\xcf\x7b\xda\x44\x3a\x62\xda\x13\xc9\xb2\x08\xc8\x27\xe7\x58\xe6\x7c\x14\x6d\x2b\x4e\x0c\x5c\x5a\x05\x50\x9c\xdb\xc9\xb2\xfc\xd9\x87\x46\x6e\x22\x56\xed\xfc\xfe\x3b\x25\x82\x49\xd1\x7f\x5a\x5b\xaa\x62\x5e\x8f\xa8\xcb\x36\xbd\x63\x4a\xb9\x9f\x1d\x8d\x3d\x35\x3b\x8e\x9a\x1f\xbc\x48\x65\x63\xd4\xcd\xa6\xb7\x01\x03\x55\xf6\xe3\x1d\xaf\x0d\xa9\x3f\x3e\x1f\x3c\xbe\xc5\x86\x0f\x03\xb5\xf2\x60\xd1\x09\xe4\x9c\x93\x66\xef\x36\xf8\x96\x13\x56\xdd\xb8\x7d\xe7\xdf\xbf\x03\xc8\xb4\xff\x74\xb8\xef\x72\x46\xd6\xea\xf2\x59\xeb\xdc\xa1\x44\x52\x14\xad\xc1\x3b\xee\x47\xf2\x81\xc1\x68\xae\xf9\xde\x2e\x96\xe5\x8d\x3a\xf5\xee\xf7\x92\x53\x65\xc3\xbb\x6e\xa0\xe4\xc2\xfd\xff\x70\xf1\xff\x00\x00\x00\xff\xff\xe0\x57\xaf\x8e\x94\x60\x00\x00")

func modelModelSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_modelModelSwaggerJson,
		"model/model.swagger.json",
	)
}

func modelModelSwaggerJson() (*asset, error) {
	bytes, err := modelModelSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "model/model.swagger.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"model/model.swagger.json": modelModelSwaggerJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"model": &bintree{nil, map[string]*bintree{
		"model.swagger.json": &bintree{modelModelSwaggerJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

