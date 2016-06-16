// Code generated by go-bindata.
// sources:
// root-fs/bower.json
// root-fs/favicon.ico
// root-fs/fcs-lpc-motor.html
// root-fs/index.html
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
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

var _bowerJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x5f\x6f\xd3\x30\x10\x7f\xef\xa7\xa8\xca\x5b\x45\x9c\x76\x2b\x43\x4c\x08\xa1\x09\xde\x91\xc6\xdb\xc4\x26\xc7\xb9\x26\x46\xb6\xcf\xd8\x17\x6d\xd5\xb4\xef\x8e\x1d\x37\x6d\x52\x92\x96\xa7\x36\xfe\xfd\xf1\xcf\x77\x67\xbf\xce\xe6\xf3\x85\xe1\x1a\x16\xb7\xf3\xc5\x56\xf8\x4c\x59\x91\x69\x24\x74\x8b\xf7\x11\xaa\x51\x83\xe5\x55\x0b\xd7\x44\xd6\xdf\xe6\x79\x25\xa9\x6e\x0a\x26\x50\xe7\x15\x66\xca\x7b\xca\x47\x94\xbc\xa1\x1a\x9d\x0f\xc2\x87\xf0\x19\x16\xee\xa1\xe0\x9e\x24\x98\xf9\x9d\x34\x40\xf3\xcf\x45\xfc\xf9\x2a\xc0\x19\x26\xea\x2f\x8b\xc0\xfa\xd5\x2a\x4b\xf0\xc2\x49\x4b\x12\x4d\xdc\x36\xd9\x69\x2e\xfb\x5f\x58\x36\x0a\x7e\xee\x6c\x0c\xf6\x90\x64\x4a\x0a\x30\xbe\x4d\x7a\x77\xff\x2d\xf1\x64\x65\xd0\xc1\x31\xc3\x72\x99\xb3\x65\x0b\xc5\x63\x63\x09\x4f\xc9\xc9\x77\x6b\x05\x3e\x83\x7b\x0a\x47\xb3\x68\xc0\xd0\x61\x9d\xc0\x53\xff\xbf\xef\xc7\xb5\x60\x4a\x30\x42\x42\x3c\xed\x6b\x22\x59\x6e\xc1\x65\x1a\x4c\x13\x03\xfd\x40\xb5\xd3\xe0\xbe\x2b\xd0\xd1\x35\x3f\xa2\xef\x1e\xd7\xec\x8a\x5d\x75\xde\x09\x28\x1d\x0f\x31\x32\xcb\x0d\xa8\x69\x79\x9f\x15\x6d\x56\xec\x66\x68\x53\x03\x2f\x2f\xdb\xf4\x59\xd1\x66\xcd\xae\x87\x36\x5e\xc9\x40\x98\x36\x48\x78\x4a\xf0\x69\x28\x25\x44\x55\xf0\x33\xda\x3d\x21\xed\xbb\xe9\xc4\x0e\x8b\x30\x48\xd9\x16\x0d\x6d\xb9\x68\x5b\xfa\xb8\x62\x9b\xd3\x60\x45\x43\x94\x66\x64\xdc\x3c\xe1\x29\xd8\x7a\x3d\xd4\x4a\x02\x3d\xad\x8c\xe8\x78\x2d\x84\x43\xa5\xfe\xb3\xb2\x23\xe4\x7d\x98\x0f\x27\x61\x04\x9a\x8b\xa7\xe9\x91\x46\xbb\xad\x39\x81\x93\xfc\x4c\x9e\x8e\x31\x2a\x17\xdc\x95\xd3\xd2\x88\xa6\x82\xac\x4f\x3b\x5c\x55\x0a\x2e\x86\x1f\xd0\xf6\x45\xd8\x9c\x38\xf1\xc2\x9f\x31\x08\x68\xd4\x5d\xb3\x43\xed\xa4\x0b\x05\x89\x8f\xd3\xa8\xec\x88\xa6\xed\x3e\x0e\x64\x5b\x74\xa3\xed\x3f\x80\xfb\x8c\x37\x03\x55\x78\xd3\xb0\x72\xf8\x9c\x11\xbc\x10\x77\xc0\x27\x2d\xfe\x61\xee\xfd\x56\x9d\xdf\xef\x3f\x0d\xb8\x5d\x3b\xd9\xe1\xfa\xb3\xd5\x30\x9d\x82\x97\x4c\xf1\x1d\x36\x34\x1d\xf2\xc8\x49\x6f\xc8\xc9\xa4\x96\xa1\xd1\x58\x9d\x79\x3d\x5a\x3c\xc5\xda\x0c\x76\x8f\x83\x36\x5d\xd2\x16\xed\xee\x46\x50\xbd\xcd\xde\x66\x7f\x03\x00\x00\xff\xff\xd7\x73\x78\xe9\x46\x06\x00\x00")

func bowerJsonBytes() ([]byte, error) {
	return bindataRead(
		_bowerJson,
		"bower.json",
	)
}

func bowerJson() (*asset, error) {
	bytes, err := bowerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bower.json", size: 1606, mode: os.FileMode(420), modTime: time.Unix(1461078977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _faviconIco = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x94\x7d\x4c\x55\x65\x1c\xc7\x9f\xab\xe6\xda\x5a\x41\xb9\xb5\x85\x93\xdd\xd6\x8b\x28\xe5\xac\x10\x73\x72\x31\x5e\xe6\xc2\xd8\xa5\x49\x2f\xda\x35\xbb\xad\xc1\x6e\xbc\x94\xcb\x90\xba\x8b\x09\x82\x12\x6f\x37\x27\x83\x08\x0a\x48\x97\x97\x80\x41\xca\xa0\x04\xe2\xc2\x15\x24\x0b\x3d\xbc\x5c\xf4\xc2\xbd\x89\x88\xbc\xc4\x25\x2e\x01\x5a\x9b\x7f\x7c\x3a\xf7\xdc\x51\xba\x66\xbf\xb3\xef\xce\x73\x7e\xe7\xfb\x79\x9e\xf3\xfc\x9e\xe7\x39\x42\xa8\xe4\xcb\xd7\xd7\x73\x57\x0b\xc3\x0a\x21\x1e\x16\x42\x04\xc8\x92\x53\x72\xc6\x9b\x5f\x8a\x55\xf7\x79\x75\xb7\xe8\xf8\x32\x5e\xdf\x64\x8a\xb1\x34\x14\x44\xd2\x60\x7a\x81\xea\xc3\xc1\x14\x25\xf8\x91\x1e\x2b\x2c\x29\x31\x42\x7f\x37\x6e\xe4\xdc\x71\xf5\x40\x63\xae\x64\x6f\x2f\x67\x66\xb4\x97\xb9\xdf\x1c\xcc\x4e\xd8\x98\x19\xeb\x65\x7a\x54\xc2\x71\xa1\x91\xca\x2c\x2d\x89\x51\x42\xd2\x47\xdc\xf6\x41\x72\xcc\x8f\x9d\x57\xbb\xec\x2d\xee\xb9\x49\x27\x8b\xf3\x73\xcc\x4e\x5e\xc1\x3d\x7d\x8d\xe6\xd2\x5d\x54\x1d\x0c\x60\xe2\xca\x45\xfe\x98\x9d\xc6\xed\x9a\xa0\xa7\xa3\x8a\xf8\xe8\x07\xdc\xb1\x5b\xff\xed\xe3\xaf\x99\x41\xe9\xd6\x0d\x17\x37\x17\x17\xb0\xb7\x95\xd1\x59\x99\x44\x73\xf1\x1b\xb8\xc6\x87\x39\x9e\xea\x47\x4d\x76\x28\x8b\x0b\x0b\x98\x12\xd7\xe2\xb4\x59\xf9\xa9\xad\x0a\xed\x66\x21\x79\xd8\x89\xde\x7a\xfd\xfc\x78\x1f\xb7\xfe\x9c\x67\xdc\xd6\xc2\xe5\xb6\x52\x86\x3a\x4f\x32\xd4\x5d\xab\x30\xb5\x47\x82\xa9\xce\xd6\x28\xed\x8f\x5f\x16\x64\xc6\xad\x65\xa8\xdf\x4a\xe1\x11\x03\x21\xeb\x84\xde\xde\x5a\x68\xb9\x31\x3b\xc6\xef\x23\x3d\x4c\x0e\x9f\x57\x7c\x1e\xb9\xc6\x1d\xf4\xb6\x16\x52\x9a\x74\x2f\xdf\x1e\xf6\xf2\xa9\x31\x82\x84\x17\x05\x15\x26\x03\x03\x17\xad\x04\x3f\x2e\x2c\x7d\x8d\xf9\xdc\x9c\x9b\xc4\xf6\xc3\x51\xa6\xaf\xda\x18\xe9\x6d\xa6\xc9\xa4\xa5\x26\xe3\x59\x2e\x75\x7e\x4d\xd9\xbe\x55\x9c\x3c\xe4\xe5\x3f\xd4\x0a\xf2\xf7\x6b\x38\x51\x94\x82\xd3\x2e\xf1\x5a\x54\x10\x3f\xd7\xa4\xb1\x30\x73\x0d\x6b\x45\x22\xce\x5f\x4e\x33\x25\xf7\xd1\xdf\x5a\x42\xbf\xa5\x44\x61\x06\xbb\xcc\xf4\x9c\xf1\xb6\x4f\x57\x1a\x19\x1e\xe8\x22\x73\x7f\x34\x99\x07\x74\x04\x3d\x71\x3f\xed\xe5\x71\xf2\x5a\x39\x69\xfc\x4c\x4b\x5f\x4b\xb1\xe2\xb3\x9d\xfd\x8a\x73\xf5\x69\x58\xab\x8d\xca\xb3\xd4\x61\xa6\x28\x55\x43\xf9\xa7\x3a\xec\xfd\x5d\xbc\x1e\xee\xc3\xf3\x4f\x0a\xd6\xf9\xa9\x38\x95\xb7\x55\x5e\xe7\x41\xea\xb2\x37\xf1\x7d\x71\x8c\xe2\xaf\xcf\x0f\xe7\xd8\x3b\x2a\x72\x75\x02\x6b\x5d\x8e\x52\xb7\xf7\x5f\x12\x1c\xd0\xf9\x2b\xfc\x7b\x6f\x69\xd8\xb0\x46\xb0\xf1\xd1\x65\x7c\xfe\xee\x4a\xcb\x84\xe3\x2c\x56\x73\x12\x75\x79\x61\x0a\x6f\xce\xd2\x90\xb3\x5b\x70\xe8\x15\x41\x6e\x9c\x3f\x3d\xed\xe6\x7f\xea\xea\x51\xcc\x36\x7f\x02\x1e\x11\x6c\x59\xbf\xdc\x52\xf0\xa6\xd0\x37\x97\x1b\xb8\x3e\x64\xa5\x36\x77\x87\xf2\xbe\xf2\x13\x0d\x07\x77\x0a\x3e\x92\xc7\x6d\x90\xe7\xec\x9a\xba\x8e\xf9\x0b\x23\x19\x1f\x44\xf3\x6a\xa4\x3f\x81\xab\x05\xcf\x3c\xb6\x8c\xf0\xe7\xee\x51\xf6\x73\xce\x5e\x5f\xc9\xd6\x59\xc5\xb8\xf3\x82\xc2\x57\x64\x46\x2b\xb5\x4e\xde\x21\xf8\xe6\x58\x02\x9d\x2d\x66\x74\x91\x3e\x68\x43\x7c\xe8\x6a\x6f\xe2\xa9\x35\x2a\xb6\x6d\x5c\x21\x2d\xed\x3f\x79\x1c\x75\x41\xf2\x06\xf7\xaf\x97\xba\x99\x99\x1a\xc5\x39\xd8\x4d\x7a\x5c\x20\x6f\x47\x08\x76\x87\x0a\xf2\xd2\x74\x94\x98\x8c\x1c\xcd\x36\x92\x9a\xbc\x87\x90\xa7\x57\xb8\xb7\x6f\x5e\x79\xc7\x19\x48\x8a\x12\xea\x7d\x3b\x7d\xa5\x53\x27\x32\x18\x75\x4a\x5c\x75\x48\x74\xfd\x58\xc5\x99\xef\xca\x70\x5c\x96\x88\x08\xf2\x63\xfd\x6a\x95\x67\xce\x52\xe4\xa6\x3b\xd9\xdb\x63\x57\xa8\xd0\xef\xd9\xfe\xa0\x25\x25\x3e\x8c\xbc\x74\x03\x59\x46\x03\x7b\x63\xc3\xd8\x12\xf8\x90\x25\x78\xed\xf2\xff\x9c\xdf\x39\x1f\xb5\xba\x7d\xb9\x5a\x9d\xae\xf2\x4a\xfe\x63\xfc\xaf\x96\x7c\x1e\xc6\xc3\xfe\x1d\x00\x00\xff\xff\xb8\x6d\x5d\xdb\x7e\x04\x00\x00")

func faviconIcoBytes() ([]byte, error) {
	return bindataRead(
		_faviconIco,
		"favicon.ico",
	)
}

func faviconIco() (*asset, error) {
	bytes, err := faviconIcoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "favicon.ico", size: 1150, mode: os.FileMode(420), modTime: time.Unix(1455561960, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3b\xed\x72\xdb\xb6\xb2\xbf\xa5\xa7\x40\x98\x7b\xa7\xca\x07\x29\x59\xb1\x6f\x53\x45\xd2\xbd\xa9\x9b\x4e\x73\x27\x5f\x13\xbb\xed\x69\x33\x99\x0c\x44\x42\x12\x63\x8a\x60\x48\xc8\xb6\xe2\xf1\x3b\x9d\x67\x38\x4f\x76\x76\xf1\x41\x82\x14\x29\x5b\x8a\x9d\x9c\x73\xa6\x9d\x69\x0c\x2d\x80\xdd\xc5\x62\xb1\x1f\xc0\x72\x78\xe7\xa7\xd7\x87\xc7\x7f\xbc\x79\x46\xe6\x62\x11\x8d\xdb\xc3\x3b\xae\x4b\x0e\x79\xb2\x4a\xc3\xd9\x5c\x90\x7e\x6f\xef\x7f\xc8\xf1\x9c\x91\x19\x77\xa3\x2c\x13\xe4\xe9\x52\xcc\x79\x9a\x79\xe4\x69\x14\x11\x39\x26\x23\x29\xcb\x58\x7a\xca\x02\xaf\x4d\x08\xcc\xfe\x35\x63\x84\x4f\x89\x98\x87\x19\xc9\xf8\x32\xf5\x19\xf1\x79\xc0\x08\xfc\x9c\xf1\x53\x96\xc6\x2c\x20\x93\x15\xa1\xe4\xc7\xa3\x9f\xdc\x4c\xac\x22\xa6\xe6\x45\xa1\xcf\x62\x98\x2b\xe6\x54\x10\x9f\xc6\x64\xc2\xc8\x94\x2f\xe3\x80\x84\x31\x00\x19\x79\xf1\xfc\xf0\xd9\xab\xa3\x67\x64\x1a\x46\x4c\xd1\x1a\xb7\xdb\xc3\x28\x8c\x4f\x80\x85\x68\xe4\x84\x8b\x84\xa7\xc2\x21\xf3\x94\x4d\x47\xce\x84\x9f\xb1\xf4\x83\xcf\x01\x18\xb3\x58\x64\xdd\x84\x47\xab\x05\x4b\xcd\x5f\x0f\x17\xec\x20\x86\xcc\x4f\xc3\x44\x90\x2c\xf5\x6b\x66\x7d\xfc\xb4\x64\xe9\xaa\x1b\x84\x99\xd0\x6d\xef\x63\xe6\x8c\x87\x5d\x35\x6b\x2b\x0e\xc2\x94\xc7\xee\x34\x62\xe7\x6e\x44\x57\x7c\x29\xd6\x00\xae\x1f\xd1\x2c\x63\x59\xc1\xdb\x76\xb8\xe9\x52\xf0\x59\xca\xcf\x5c\xc1\xce\x05\x4d\x19\x6d\x00\x1b\xfc\xdb\xb2\xce\xd3\x45\xd1\xda\x11\x49\xe8\xf3\xd8\x6e\xee\x88\x26\xa1\x33\x66\x37\x77\x10\x59\x42\x13\x96\xba\x93\xa5\x10\x3c\x2e\xfd\xd8\x9e\x25\x35\xdb\xa7\x69\x60\x35\x77\x45\x13\x84\x34\xe2\xb3\xd2\x8f\x9d\x51\xa5\x14\xe0\x20\xa1\x98\x45\x35\xa0\x5d\xd1\xe2\xc6\x95\x05\x67\x41\x76\x46\x1a\x27\x70\x24\xac\xf6\x4d\x20\xda\x5d\xe3\x35\x2e\xc1\x16\x56\x73\x57\x34\x0b\x2a\x58\x0a\x3b\x59\xf9\xb9\x2b\x3a\xb0\x3d\x3c\x8a\xdc\x39\xa3\x41\x65\x77\x6b\x7a\x76\x26\x12\x85\x01\x9a\x4b\xeb\xc7\xae\xa8\x04\x9d\xd8\xcd\x9d\xd1\xf0\xd9\x2c\x62\x65\xc5\x2b\xc1\x76\x47\xcc\xa3\x09\x4d\xcb\xbf\x0a\x8b\x12\xf0\x85\xbb\xe0\xc1\x32\x02\x27\x16\x8c\x9c\xa9\x9f\xb9\x51\xe2\x03\x48\xf0\xd4\x85\xd3\x19\xc6\x30\xac\x35\x04\xfd\x48\x22\xd8\x59\x68\xb7\x86\xd2\xb1\x8d\xc1\x4b\x11\x32\x98\x73\xf0\x9c\x17\xb2\x4d\x08\x78\x12\x18\xb5\x1a\x90\x49\xc4\xfd\x93\x27\x1a\x3a\xe1\xe7\x6e\x16\x7e\x0e\xe3\x19\x74\xf0\x14\x37\x0f\x40\xa6\x17\xb5\xd8\xa5\x51\x38\x8b\x07\x04\x9c\x24\x28\x8f\xe9\x59\xd0\x14\xc8\x0f\xc8\x41\x72\x5e\x80\xce\xdd\xb3\x30\x10\xf3\x01\xe9\x1f\xf4\x2c\x78\x18\xe7\xf0\x5e\x0e\xbf\x6c\xb7\xb4\xec\x62\xae\xbc\x31\x70\xda\x6a\xb9\x67\x6c\x72\x12\x02\xd1\x24\x61\x34\xa5\xb1\xcf\x06\x24\x06\x99\x3d\x81\x3e\xc5\xdf\x80\xf4\xf0\x47\x42\x83\x40\x72\x2d\x7f\x4d\xa8\x7f\x02\x9e\x06\x9c\xf6\x80\x08\x98\x96\x25\x70\xf8\x62\x01\x5d\x97\x28\x94\xae\x96\x0a\xb6\x21\xd0\xc0\x3f\x52\x7c\x2e\xfa\x64\xf4\x29\x04\xac\x88\x84\x48\x29\x7f\x98\xd3\x38\x88\xd8\x07\x25\x62\xb2\x84\x30\x23\xa6\x0b\x36\x72\x4e\xc3\x2c\x04\xd9\x3b\x64\xc1\x20\x1c\x81\x2d\x49\x40\xc2\x0e\xa1\xbe\x08\x39\x4c\xed\xea\x09\x09\x38\xd3\x33\x60\xf6\x37\x10\x5d\x40\xb1\xef\x0d\x15\x20\x3c\x18\xe2\xdd\xc7\x2d\x03\xfa\xe1\x62\x46\xe8\x29\x05\x23\xa1\x02\x80\xb9\x10\xc9\xa0\xdb\x3d\x3b\x3b\xf3\x30\xe0\xf1\x78\x3a\xeb\x02\x31\xf0\x35\x01\x9b\xd2\x65\x24\xba\x18\x7e\x64\x48\x82\x7f\xd8\xf3\x92\x78\x06\x74\x23\x31\x72\x14\x12\xa7\x2b\x77\xbf\x5b\x5e\x16\xac\x18\xa1\x3f\xbf\x7e\xfb\x12\x94\xf3\xf9\xab\x37\xbf\x1e\x13\x0c\xba\x40\x21\xa5\xec\x1d\x58\xf7\x21\xc4\x3f\x27\x40\x1f\x02\x0d\x0e\x41\xc6\x8c\x77\x7a\xf7\x1c\xf2\xdb\xd3\x17\xbf\xc2\xb0\xb7\x6c\x0a\x51\xd6\x1c\x55\xbb\xab\xb0\xc0\x16\xb9\x4a\x90\x52\x6e\x1b\x05\x71\x8b\x2b\x6d\x0d\x2d\x9b\x4b\xd4\xee\x98\x7d\x72\x48\x44\x27\x78\x06\x21\x22\x84\xbd\x3a\xa5\xd1\xd2\xde\xbb\x94\x7d\x5a\x86\x29\x04\x83\x10\x4c\x59\x48\x9a\xb0\x9a\xcd\xcc\xb1\xbe\xc9\x01\x62\x95\x94\x06\x18\xcc\x75\x88\x95\xc0\x89\x0c\xb4\x46\x8e\xd1\x79\x83\x23\x5b\x4e\x16\xa1\x80\xf0\xce\x0e\x07\x48\x4a\xc3\x0c\xb0\x1d\xc9\xce\x61\x29\x54\x00\x12\xba\x21\xb7\x5d\x6d\x36\x34\x0a\x5b\x00\xbf\x4c\xa8\xd8\x7a\xa3\x42\xcf\x0e\x9e\xb0\x30\x1b\xd4\x5a\x92\x87\xd0\x87\xba\xaf\x58\x51\xba\x3f\x20\xd3\x65\x2c\xf7\xb4\xc3\x4e\xe1\x3c\xdd\x93\x67\xb4\x85\xe1\x13\x87\x58\x18\xc6\x74\xd6\xce\xcb\x80\x38\x0f\xd4\x60\x3c\x98\xad\x53\xd8\x76\xa9\x29\x23\x12\x70\x7f\xb9\x80\x0e\x6f\xc6\xc4\xb3\x88\x61\xf3\xc7\xd5\xf3\x00\x50\x40\xbf\xa3\x86\x97\x51\x03\x1c\xd1\xe1\xdf\x07\x0e\xc6\xea\x4a\xdd\x06\xe4\x9d\x02\x7a\xea\xf7\x03\xe7\x3d\x76\x6a\xed\x33\x7d\xea\x27\xf4\x69\xd4\xd6\x04\x60\x46\x69\x6c\xd1\xa1\x46\x63\x87\x56\xde\x75\x76\x30\xf0\x1f\x7f\x39\x0b\x4a\xc0\x1d\x05\xe9\x76\x53\x26\x96\x69\x4c\x40\x8c\x9e\x12\xe3\x0b\x24\xdf\x29\x44\x78\xf9\x50\x9a\xbb\x06\x93\x02\xf2\xf1\xee\xff\xaf\x63\x8f\x79\x96\xa6\x3c\x7d\xc9\xb2\x0c\xe2\x53\xe8\x9e\x70\xce\xb1\xff\x12\xb1\x15\xf9\xc3\xb0\x5b\xb8\x17\xb4\x13\x9b\xbd\x0d\xd8\x63\xe9\x6b\x94\xa5\x0e\x41\x85\xfd\x25\x18\x8c\x85\x56\x63\xd4\x42\xe3\x6e\x60\x55\xda\xdc\xef\xf5\x7a\xff\x2d\x57\x39\x67\x98\xb1\x59\x00\xe3\x3e\x7a\x04\x73\x04\x09\x6a\xf2\x44\xd0\xf5\x7f\x40\x3d\x5a\x75\x5c\x93\xad\x40\x26\x18\x7e\xe6\xb1\xa0\x91\x12\x50\x5b\x3a\x87\x60\xa5\x88\xe7\x9e\x61\x5f\x3a\x1b\xdd\xdf\x18\xa6\xe8\x49\x1c\xac\x03\x48\x75\x00\x61\x02\x8f\x96\x42\x3a\x9c\x96\xe0\xc9\x40\xfa\x97\x96\x4c\x39\x75\x7b\xc2\xe1\xe0\x2d\xb4\xe7\x69\x45\x6c\x2a\x4c\xbb\x70\x43\xae\xcf\x23\x0e\xbe\x0a\xd4\x1f\xf8\x56\xc4\x67\x29\x5b\xb9\xe0\x02\x1f\x92\xbb\x8c\xb1\x7b\xd7\x60\xcd\x03\xfd\x13\xa0\x07\x95\x85\x3d\x5e\x5b\x17\x46\x37\x6a\xd0\x3a\x0b\x11\xb2\x3e\x4b\xe9\x4a\xb2\x68\x98\xc1\x19\x6e\xc6\x22\x26\x55\xd4\x85\xb8\xc3\x8c\x07\x03\x56\xc5\xae\x02\x93\x26\x02\x93\x88\x62\x40\x51\x3b\xc5\x13\xa1\xd0\xae\xdd\xda\x74\x7b\x01\x1e\xf8\x6a\x9f\x69\xe4\xd5\x9d\xc6\x1c\xb5\x10\x94\xa7\x62\x10\x35\xb4\x2e\x36\x31\xe3\x16\x14\x9d\x60\xb1\xbb\x98\xfe\x4f\x23\x48\x44\x21\xfc\x51\x72\xce\xc7\x76\xef\xe3\x04\xe8\x17\xa1\x4f\x23\x90\x88\xb2\x03\x15\x86\xfb\xfb\xc9\x79\xf1\x4f\x41\xa8\x3a\xcf\xc5\x0d\x03\xe2\x4d\xeb\x31\xe3\xef\x95\xf5\x5a\xf1\xef\x7e\x84\x23\x15\x4e\x43\x16\xe8\x25\xb7\x5a\xf7\xbb\xed\xb6\x1d\xbe\x0c\x6d\xfb\x6e\xbc\x55\x9d\xe6\x4c\xc3\x73\xf0\x70\xda\xdb\xa0\x14\x1d\x39\xc3\x4c\xd1\xdb\x23\x5d\x53\xee\xf4\x8a\x54\x8a\x60\x7b\xe4\x80\x85\x5e\x62\x34\x63\xa5\x6f\x2a\xec\x2d\x1c\x5c\x31\x47\xe3\x0a\xc2\x53\x43\x56\x6f\xac\xd2\x00\xb5\x44\x67\xfc\xe2\xe8\xe8\x98\xfc\x7c\x78\x44\x5e\xa2\x5d\xc9\xc0\x0a\x85\xa7\xca\x43\x76\x2b\xac\xb5\x2b\xe8\xf4\x59\x70\x72\xeb\x24\x77\x39\x17\xb8\x53\x5a\x8c\x9d\x6d\xaa\x0e\x85\x4b\xc1\x73\x94\x9a\x29\x35\x40\x8e\x18\x63\xd0\x30\x20\x43\x60\x3e\xce\x29\x2d\x65\x20\x21\xd5\x7d\xe4\xc0\xf9\x70\xc6\xaf\xba\x4f\x61\x5b\x60\xcc\xb8\x58\x01\x62\x98\xa7\x79\x53\xa7\x5c\x20\x43\x03\xca\x25\x0d\xbb\x38\x96\xeb\x77\xff\x06\x46\xf1\x87\x03\xaf\xdf\xdf\xf3\xf6\xf6\xbe\xf7\xfa\xfb\x07\x83\x83\x5e\x7f\xaf\x0e\xd8\x1f\x5a\x09\xe1\x06\x94\x7f\xd6\xa2\x7c\x54\x37\xdb\xc0\x14\x93\x06\x28\x65\x2b\x77\x07\xac\x3b\x15\xcb\x0c\xd6\xb8\x0e\xd3\x42\x95\xab\x27\x96\x84\x71\x57\x8c\x7c\x8b\x73\xe8\x54\xc4\x22\x0d\x96\x32\x3f\x0c\x84\x7c\x71\x61\xda\x97\x97\x8e\x52\xde\xea\x0a\x61\xc6\xf8\x90\x2f\x16\xe0\x23\xb3\x61\x91\xd4\xd5\x0d\x7b\x76\x0e\x2d\x51\x33\xc8\x02\x65\xc5\x7a\x8b\xcb\x9c\x26\x8e\x0a\x22\x96\x3e\xaa\x23\x4d\xcc\x91\xd6\x1a\xee\x56\x74\xaa\x76\x4e\xe1\xc0\x1a\x67\x95\xb6\xc1\x17\x11\x91\xcd\x91\x73\xee\x94\x77\x03\xba\xae\x9e\xf4\xf9\xda\x93\x54\x74\x00\x5e\xa6\xb2\xe7\x39\xb8\x7e\x1a\x8d\x40\x0a\xe5\x29\x0a\x54\x48\xc1\x3e\x26\x18\x0f\xa7\x37\x2a\xa1\x05\x8f\x31\xba\x6f\x92\x92\xee\xbe\xde\xe4\xaa\xb4\x36\x4f\x86\x84\xd5\xa7\x8b\xf2\x0c\x0d\x6b\x58\x7c\xe5\x17\xfe\x50\x1a\x9b\x91\xd7\xaf\x5e\xfc\x51\x36\x28\xd6\x55\x63\xe9\xc8\xb5\x6d\x7d\x5e\xb3\x76\x7a\x4c\xbb\x18\x52\xe3\x24\x76\x4e\x18\x30\x18\x94\x01\x67\xca\x91\xed\x90\x65\x03\xe5\xee\xcc\x91\x81\x9f\x98\xd7\x0c\xc8\xab\xe5\x62\xc2\xd2\x87\x2a\x07\x03\x47\x2a\x63\x59\xf9\x4f\x0a\x8c\xac\x1a\xf2\x0b\xcc\x19\xd0\xe4\x6e\xcc\x19\x8c\x59\x56\xae\x14\x9b\x5e\x18\x83\x1f\xf8\xe5\xf8\xe5\x0b\x98\x88\xf1\x34\x02\x4d\x04\x5d\x89\x7e\xab\xe1\x6f\x5d\xf4\x5b\xb2\x77\xea\xc6\xbe\x29\xfa\x2d\xdd\xe2\x42\x9c\xc7\x64\x04\xb1\x39\x1c\x04\x3f\xe2\x1e\xf4\x7a\xc8\x97\xee\x3d\x9b\x87\x18\x78\x42\x68\xb1\x8e\xef\x1d\xe6\x12\xa7\xec\xfd\xd5\x78\xe5\x70\x8d\x39\x47\x85\xf7\xc2\x38\xb5\x12\x43\xaa\x08\xc7\x55\x31\x6c\x09\xa4\x02\xde\x32\x0c\xc3\x61\x05\x01\xbc\x62\x8e\xf8\xa6\x70\x46\x41\xdd\xd5\xe0\x18\xd2\x1b\x1a\x61\x6f\x11\xb2\xd8\xb7\x53\x43\x8b\x17\x54\x44\x60\x64\xe4\x14\xb1\x00\x39\x52\x92\x26\xa0\x45\xa7\x54\xdd\x24\xf4\xe5\x81\x2f\x85\x01\x30\xdb\x35\xb1\x80\x52\x76\x30\xe7\x11\xd3\x27\x42\x20\x62\x73\x52\x44\x3a\x06\x00\x9c\x4d\xf8\xc7\x1c\x28\x68\x6a\xd7\x5b\x0f\xfe\xd3\x02\x43\xd3\xc4\x48\x5d\x0b\xf1\x50\x60\xb6\x51\x26\x42\xe4\x82\x47\x8e\x15\x96\xa2\x54\x1d\xf0\x59\x71\x5c\xa5\x54\x4e\xf3\x8d\xbe\x29\x45\x73\xcf\x5d\x1e\x47\x10\xd1\x38\x44\x05\x5b\x19\x5e\xdd\xe1\x0a\x03\x7d\x21\x90\x87\x75\x38\xe7\x2c\x14\xfe\x9c\xc8\x3d\xd7\x51\x49\xe5\x9a\x60\x0b\xca\x9f\x6f\x89\x72\x21\xc6\x2b\xa5\xa5\x54\x60\x4b\x79\x49\x5b\xf2\x2d\xc4\x75\x2b\x84\xb7\x90\xd6\x4b\x1e\xb0\x2d\x65\x05\x36\xee\x9b\x68\xd6\x6d\xd0\xdd\x42\x52\x6f\xdf\xbc\xdc\x5a\xab\x92\x05\xd8\x22\xcd\x60\xce\xb0\x09\x53\x00\xab\x22\x70\x13\x6a\x74\x23\x94\xb6\x90\xc6\x1b\x7d\xdb\xb1\xa5\x44\x68\x3c\xc3\x2b\xcb\xaf\x21\x92\x1b\x22\xb5\x85\x4c\x8e\x49\xe7\x1f\x7f\x3f\xbc\xb7\xa5\x48\xd0\xbb\xb9\xbd\x75\x46\xb3\xe5\x14\xb2\x98\x91\x73\xe8\xdc\x8a\x78\x6e\x81\xec\xd7\x11\xd5\xde\xb7\x11\xd5\xcd\x92\xfd\x3a\xa2\xea\x7f\x1b\x51\xdd\x2c\xd9\xaf\x23\xaa\x47\xdf\x46\x54\x37\x4b\xb6\x14\x6c\xe6\x01\x26\xb4\xf3\xe8\xd6\x24\xcd\x2f\x68\x26\x88\x62\x85\x2c\x93\x00\x82\xeb\x81\xe9\x2a\xdf\x5b\x69\x76\xd5\x18\x45\x5e\xde\x59\xb5\x8b\x2c\x72\x68\x15\x8e\xe0\xa3\x80\x95\x0f\xe6\xe9\xa0\x95\x0d\x9a\x64\xb0\x9c\x1b\x41\x7e\x55\xcd\x05\xe5\x80\x01\x84\xf4\x29\x44\xf9\x6d\x4c\xc1\x30\x03\xdb\x90\x80\xdd\x71\xdd\x07\x37\xf7\x5f\xbb\x85\x79\x85\xe2\x92\x60\xb6\x00\xd9\x6f\xd6\x56\x95\x5b\x35\xa9\x9e\xda\x0d\xd7\x5f\x04\x0d\x2f\xea\xcd\x49\x9f\xaa\x50\xc0\x35\x9b\x04\x4c\x41\xd4\xe3\x9f\x12\x46\xf1\x1e\x22\xd3\xa9\xcb\xb5\x1c\x0f\x47\x99\x47\x92\xfd\x3d\xef\xd1\xf7\x07\x3a\xcb\xb2\x07\xbd\x53\xca\xf6\xde\x93\x49\xdf\x74\x19\xd5\xe7\x82\x77\xf7\xfb\x8f\x0f\xa6\xfb\x45\x5a\x79\x77\x3a\x9d\x3e\xd9\x88\xdf\xbe\xfa\xb5\x16\x62\x81\x61\x35\x27\xee\xe6\x34\xd5\x0a\xde\x2c\x1c\xeb\xd3\xed\x5c\x35\x41\x38\xed\xab\x2c\xb8\x32\x65\x0a\x3b\xe0\x9e\xb0\xd5\x84\x63\xc6\x37\xe5\x20\x7d\x25\xcd\xcd\xe9\x6f\x09\x65\x29\xb3\x26\x77\x54\x51\x06\x55\x85\x01\xeb\x14\x95\x7c\xbf\x88\x66\x99\xc6\x66\xf2\x5f\x2c\xfe\x4a\x65\x43\x7e\x93\x8e\x8d\x8b\x0b\x79\xef\x7e\x79\x69\xdd\xb0\x6f\xba\x66\x2b\x5f\x9f\x6b\xa5\x36\xc7\xc3\x17\x91\x7b\x71\x21\x4f\xd3\xe5\x25\xb4\xe0\x9c\xc0\x5f\x35\xc8\x21\xb0\x7a\xf3\xc8\x7e\x71\x21\x1b\x78\x89\xbb\xc0\x62\x0a\x98\x14\xc6\xf2\x17\x3d\x97\xbf\xe8\x39\xfe\x32\xe6\x72\xe4\xbc\x7b\x67\xda\xef\xdf\x43\xe2\x1f\x84\xca\xda\x0d\x4b\xe5\x3f\x25\xd6\x4a\x6f\xe3\x88\x40\x9f\x8a\xf7\xcd\x68\xf3\x07\x04\x75\x6c\xb0\xea\xc1\x15\x34\x91\x6b\x3b\x62\x71\xa0\xaf\x95\x9d\xab\xd6\xab\x1f\xe9\x6b\xdf\xe3\xad\x5b\x37\xeb\xfa\xcd\xb6\xa7\xad\x06\x7b\x4a\xaa\x16\xe8\x2a\x6b\x8a\xd7\x67\x30\x6e\x40\x88\x05\x30\xeb\xd5\x37\x71\xea\xfa\xed\x47\xce\x23\x46\xe3\x87\xea\x52\x4d\x5e\xc1\x4d\x69\x94\xb1\xe2\x16\x4e\xca\x6e\xbb\x39\x1a\xa8\xaf\xf6\x90\x39\x7c\x23\xb3\x7e\xd2\x73\xeb\x27\x4e\x29\x8b\xd9\xba\xf4\x53\xf7\x7d\x78\x55\x97\x15\xfd\x1d\x25\xf8\x01\x16\xf2\x7a\xb2\xf9\x50\x96\x62\x28\x00\x6e\x05\x9e\x6c\xed\x50\xb6\xbc\xd2\xd3\x52\x56\x53\xfe\x32\xf5\x7f\x99\xfa\xff\x00\x53\x7f\x4d\x93\x5e\xb2\x9b\xd7\x36\x89\xb6\x1f\xb9\x0d\x7b\xa7\x4f\x62\xb3\xc9\xb3\x4c\x9c\x64\xa4\x04\xd1\x56\xb0\x80\x5c\xc3\xda\xe0\xe3\x02\xcc\x23\x23\xfd\xc6\xaf\xc8\x58\xc6\xa6\x30\x85\x0e\xf8\x01\x47\xfd\xe6\x27\x0c\xab\xa1\xc0\x52\xc9\xa6\x04\x4a\xa3\x44\x88\x31\x4b\xd2\x40\x3e\xd1\xf6\xcc\x9f\xd3\x18\x81\xca\xb2\x75\xfe\xff\xe8\xf5\x2b\x2f\x93\x5c\x86\xd3\x55\x07\x3a\xee\x5d\xcf\x88\x6d\x7a\x98\x40\xee\xbe\xb2\x09\xbb\xf9\x67\x84\x8a\x62\x6f\x7c\x21\x28\x7c\x32\x39\xd4\x21\x7d\xcd\x63\xc1\x15\xaf\x05\xa5\xe7\x82\xf2\x55\x7e\xf1\x60\xb0\xa6\xa1\xe6\x6d\x30\xe7\xc0\x41\x25\x82\xed\x08\xc1\x4c\xcc\x39\x16\x1d\x4a\xf5\x04\x6e\x01\x42\x7e\x41\x88\x7e\x13\x2c\xe1\x29\xe5\x99\xad\xed\x68\xa5\xa0\xce\x7c\xe1\xc8\xc2\x5c\x45\x0b\xef\x7f\xc9\x5b\x05\xbe\x8a\x9a\x9d\x54\x96\xb2\xca\xe2\x2d\x56\xd7\xc6\xea\xf7\xcd\x22\x22\x69\x88\x8b\xf4\xed\xb7\xb6\x28\x4d\x4c\xab\x41\x3a\x3e\xec\xe9\xb8\xb0\xa7\x23\xc2\x07\x7b\x4e\xb1\x96\xf2\x0e\xab\x89\xe5\x55\x01\xc6\x0a\xcb\x98\xd1\x6d\xc1\x71\xb2\x28\xa2\xc4\x46\x8e\x71\x50\x03\xbf\x8f\x7a\xbd\x5e\xce\x31\xde\x29\x37\x32\x78\x2d\x7e\xe4\x2d\xa7\x6b\x6a\xcf\xae\x66\xad\x3a\xbe\xc2\xa5\xfb\x43\x2e\xd7\x1f\x0a\x36\xcd\x65\x2f\xe9\x04\x0c\x6f\xd5\xb3\x7b\x0d\x4c\x5b\xe6\xdc\xbe\x13\x68\x7a\x24\xae\xb3\xea\x96\x65\xba\xf6\x9d\xc0\x8e\xaf\xb2\xa6\x22\xa1\xb1\x04\x3e\xaf\x47\x6c\xaa\x2c\xbc\x45\xab\x76\x73\xb1\x1c\x20\x53\xd5\xc8\xa6\x8e\xbf\x54\x5b\xb9\xb9\x88\xbf\x88\x3a\x34\x86\x56\xfe\x11\x40\x28\x1f\xfa\x5c\xfd\x2d\xc0\xd5\x22\x32\xdf\xb3\x6c\x92\x67\x5e\xa0\xbf\x97\x9c\x93\x8c\x83\x5f\x91\x85\x8e\x6a\x0d\xb5\xdf\x83\x21\xb6\xf2\xa7\x01\x9b\x38\x29\xf3\x20\x49\xad\x53\xda\xc6\x9f\x64\xe4\x28\xd7\xa2\x1d\xfc\x48\x51\xa6\x31\x9c\xef\x8f\x9f\x4b\x19\x2b\x84\xe4\xe7\x30\x62\xc3\x2e\x40\x75\xbf\xac\x80\xce\xab\xc9\x92\x88\xd3\x40\x6b\xb0\x8b\xc5\xee\xf2\xbb\xb6\x52\x19\x5a\xb9\x42\x8d\x34\xc5\x76\x58\xbd\x24\xe9\xaa\x32\x72\xc4\xa5\xb3\x58\x68\x49\x8e\x1c\x53\x1c\xaf\x88\xaa\x11\x20\x60\x36\xea\xf7\xc0\x0c\x14\x15\xea\xdb\x44\x8a\xbe\xfa\x66\xe0\xc3\x95\x19\x71\xab\x5c\xed\x92\x17\xaa\x97\xe1\x85\x03\x2a\x8b\xb4\x28\xf7\xaa\x4a\x12\x18\xca\xbf\x07\x74\xaa\x82\xf5\xf5\xb4\x8a\x54\x2b\x8a\xac\xab\xf9\xa5\x51\xd6\x13\xa4\xf1\x74\x41\x41\xb3\xd1\x3e\xd1\x7f\x68\x74\x46\x57\x80\x0a\x30\x0b\x57\xce\x21\x70\x82\x7c\x36\xe7\x11\x28\xe0\xc8\x99\x31\x41\x7a\xde\x63\x6f\xef\xc0\xc9\x6f\x2d\x0c\x8d\x6d\xe5\xaa\xc5\x2a\x4b\x57\xe4\x52\x8c\x00\x74\x75\xcc\x15\x82\xde\x9e\x0a\x7e\x41\x2b\xd6\x88\xbc\x45\x68\x2d\x8d\xfa\xfd\xbb\x39\xbf\x51\x18\xf5\x1a\xef\xa1\x63\x70\x5d\x59\xa4\x43\x7e\x53\x59\xe4\x38\x97\x3a\xfc\xd7\x4a\x59\x1b\xf7\xc3\x19\xc7\x0f\x11\x64\xc4\xfe\x5f\x5e\x7e\x46\x4c\xd0\xae\x84\xae\x8e\x30\x9e\xe0\x0e\x8e\x97\xc3\x32\x0f\xcb\x26\x3b\xbd\x86\xa8\xbd\xe6\x2e\xbb\x72\x01\x2d\x45\x4c\x18\x56\xe6\x67\x9b\xca\xe8\x1b\x6f\xa8\xed\xaa\xba\x46\x97\xa7\xab\xc1\xe4\xa7\x9b\x96\x6d\x3d\x90\x55\xf7\xc6\x3a\xf7\x8d\xcd\xcc\x61\x57\x54\x27\x6d\x4e\x88\x65\x79\xb9\x1e\x6e\xd5\x32\xd5\xe3\x2a\xd9\xe7\x5c\x65\x35\xc3\x95\xa5\x2a\x61\x15\xa7\x4a\x8d\x1a\x6f\x93\x64\x96\x84\xb6\xbb\x4a\x69\xaf\x29\xbf\xab\xb0\xf4\x8a\x4f\x3e\x16\xaa\x95\xb2\x4f\xa0\x59\x95\x3c\x0f\x46\xc0\x19\xfb\x24\xaf\x27\xa4\xfe\xa5\x58\xd3\x86\x50\x68\x19\x60\x8c\x41\xbc\x55\xe9\x26\x3f\xf6\x3e\x92\x95\x74\x3c\xed\x7c\x77\x77\x4d\x24\xdf\x49\x7c\x38\xad\x54\xf1\xe6\x0c\xe7\xfd\xb1\xe4\x11\xec\x65\x7f\x2c\x6d\xa9\xf3\x00\xa6\x3c\x70\xd4\x09\x45\x33\x2b\xff\x87\xde\xb7\x0c\xc8\x64\xc2\x1e\x09\x8c\x9a\x91\x4e\x4e\x00\xe4\x15\x77\xae\xad\xf4\xb7\xf5\x80\xa3\x2b\x30\xeb\x1f\x70\x6a\x2b\x35\x37\x1c\x8f\x7f\xf1\x68\xcf\x03\x23\x24\x8a\x2f\x22\xf5\x11\xde\xd7\xe1\x91\xb9\xd9\x33\xe1\x92\xac\xbf\x43\x77\x0e\x51\x97\xfa\x32\x41\xc6\x75\xf9\x07\x37\x5e\x9f\x2d\xb6\x0d\x8b\xac\x24\xec\xa5\x12\xe8\x97\x47\x47\xc7\xb0\x1d\x2c\xa5\x62\x99\x96\xe2\x22\xc4\x51\x1c\xfb\xd8\x4a\x8d\x44\x31\x21\xbf\xc0\x2f\x24\xa3\x9f\x52\xad\x42\x58\x53\xef\xba\x46\xf9\x69\x3c\x5b\x46\x70\xcc\x92\xbc\xe6\xe5\x3a\xe4\x8b\x1c\xeb\x4b\x68\xab\xaa\xa3\xeb\xd0\x53\x25\x40\xd7\xa6\x75\xb3\x6e\x57\x1f\x9b\xeb\xa7\x6c\xed\xb3\x30\x0e\xf8\x99\xc7\x63\x74\x98\x58\x6d\x6b\xfb\xda\xcb\x6d\x5c\xa3\x2a\x95\xbe\xea\x60\xe7\x05\xd5\xff\xc6\xe7\xda\x9f\xd3\x54\x5c\x7d\xa4\x61\xe8\x14\xf7\x3f\x1f\x6e\x7d\x3f\xfd\xb8\x32\xc5\xa0\xa8\x7e\x4e\xb7\x6d\x1a\xf4\x3b\x9b\x1c\xd2\xc5\x2e\xa7\x1c\x3f\xee\x35\x1a\xad\x36\x09\x14\x17\x42\x5f\x40\x26\x3f\xf6\xd5\x9f\xee\xfe\x2e\xbb\x08\x28\xb3\xf9\x7c\xf7\xc6\x94\x57\x51\x75\xb6\x56\xcc\x8a\x66\xfe\x33\x00\x00\xff\xff\xd4\xd9\x8a\x71\x91\x46\x00\x00")

func fcsLpcMotorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_fcsLpcMotorHtml,
		"fcs-lpc-motor.html",
	)
}

func fcsLpcMotorHtml() (*asset, error) {
	bytes, err := fcsLpcMotorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 18065, mode: os.FileMode(420), modTime: time.Unix(1466006150, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x59\xdd\x72\xdb\xb8\xf5\xbf\x96\x9e\x02\x8b\x1b\x4b\x23\x91\x72\xfe\xff\xce\x4e\xc7\x92\xdc\x71\x1c\x67\x36\x6d\xd6\xc9\xc4\xce\x74\xda\x24\x93\x81\x48\x50\x42\x0c\x12\x2c\x00\x4a\xd1\x66\xf5\x4e\x7d\x86\x3e\x59\xcf\x01\x40\x8a\xfa\x88\x3f\xb6\xb9\xd0\x88\x00\x0e\x7e\xe7\xe0\x7c\x83\x9c\xfc\xf4\xe2\xcd\xe5\xed\x3f\xde\x5e\x91\x85\xcd\xe5\x79\x77\xf2\x53\x14\x91\x4b\x55\xae\xb5\x98\x2f\x2c\xf9\xbf\xd3\x67\x3f\x93\xdb\x05\x27\x73\x15\x49\x63\x2c\xb9\xa8\xec\x42\x69\x13\x93\x0b\x29\x89\xa3\x31\x44\x73\xc3\xf5\x92\xa7\x71\x97\x10\xd8\xfd\xde\x70\xa2\x32\x62\x17\xc2\x10\xa3\x2a\x9d\x70\x92\xa8\x94\x13\x18\xce\xd5\x92\xeb\x82\xa7\x64\xb6\x26\x8c\x3c\xbf\x79\x11\x19\xbb\x96\xdc\xef\x93\x22\xe1\x05\xec\xb5\x0b\x66\x49\xc2\x0a\x32\xe3\x24\x53\x55\x91\x12\x51\xc0\x24\x27\xaf\x5f\x5d\x5e\x5d\xdf\x5c\x91\x4c\x48\xee\x79\x81\xbc\x5e\x6c\x42\x26\x0b\xce\x52\x7c\x80\xc7\x9c\x5b\x46\x0a\x96\xf3\x29\x5d\x0a\xbe\x2a\x95\xb6\x14\x64\x28\x2c\x2f\xec\x94\xae\x44\x6a\x17\xd3\x94\x2f\x81\x5f\xe4\x06\x43\x92\x8b\x42\xe4\x55\x1e\x99\x84\x49\x3e\x7d\x16\x9f\x0e\x81\xa9\xb0\x82\xc9\xf6\x54\x05\xe7\x74\x63\x36\x83\xa9\x35\x37\xb4\xcd\x30\x59\x30\x6d\x38\x30\xa8\x6c\x16\xfd\xb9\x5e\xb2\xc2\x4a\x7e\xfe\xf2\xf2\x86\xbc\x7e\x7b\x49\x6e\xb9\xb1\x33\x5e\x24\x8b\xc9\xc8\x2f\x78\x22\x93\x68\x51\x5a\x62\x74\x32\xa5\x33\xb5\xe2\xfa\x73\xa2\xf2\x52\x15\x20\xaf\x19\xad\xf8\x6c\x3b\xfa\xb2\x37\x8e\xa4\xb0\x3c\x06\xf1\xe3\x2f\x20\xcd\x64\xe4\x91\x02\xac\x14\xc5\x1d\x58\x47\x4e\xa9\xc8\xbd\x12\x16\x9a\x67\x47\x58\x94\xac\xc4\x93\xa1\x2d\x76\x07\x51\x22\x99\x31\xdc\xc4\xa8\x66\xfa\x00\x6c\x96\x80\x38\x65\x12\xe5\xca\x2a\xbd\xb3\xc3\x81\xf9\x67\xe2\x1c\x8d\x7c\x0b\x03\x42\xd0\x25\x32\xa9\x56\xd1\xfa\x8c\xb0\xca\xaa\x71\x58\xd9\x84\xff\x99\x4a\xd7\x2d\xf2\x0c\xcc\x18\x65\x2c\x17\x12\xe8\x4f\xde\xa9\x19\x30\x3b\x19\x92\x93\x5f\xb8\x5c\x72\x2b\x12\x46\xae\x79\xc5\x61\xa6\x99\x18\x92\x0b\x0d\x86\x1c\x12\xc3\x0a\x13\x81\x09\x45\x36\xde\x85\x5b\x71\x74\xe4\x33\xf2\xff\xa7\xa7\xbb\xdc\x41\x9f\xb5\xe4\x93\x51\xed\x60\x93\x5a\xc7\x14\xfc\x81\x18\xab\x45\x62\xe9\xb8\xbb\x64\x9a\x80\x06\xc8\x14\x84\xed\xa0\xa7\x9c\x11\x7a\x3d\xba\xa0\xc3\x6e\xc7\xaa\x3b\x5e\xc0\x10\x9f\x59\x9a\xe2\xca\xb7\x6f\xf1\x05\x3c\x6d\x36\x38\x07\x3a\xb6\x9f\xab\x32\x65\x96\x6f\x37\x69\x60\x07\x47\x04\xb0\xce\xd7\x33\x92\x31\x69\x38\xcc\x76\x7e\xdb\x3e\x6f\xe0\x07\x4e\x57\x78\xa2\x24\x4f\xcd\x19\x29\x2a\x29\x91\x0c\xb0\xd8\x76\x84\x2e\xc3\xf2\x66\x0c\x1b\x37\xe3\x6e\xb7\x97\x55\x45\x62\x85\x2a\x7a\x7d\x44\xc0\x03\x38\x49\xe1\x08\x73\x6e\x2f\x95\xba\x13\xbc\x47\xc1\x75\x3f\xdf\xbe\xf9\xdb\xd5\x35\xed\x8f\x3d\x11\x1e\xee\x90\xe6\xfd\xcd\xd5\x3b\x47\x02\x4a\x88\x03\x09\xfe\x85\x99\x1a\xd9\xfd\x8f\xbb\x9b\x5e\xbf\xdf\xed\x3a\xd2\x52\x2a\x96\xde\x38\x9d\xbe\x84\xc8\x06\x9a\x46\x2e\x8c\xf4\x46\x36\x54\x88\x43\x2d\xf8\x8a\x20\xe5\x3b\x37\xd1\x43\x9e\x7e\x2d\x56\x05\x62\x71\x48\x19\x2d\x10\x07\xe0\x10\x40\x45\xde\x3c\x9d\x8e\x5d\x97\xa0\x6b\x42\x13\x2b\x51\xd9\x9d\xda\x46\x8d\xa8\x6e\x12\x73\x08\x52\x79\x19\x23\x88\x98\x9c\x15\xa9\xf1\x3b\xbc\xc2\x83\x58\x31\x24\xc1\x4a\x5a\x5c\xd8\xc0\x0f\x61\xd0\x34\x31\x12\xc5\x06\x24\xea\xfd\xf5\xe6\xcd\x75\x8c\xde\x52\xcc\x45\xb6\xee\xc1\x42\x1f\x25\xdf\x34\xc2\xe3\xdf\x85\x79\x2e\x0a\xa6\xd7\x37\x8e\xce\x9f\x1f\x94\xd5\xd6\xd4\x65\x10\xa2\x7d\x44\xbe\x84\x30\x6e\x14\x95\x29\x9d\xc3\xea\x5b\x25\xd7\x39\xe0\xa6\x2a\x0f\x04\xb1\x54\x90\xbc\x6e\x99\x06\xd3\xc5\x25\xd3\x30\x77\x25\x79\x0e\x7f\xe0\x0d\x7b\x1a\xf2\x0a\x6a\xf4\x73\x4c\x3d\x5e\x3b\x47\x95\xe3\x75\x83\x82\xc4\x96\x7f\xb5\xc0\x8a\xc5\x4b\x26\x2b\xe7\xb6\xdd\xc7\xab\x27\x9c\x1c\x2b\x8c\xbd\xff\xe0\x4f\x39\xad\xc7\xeb\x6d\xf1\x51\x82\x00\x7f\x1c\x7d\x34\xc2\xfa\x61\x14\x14\x1e\xa9\xe6\x3d\x8a\x1b\x40\x52\x12\x0e\x1d\xc7\x31\xe9\xe1\x1c\xd7\x53\x3a\x70\xdb\x62\x3f\x1c\xd0\x7e\x13\x37\x2e\x2f\x02\xbe\x5f\x77\xa3\xb0\x82\xd1\x1a\xf4\xee\xa6\xcf\x3c\xed\xf0\x29\x76\xf0\xa8\xf8\x8c\x3a\x06\x60\x04\xf5\x4a\x07\xe8\xeb\x2a\x9f\x41\xac\xa4\x2a\xa9\x9c\x06\x40\x27\x41\x19\xcf\xd7\xaf\xd2\x9e\xcb\xdd\xc0\x23\xa2\x03\xc7\x79\x40\xa3\xfa\x1c\x88\x08\x43\x23\x05\x1c\x87\xf6\x3d\x22\x1e\xe9\x01\x9d\x80\xd0\x83\x3d\xa3\xa2\x44\x68\xd5\xc7\x38\x40\x4d\xdb\x78\x40\x06\x46\x5b\xfc\xaa\x8a\xb6\x81\x1c\x51\xed\xf6\xe0\x86\x16\x9d\xe3\x03\xb5\x3c\x87\x12\xc6\x6c\xa5\x39\x1d\x12\x5a\x2a\x23\x90\x1c\x9f\x75\x99\x1b\xfa\x29\xa8\x1d\x3a\x14\xab\x70\x8b\x53\x95\x1f\x61\xba\x02\x2b\xf5\x70\x5d\xc0\xd2\xe9\x18\xfe\x26\x1e\x3c\x96\xbc\x98\xdb\x05\xcc\x0c\x06\xdb\xac\x22\xd0\x67\xdc\xfa\x07\x81\xc8\x8d\x30\x08\x7c\x9f\xbe\x73\x55\x80\x96\x1d\xef\xad\xd2\x45\x8a\xba\xed\xe0\xfe\x58\x14\x05\xd7\xbf\xdc\xfe\xfa\x1a\x90\xbc\x74\x1f\x44\xfa\xc9\x25\x8d\x5d\xb5\xfc\xdd\xe5\xf7\xef\x6a\xa6\xc0\xf6\xeb\x01\x61\x7c\x89\x80\xb2\x0f\xa1\x9a\x3b\x97\xc5\x5d\x31\x93\x78\x0c\x7a\x0d\xa7\xb9\x58\x32\xe1\x3a\x1f\x5a\x2f\x42\xbf\x82\x8b\xae\xd6\x88\x9c\xcd\xf9\xe8\x4b\xc9\xe7\xe3\x19\x33\xfc\xe7\x3f\x0d\xfd\xd1\xbc\x05\x83\x58\xad\x9a\x91\xa0\x5f\xf5\x43\x61\x77\x42\xc2\x18\xd0\xdc\x3c\x19\x10\x3a\xa5\xe3\x66\x2d\x61\x6d\xf1\x13\x87\x10\x9b\x12\x7a\x9f\xde\xc9\xf8\xa4\xef\x09\xc1\x6c\xde\x6a\x53\xb4\xd9\x24\x61\xfb\xe6\xaa\x6b\xbe\x43\x44\x56\xcc\x19\xac\x9e\x5e\x2d\xb0\xfa\xf4\x12\x74\x4c\x7d\x61\x7b\xa7\xfd\xe9\xf4\x84\x9c\xf4\x3d\x6d\x6c\xaa\x99\xf7\xce\xde\xb3\xfe\x76\x93\xc8\x70\x87\x80\x50\xff\xfa\x26\xeb\xf9\x33\x4d\xc1\x6b\xfa\x50\x17\xc0\xff\x8a\x9d\x8d\xb8\x1c\x84\x1a\x26\xe1\x21\x60\xf9\x7e\x23\xec\xa1\xd4\x69\x6d\x05\xb0\x6a\x15\x8a\xda\x41\x45\x6b\x42\x28\x24\x0f\xac\x8a\xe0\x07\x37\x2a\xb9\x83\xdc\x46\x57\xe6\x6c\x34\xa2\xa0\x48\xa4\xc3\xa6\x03\x75\x3a\x42\xda\xa6\x48\x37\x01\xf8\xd8\xdd\x48\xbb\xbb\x7b\x55\x3b\xde\xa3\xf6\x2f\x21\x85\x28\xba\x9b\x00\x9c\xfb\xab\x42\x95\xbc\xd8\x3b\xe2\xe6\x90\x2a\x91\xca\xf0\x03\xb2\x43\xba\x9c\x1b\x03\xee\x78\x3c\x97\x3b\xea\x56\xcb\x15\xa4\x7f\x01\x8f\xae\x97\x68\x67\x64\x97\x94\x4a\xec\xeb\x3d\x80\x63\xd0\x10\xd5\x09\x7d\x1b\xc2\xf5\x0a\xc2\xbf\xaf\xd1\xef\x8d\x3b\x63\x21\x4f\x99\xc8\x8b\xe2\x94\xdb\x41\x9f\x6a\x03\x4c\x5d\xdf\xe6\x65\xc7\xac\x9b\x33\xb8\x10\x95\x70\xbe\x21\x61\x99\x85\x96\x28\xd1\x1c\x9a\x1e\xbc\xad\x18\xb2\xc4\xfe\x16\xba\x5e\x3c\xf4\x10\x22\xdf\xe2\xd5\x49\x73\xb2\x86\x22\x18\xc7\x88\xe0\x9d\x6c\x1c\x3a\x15\xd7\xd0\xd5\x92\xfa\x94\x02\x7d\xe7\x07\x77\x9a\x4f\xe4\xa7\x70\x38\x37\x89\x76\xeb\x6c\x05\xdb\xa6\xa7\xe9\x9e\x46\xa1\x34\xbd\xbf\xbd\x0c\x3d\x8c\x3b\xd2\x6e\x0a\xef\x35\x5a\x3c\xc8\x61\x6e\x29\xb8\x95\xa3\xd8\xaf\x32\x2e\xd9\x7c\xb7\xb2\x74\x3a\xd8\xfb\x3e\x42\xdf\xdb\x3a\x07\xe1\x25\x0a\xd0\x7c\xcc\xc0\x51\x96\xbc\x36\xa7\x9f\x1e\x3f\x19\xce\x69\xea\x00\x2d\xe8\xef\xc9\x60\x58\xad\xfa\x3b\x75\xc0\xc3\xc1\xfc\xd3\xd1\x58\x31\x97\xfc\x18\x9c\x5b\xc0\x08\x25\xff\xf9\x37\x75\x66\x69\x2a\xa0\xcf\xa5\x8e\x0a\xab\x6a\x53\x02\x9b\x0a\xf8\x44\x19\x10\x04\xab\xdc\x11\x29\x1c\x3e\x64\xe4\x01\x8a\x71\x49\xbd\x83\x86\x70\x68\x99\xe4\x0f\xb1\x6d\x8c\xdc\xe6\x4a\xdf\xf8\x59\xe4\xb4\xe5\xe2\x4c\x15\x98\xfc\x41\xdb\xef\x30\x79\x57\x15\x9e\xc3\x86\x70\xb8\xb7\xfd\x38\xe0\x1b\xab\xca\x80\xdc\xdd\x45\xff\x21\xba\xc9\xb2\xad\x72\x7e\x80\xb4\x78\xab\x6d\x4c\xfa\x34\xbc\x1c\x7a\x8d\x63\x6e\x8b\xf3\x4f\x8f\x82\x80\x76\x3c\x3e\x83\xbb\x85\x84\xf8\xfb\xef\xc4\xea\xaa\xf6\xb8\x83\xe4\xb8\x1f\xdb\xa1\x1b\x74\x3d\xa8\x33\xf2\x91\xae\x1a\x22\xd7\xdf\x1e\x8f\x2d\xba\x38\x8c\xb6\xed\x2a\xd2\x7d\x6a\x90\x53\x61\xb0\xfb\x4a\xdb\xc7\xc7\x04\x4d\x35\x74\xdb\x2a\xa7\x0d\xa1\x66\xc2\xec\x91\x4d\xdb\x64\xee\x30\x47\x7a\x5c\x10\xfd\xb0\xc3\x6d\xf5\xb8\xb0\x1e\x3a\xdc\xef\xeb\x5c\xa4\xed\xbb\x42\x5b\xe4\xf0\xf8\x98\xed\xd5\x2c\x17\xf6\x7f\xdf\xde\xe8\xc1\x3f\x6c\x43\xa5\xdb\xd9\x1c\xdc\x41\x1e\x6c\x41\x02\xd5\xf1\x16\xe4\x80\xec\xfe\x0e\x04\x55\xaa\x66\x5f\x1e\x6c\x30\x9c\xda\x81\x30\x16\x4e\xfa\x9d\x42\xa8\x79\xc2\xc1\x83\x53\xbc\xb3\x1f\xb9\x6b\xc1\xae\x7e\xd3\x4d\xc0\xe0\x03\xe5\x5a\x53\x57\xd1\x29\xdd\xcf\xa1\xff\xaa\xb8\x5e\xdf\x70\xc9\x13\x70\xc4\xfa\x72\x02\x8f\x26\x62\x92\xeb\x60\x8a\x52\xb2\xf5\x95\xd6\x40\x80\xd8\x75\xff\x60\x56\xc2\x26\x8b\xc0\x02\x82\x6d\xce\xe9\xa7\x00\x9f\xc0\x3d\x80\xd0\x39\x2f\xa2\x54\x41\x26\x39\xbb\x3f\xf1\x19\x01\xb7\x8f\xd2\x05\x79\x3b\x3c\xdd\xcb\x2e\x6f\xf5\xc3\x93\xb4\xfa\xa2\xfb\x90\x8b\xac\x32\x60\x81\x08\x45\xa1\x83\x7a\xf8\x19\xae\x5a\x3b\x49\xe5\x50\x83\x81\x55\xdf\xf3\xdf\xec\x4a\x91\x9b\xf9\x11\x29\xdc\x5b\x34\xfe\xd5\xdd\x9a\x26\xa5\xc6\x17\x84\x6b\xc9\xa7\x1f\x29\x4e\x82\x3a\xc5\x1c\xae\xef\x92\x67\xf6\x23\x3d\xf7\x21\x1b\x7a\x49\x33\x0f\xb6\x86\xa7\x70\xb7\xa1\x1f\x0b\x1a\x78\xfb\x80\x6d\x05\x2b\x52\xed\x06\x2b\xf1\x84\x1d\xc7\x7d\x30\x45\x0a\x08\x57\x2c\xeb\x93\x99\x6e\x78\xf9\x43\x34\x44\x74\x32\x02\x21\xcf\xe9\x0f\x52\x21\xc2\xb6\x95\x35\x83\xe4\x78\x37\xde\x3a\x03\x5e\x6b\x1f\xe7\x0d\x35\x43\xb7\xe3\x1e\x8e\xde\xef\x96\x73\x1a\x12\xd3\x3e\xc7\xdf\x44\xd9\x66\xe8\x13\xe9\xf2\xbe\xae\xbc\x61\xad\x39\xbe\xdd\xde\x63\x3e\xde\xc2\xc8\x9d\x4b\x29\xf0\xb5\x3c\x00\x41\x08\x29\x9d\xd7\xb6\x4b\x65\xec\x02\x79\x8b\xec\x53\x54\x84\x44\x11\xde\x92\x44\xda\x50\xe6\xdc\x2e\x54\x3a\xc5\x97\x16\x96\x36\xb3\xcc\xa5\x90\x29\x5d\x58\x5b\xc2\xcd\xaa\x79\x7b\x3c\x82\x0b\xa2\xbb\x1e\xfe\x05\x5f\x15\xe2\xbb\xa7\x02\xbf\xb5\xbc\x7f\xf7\x0a\x7d\x34\xc6\xf7\xf2\x2d\x21\xda\xf5\x78\xe2\x5f\xf3\xcf\x2a\x6b\xe1\x5e\x1e\x52\x25\xa6\x37\x91\xdc\x81\xb7\x7a\x09\x5f\x82\x80\x3d\x7c\x27\x31\xa0\xfd\x8f\x94\xb8\x8f\x01\xb0\x98\x28\xa9\x74\x56\x49\x98\x12\x29\x12\xd7\xe7\xaa\xa5\xd9\xd5\xd9\x80\x82\xaf\xbf\x08\x4b\xe4\x9f\xaf\xde\xa2\x3f\xf6\x6e\x44\x5e\x49\x77\x4d\x21\x11\x09\x3c\x26\xa3\xb6\x54\xb5\xcf\x82\xbd\x62\x56\x96\xf8\x8e\x0e\xae\xe8\x69\x2f\x95\xfd\x31\xd9\x35\xf6\x7e\x4a\xf7\x17\x88\x87\x93\x7a\x43\xf7\xe0\xcd\xb2\xa1\xbc\x37\xb3\xc3\x16\x80\xdf\x7e\x86\x81\x5b\x3d\x0c\xdd\x27\x8c\xaa\x80\x5b\x8e\x92\x2e\x63\x3b\x3d\x52\xc8\xa7\xaa\xb2\x78\x63\xc3\xcf\x14\x92\x24\x80\x02\x87\xf7\x7f\xf4\x1c\xeb\xf4\x04\x7d\x15\x74\xec\x52\x32\xe8\x00\x66\xe1\xd8\x93\x9d\x2f\x2d\x38\x7f\x3e\x19\x1d\xce\xe1\xfe\x11\x00\x20\xd2\x64\x84\x42\x9c\x77\x27\x23\xff\xc1\xec\xbf\x01\x00\x00\xff\xff\xa5\x4b\xf2\x32\xf8\x1b\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 7160, mode: os.FileMode(420), modTime: time.Unix(1466064858, 0)}
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
	"bower.json": bowerJson,
	"favicon.ico": faviconIco,
	"fcs-lpc-motor.html": fcsLpcMotorHtml,
	"index.html": indexHtml,
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
	"bower.json": &bintree{bowerJson, map[string]*bintree{}},
	"favicon.ico": &bintree{faviconIco, map[string]*bintree{}},
	"fcs-lpc-motor.html": &bintree{fcsLpcMotorHtml, map[string]*bintree{}},
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: k}
	}
	panic("unreachable")
}
