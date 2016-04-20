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

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x5b\x7b\x6f\xdb\xb8\xb2\xff\xdb\xfe\x14\xac\x7a\x2f\x9a\x36\x95\xec\x38\xc9\xbd\x81\xd7\xf6\x39\xdd\x6c\x17\xa7\x07\x69\x1b\x34\xdd\x3d\x58\x2c\x16\x85\x2c\xd1\xb6\x1a\x49\x54\x25\x2a\x8e\xb7\xc8\x77\x3f\x33\x1c\x52\x2f\xcb\x4e\xe2\xa4\xfb\x02\xb6\xa6\x86\xe4\x70\x38\x9c\xc7\x8f\x8f\x8c\x9e\xfc\xf0\xfe\xf4\xe3\x2f\xe7\xaf\xd9\x42\x46\xe1\xa4\x3b\x7a\x62\xdb\xec\x54\x24\xab\x34\x98\x2f\x24\x1b\xf4\x0f\xfe\x8f\x7d\x5c\x70\x36\x17\x76\x98\x65\x92\xbd\xca\xe5\x42\xa4\x99\xc3\x5e\x85\x21\x53\x6d\x32\x96\xf2\x8c\xa7\x57\xdc\x77\xba\x8c\x41\xef\x9f\x32\xce\xc4\x8c\xc9\x45\x90\xb1\x4c\xe4\xa9\xc7\x99\x27\x7c\xce\xe0\x73\x2e\xae\x78\x1a\x73\x9f\x4d\x57\xcc\x65\xdf\x5f\xfc\x60\x67\x72\x15\x72\xea\x17\x06\x1e\x8f\xa1\xaf\x5c\xb8\x92\x79\x6e\xcc\xa6\x9c\xcd\x44\x1e\xfb\x2c\x88\x81\xc8\xd9\xd9\x9b\xd3\xd7\xef\x2e\x5e\xb3\x59\x10\x72\x1a\x6b\xd2\xed\x8e\xc2\x20\xbe\x04\x11\xc2\xb1\x15\x44\x89\x48\xa5\xc5\x16\x29\x9f\x8d\xad\xa9\x58\xf2\xf4\x93\x27\x80\x18\xf3\x58\x66\xbd\x44\x84\xab\x88\xa7\xe6\xd7\xc1\x09\x5b\xc8\x21\xf3\xd2\x20\x91\x2c\x4b\xbd\x96\x5e\x9f\xbf\xe4\x3c\x5d\xf5\xfc\x20\x93\xba\xec\x7c\xce\xac\xc9\xa8\x47\xbd\xee\x25\x41\x90\x8a\xd8\x9e\x85\xfc\xda\x0e\xdd\x95\xc8\xe5\x1a\xc1\xf6\x42\x37\xcb\x78\x56\xca\x76\x3f\xde\x6e\x2e\xc5\x3c\x15\x4b\x5b\xf2\x6b\xe9\xa6\xdc\xdd\x40\x36\xfc\xef\x2b\xba\x48\xa3\xb2\xb4\x23\x93\xc0\x13\x71\xb5\xb8\x23\x9b\xc4\x9d\xf3\x6a\x71\x07\x95\x25\x6e\xc2\x53\x7b\x9a\x4b\x29\xe2\xda\xc7\xfd\x45\xa2\xde\x9e\x9b\xfa\x95\xe2\xae\x6c\xfc\xc0\x0d\xc5\xbc\xf6\xb1\x33\xab\xd4\x05\x3a\x68\x28\xe6\x61\x0b\x69\x57\xb6\xb8\x70\x75\xc5\x55\x28\x3b\x33\x8d\x13\x70\x89\x4a\xf9\x31\x18\xed\x6e\xf1\x9a\x97\xe4\x51\xa5\xb8\x2b\x9b\xc8\x95\x3c\x85\x95\x6c\x7c\xee\xca\x0e\x62\x8f\x08\x43\x7b\xc1\x5d\xbf\xb1\xba\x2d\x35\x3b\x0f\x12\x06\x3e\x86\xcb\xca\xc7\xae\xac\xa4\x3b\xad\x16\x77\x66\x23\xe6\xf3\x90\xd7\x0d\xaf\x46\xdb\x9d\xb1\x08\xa7\x6e\x5a\xff\x2a\x23\x8a\x2f\x22\x3b\x12\x7e\x1e\x42\x12\xf3\xc7\xd6\xcc\xcb\xec\x30\xf1\x80\x24\x45\x6a\x83\x77\x06\x31\x34\xeb\x8c\xc0\x3e\x92\x10\x56\x16\xca\x9d\x91\x4a\x6c\x13\xc8\x52\x8c\x0d\x17\x02\x32\xe7\x57\x55\x66\x0c\x32\x09\xb4\x5a\x0d\xd9\x34\x14\xde\xe5\x77\x9a\x3a\x15\xd7\x76\x16\xfc\x1e\xc4\x73\xa8\x10\x29\x2e\x1e\x90\x4c\x2d\x5a\xb1\xed\x86\xc1\x3c\x1e\x32\x48\x92\x60\x3c\xa6\x26\x72\x53\x18\x7e\xc8\x8e\x93\xeb\x92\x74\x6d\x2f\x03\x5f\x2e\x86\x6c\x70\xdc\xaf\xd0\x83\xb8\xa0\xf7\x0b\xfa\x4d\xb7\xa3\x75\x17\x0b\xca\xc6\x20\x69\xa7\x63\x2f\xf9\xf4\x32\x80\x41\x93\x84\xbb\xa9\x1b\x7b\x7c\xc8\x62\xd0\xd9\x77\x50\x47\xf2\x0d\x59\x1f\x3f\x12\xd7\xf7\x95\xd4\xea\x6b\xea\x7a\x97\x90\x69\x20\x69\x0f\x99\x84\x6e\x59\x02\xce\x17\x4b\xa8\xba\x41\xa5\xf4\xb4\x56\xb0\x0c\x40\x03\x7f\x94\xfa\x6c\xcc\xc9\x98\x53\x18\x44\x11\x45\x51\x5a\xfe\xb4\x70\x63\x3f\xe4\x9f\x48\xc5\x2c\x07\x98\x11\xbb\x11\x1f\x5b\x57\x41\x16\x80\xee\x2d\x16\x71\x80\x23\xb0\x24\x09\x68\xd8\x62\xae\x27\x03\x01\x5d\x7b\xba\x43\x02\xc9\x74\x09\xc2\xfe\x0c\xaa\xf3\x5d\xac\x3b\x77\x25\x28\x0f\x9a\x38\x2f\x70\xc9\x60\xfc\x20\x9a\x33\xf7\xca\x85\x20\x41\x00\x60\x21\x65\x32\xec\xf5\x96\xcb\xa5\x83\x80\xc7\x11\xe9\xbc\x07\x83\x41\xae\xf1\xf9\xcc\xcd\x43\xd9\x43\xf8\x91\xe1\x10\xe2\xd3\x81\x93\xc4\x73\x18\x37\x94\x63\x8b\x98\x58\x3d\xb5\xfa\xbd\xfa\xb4\x60\xc6\x48\xfd\xf1\xfd\x87\xb7\x60\x9c\x6f\xde\x9d\xff\xf4\x91\x21\xe8\x02\x83\x54\xba\xb7\x60\xde\xa7\x80\x7f\x2e\x61\x7c\x00\x1a\x02\x40\xc6\x5c\xec\xf5\x9f\x5b\xec\xe7\x57\x67\x3f\x41\xb3\x0f\x7c\x06\x28\x6b\x81\xa6\xdd\x23\x2e\xb0\x44\x36\x29\x52\xe9\x6d\xab\x22\xbe\xe1\x4c\x3b\xa3\x4a\xcc\x65\xb4\x3a\x66\x9d\x2c\x16\xba\x53\xf4\x41\x40\x84\xb0\x56\x57\x6e\x98\x57\xd7\x2e\xe5\x5f\xf2\x20\x05\x30\x08\x60\xaa\xc2\x64\x13\x57\xb3\x98\x05\xd7\xf3\x82\x20\x57\x49\xad\x81\xe1\xdc\xc6\x98\x14\xce\x14\xd0\x1a\x5b\xc6\xe6\x0d\x8f\x2c\x9f\x46\x81\x04\x78\x57\x85\x03\x2c\x75\x83\x0c\xb8\x5d\xa8\xca\x51\x0d\x2a\xc0\x10\xba\xa0\x96\x9d\x16\x1b\x0a\x65\x2c\x80\x2f\x03\x15\x3b\xe7\x04\x3d\xf7\xd0\xc3\x82\x6c\xd8\x1a\x49\x5e\x42\x1d\xda\x3e\x89\x42\xb6\x3f\x64\xb3\x3c\x56\x6b\xba\xc7\xaf\xc0\x9f\x9e\x2b\x1f\xed\x20\x7c\x12\x80\x85\xa1\xcd\xde\x9a\xbf\x0c\x99\xb5\x4f\x8d\xd1\x31\x3b\x57\xb0\xec\xca\x52\xc6\xcc\x17\x5e\x1e\x41\x85\x33\xe7\xf2\x75\xc8\xb1\xf8\xfd\xea\x8d\x0f\x2c\xa0\xde\xa2\xe6\x75\xd6\x40\x47\x76\xf8\xbb\x6f\x21\x56\x27\x73\x1b\xb2\x5f\x89\xe8\xd0\xf7\xbe\xf5\x1b\x56\x6a\xeb\x33\x75\xf4\x09\x75\x9a\x75\xa5\x03\x08\x43\x16\x5b\x56\x50\x6b\xac\xd0\xc6\xbb\x2e\x0e\x02\xff\xc9\xc3\x45\x20\x05\xef\x11\xa5\xd7\x4b\xb9\xcc\xd3\x98\x81\x1a\x1d\x52\xe3\x19\x0e\xbf\x57\xaa\xf0\xe6\xa5\x0a\x77\x1b\x42\x0a\xe8\xc7\x79\xf1\x0f\xab\xda\xe6\x75\x9a\x8a\xf4\x2d\xcf\x32\xc0\xa7\x50\x3d\x15\x42\x60\xfd\x0d\x72\x2b\xf7\x0f\xa3\x5e\x99\x5e\x30\x4e\x6c\xcf\x36\x10\x8f\x55\xae\xa1\x48\x1d\x80\x09\x7b\x39\x04\x8c\x48\x9b\x31\x5a\xa1\x49\x37\x30\x2b\x1d\xee\x0f\xfa\xfd\xff\x55\xb3\x5c\x70\xdc\xb1\x55\x08\x26\x7d\xf4\x19\xee\x11\x14\x69\x53\x26\x82\xaa\x7f\xc2\xe8\xe1\x6a\xcf\x36\xbb\x15\xd8\x09\x06\xbf\x8b\x58\xba\x21\x29\xa8\xab\x92\x83\xbf\xa2\xc1\x8b\xcc\x70\xa4\x92\x8d\xae\xdf\x08\x53\x74\x27\x01\xd1\x01\xb4\x3a\x04\x98\x20\xc2\x5c\xaa\x84\xd3\x91\x22\x19\xaa\xfc\xd2\x51\x5b\x4e\x5d\x9e\x0a\x70\xbc\x48\x67\x9e\x4e\xc8\x67\xd2\x94\xcb\x34\x64\x7b\x22\x14\x90\xab\xc0\xfc\x41\x6e\x1a\x7c\x9e\xf2\x95\x0d\x29\xf0\x25\x7b\xca\x39\x7f\x7e\x07\xd1\x1c\xb0\x3f\x09\x76\xd0\x98\xd8\xc9\xda\xbc\x10\xdd\x50\xa3\x75\x11\x42\x14\x7d\x9e\xba\x2b\x25\xa2\x11\x06\x7b\xd8\x19\x0f\xb9\x32\x51\x1b\x70\x87\x69\x0f\x01\xac\xc9\x9d\x80\xc9\xa6\x01\xa6\xa1\x8b\x80\xa2\xb5\x8b\x23\x03\xa9\x53\x7b\x65\xd1\xab\x13\x70\x20\x57\x7b\x5c\x33\x6f\xae\x34\xee\x51\x4b\x45\x39\x84\x41\xa8\x69\x1b\x36\x31\xed\x22\x17\x93\x60\xb9\xba\xb8\xfd\x9f\x85\xb0\x11\x05\xf8\x43\x7a\x2e\xda\xf6\x5e\x60\x07\xa8\x97\x81\xe7\x86\xa0\x11\x8a\x03\x0d\x81\x07\x47\xc9\x75\xf9\x4f\x39\x50\xb3\x9f\x8d\x0b\x06\x83\x6f\x9a\x8f\x69\xff\xbc\x6e\xd7\x24\xbf\xfd\x19\x5c\x2a\x98\x05\xdc\xd7\x53\xee\x74\x5e\xf4\xba\xdd\x2a\x7c\x19\x55\xe3\xbb\xc9\x56\x6d\x96\x33\x0b\xae\x21\xc3\xe9\x6c\x83\x5a\xb4\x54\x0f\xd3\x45\x2f\x8f\x4a\x4d\x45\xd2\x2b\xb7\x52\x0c\xcb\x63\x0b\x22\x74\x8e\x68\xa6\xb2\x7d\x23\xd8\x5b\x26\xb8\xb2\x8f\xe6\xe5\x07\x57\x66\x58\xbd\xb0\x64\x01\x34\x45\x6b\x72\x76\x71\xf1\x91\xfd\x78\x7a\xc1\xde\x62\x5c\xc9\x20\x0a\x05\x57\x94\x21\x7b\x0d\xd1\xba\x0d\x76\xda\x17\xac\x22\x3a\xa9\x55\x2e\x14\x6e\xd5\x26\x53\xdd\x6d\x52\x05\xf1\x22\x7a\xc1\x52\x0b\x45\x0d\x54\x8b\x09\x82\x86\x21\x1b\x81\xf0\x71\x31\x52\xae\x80\x84\x32\xf7\xb1\x05\xfe\x61\x4d\xde\xf5\x5e\xc1\xb2\x40\x9b\x49\x39\x03\xe4\xb0\x48\x4d\x91\xe6\xc7\x2e\xf2\x08\xcc\x68\x65\xea\xa7\x69\x4b\x53\x2d\x33\xaa\xdb\x90\x8a\x45\x81\x05\x9f\x28\x56\x76\x1f\xe2\xe7\xe1\x91\x73\x70\x7c\xe2\x1c\x0c\x8e\x9d\xc1\xe0\x70\x78\xdc\x1f\x8c\x2a\xbb\xc4\x2d\x9d\x0f\x9a\x9d\x8f\x36\x75\x36\x34\x92\xc6\x10\x95\xbe\xd5\x8c\x20\xe2\xbb\x32\xcf\x60\xde\xeb\x34\xad\x68\xa5\x11\x56\xd1\x3a\xae\x94\xd1\x79\xe9\x9b\x56\x63\xfe\x2a\x88\x51\x48\xe2\xa0\xf8\xaf\x5f\x4d\xf9\xe6\xc6\x22\x83\x6e\x4e\x10\x7a\x4c\x4e\x05\x28\x38\xf6\xb3\x51\xb9\xd1\x6b\x6b\xf6\xfa\x1a\x4a\xb2\xa5\x51\x85\x94\x95\xf3\x2d\x0f\x78\x36\x49\x54\x0e\x52\xb1\x51\x72\x73\x66\xdc\x5c\x5b\xbd\xdd\xb0\xb3\xd6\x3e\x65\x52\xdb\xd8\xab\xb6\x0c\x9e\x0c\xeb\x6b\x80\x84\xf6\xa6\x94\xf2\x21\x75\x34\x16\xad\x20\xb7\x77\x73\x43\x98\x46\xbd\x0b\x91\xca\x69\x54\x6d\xbf\x66\xdd\x8f\x32\xc5\x48\xc4\x08\xd9\xeb\x22\x18\x62\x7b\x17\xd8\x3e\x7a\x6e\x54\xef\xa1\x69\x1b\xa4\x6e\x7c\xe1\x07\xd9\x4a\xc6\xde\xbf\x3b\xfb\xa5\xee\xde\x95\x83\xbf\x9a\xb1\x77\xab\x96\xb4\x16\x7b\x74\x9b\x6e\xd9\xa4\x25\x64\xef\x0c\xdf\x11\x9a\x29\xf8\x97\x0a\x14\x3b\xe0\xd9\x90\x92\x8f\x31\x56\xf8\xc4\x5d\xc6\x90\xbd\xcb\xa3\x29\x4f\x5f\xd2\x8e\x08\xd2\x9a\x42\x96\xea\x9f\x14\x04\x59\x6d\x40\xfb\x88\xe0\x31\x00\x6e\x45\xf0\x26\x48\x52\x62\xc3\xa2\x13\xc4\x10\x95\xff\xf5\xf1\xed\x19\x74\x44\x74\x8b\x44\x83\x67\x1b\x58\xb4\x09\x46\xdb\xb0\x68\x2d\xd2\xd0\xf9\xf9\x26\x2c\x5a\x3b\x53\x05\xd4\xc5\x55\x3e\xdf\x0e\xce\x20\xaa\xdb\xc7\xfd\x3e\xca\xa5\x6b\x97\x8b\x00\x61\x20\x24\xfa\x75\x7e\xbf\x22\xb2\xbf\xe2\xbf\xdd\xce\x57\x35\xd7\x9c\x0b\x56\x78\x4a\x8b\x5d\x1b\x88\x8e\xf0\x86\x4d\x88\xb2\x46\x22\xf8\x59\xa7\x21\x38\x25\x0a\xf0\x95\x0b\xe4\x37\x03\xe7\x02\x73\xa7\xc6\x31\x6c\x36\xdc\x10\x6b\x4b\x00\x51\x3d\x2b\x1a\x55\x64\x41\x43\x04\x41\xc6\x56\x99\x99\xd9\x05\x69\x9a\x81\x15\x5d\xb9\xb4\xaf\x1f\x28\x4f\xad\x25\x65\xe8\x6d\x9b\xcc\x4c\xc6\x0e\x81\x34\xe4\xda\x23\x24\xc2\x72\xe3\x29\x12\x3c\x19\xe4\x54\xb2\x8c\xad\x0a\x7e\xc3\x09\x5b\x13\x03\x08\xe4\x02\x9b\x35\xb6\xc2\xc6\x0a\x68\xf9\x6d\x11\x87\x90\xf3\x61\xf7\xac\xe0\x48\x86\x87\x5b\x38\xaa\xaf\xb7\xcc\x05\xf0\xc1\x1e\xcb\x40\x7a\x0b\xa6\xd6\xc1\x9a\xbc\x57\x1d\xd7\xf6\xd2\x38\x5e\x4f\xa6\xdf\x40\x52\xe5\x59\x3b\x08\xfa\x01\xfb\x3d\x50\xce\x0f\x42\xaa\x85\x63\x3f\x04\x29\xe1\xd3\x3b\xca\xac\xfb\xd9\x7e\x80\xe7\x25\x24\x69\x21\xb9\x89\xe9\x30\x04\x8d\xa6\xd1\xd0\xc3\x44\x3d\x7f\x7b\x57\x85\x26\x51\xf6\x07\x09\xf5\x2a\x9e\xe7\x21\x44\xbf\x73\xbd\x39\xbc\x9b\x80\x6e\x3c\xc7\x83\x9d\x3f\x44\xc2\x8f\xe0\xce\x3c\x85\x61\x53\x7e\x37\xe1\xd0\xff\xed\xfe\xba\x74\x59\x3e\x03\x84\x35\xb6\x4e\xad\xbf\x92\xa0\x07\x7f\x17\x41\x07\x7f\x17\x41\x0f\x1f\x5d\x50\x28\x15\x41\x1e\xca\x45\xf0\x37\x60\xf0\xcc\xcd\x24\x23\x11\x58\x9e\xf8\x90\x7b\x86\xba\x8a\x31\xd6\xa9\xef\xb3\xb4\xa4\xd4\x8c\xc6\x56\x7b\xac\x6e\x89\xb3\x46\x95\x8b\x4e\x3c\xc4\xaa\x20\xa6\x02\x30\x55\xf0\x92\x81\x4b\x75\xf4\xf0\xb2\x8b\x00\x64\x0b\xfe\x78\x62\xdb\xfb\x8f\xf7\x5f\xb7\x83\x69\x95\x44\x60\x98\x2c\x01\xfc\x65\x5d\x7a\x46\xb0\x19\xe9\x00\x9e\xdf\x70\xbd\xb3\x19\xf3\xd0\x75\x19\x42\x01\x83\x3f\x88\x42\x27\xd1\x84\x0b\xcb\xc3\x39\x85\x26\x6e\xbe\x29\x2e\xa9\x19\x24\x70\x36\x27\x81\x47\x07\xce\xe1\xff\x1f\xb7\x34\xfa\x95\xcc\xf3\x37\x47\x61\xa9\x59\x1e\xb6\x43\xac\xa7\x47\x83\x93\xe3\xd9\x51\x89\xd6\x9e\xce\x66\xb3\xef\xee\xc0\xdf\xa9\xa4\xda\x8a\x9a\xa8\x12\xd4\x74\xd9\x06\xe2\x12\xa4\xbb\x03\x82\x87\x8d\x2e\x33\x58\x1b\xfb\x92\xaf\xa6\x02\xa1\xd0\x0c\x10\xb2\xc6\xdf\xdb\x71\x61\x8d\x65\x0d\x72\xb2\x27\x74\x77\xe8\xd2\xfd\xd5\xfa\x88\xa4\xa1\x07\x8d\x59\x1f\x63\xfb\xf0\x85\x02\xab\x07\x44\x15\xd5\x55\xc8\x55\xfd\x95\xf0\xb9\x71\x01\x77\x0b\xec\x3c\xd5\xfe\xd1\x02\x3c\x6f\x41\x9e\xb4\x79\x2b\x8f\x1d\x34\x4e\x63\x29\x41\xa9\x7a\xe5\xc6\xad\x69\xf3\x54\x42\xbb\x94\x71\x4e\xf0\x4a\xc2\x74\xba\xc2\x62\xa0\x55\x73\xc7\x04\x89\x35\xc2\x3b\x44\xfc\x75\x21\xa4\xee\x43\x02\xe3\x7e\x40\x41\x71\x54\xbb\xd2\x6e\x8c\x52\xbb\xf1\x29\x4f\xbc\xc8\x05\xf0\x9a\xce\xa3\x6b\x3a\xdc\x48\x65\x3c\xf6\xf5\xa9\xc7\xde\x57\xbc\xa4\x1a\x3e\x53\x12\x3d\xbb\x79\x6e\xb5\xc9\xa9\xef\x96\x5a\xaf\x91\x5a\x36\xb1\x65\xd1\x6c\xe7\x9b\x7a\x6d\xc3\x95\x8f\xa8\xdd\x0a\xfa\xd4\xa7\xaa\x5b\x54\x8d\x18\xa1\xd0\x75\x91\xd3\xee\xa5\xf4\x26\x44\xbb\x87\xda\xd7\x44\x5d\x5f\x83\x96\xd9\xec\xb2\x20\xed\x0b\xa1\x50\xf3\x23\xaa\x3e\x89\xee\x62\xd6\x87\xfd\x7e\xff\x4f\x50\x76\x12\xad\x6b\x17\x05\x7e\x34\x75\xae\xe3\xfd\xc7\x53\xad\xda\x15\xd8\xe6\x9a\x69\xbb\x45\x9f\x94\xf1\xe3\xe4\x4f\x50\x74\x5d\xd4\x35\x9d\x37\x67\xb2\x63\x78\x29\x4a\x75\x44\xb7\xe9\x10\xac\x8e\xe9\x58\x13\x28\x11\xa2\xdb\xf1\x48\xc9\x9c\x83\x6e\x7c\x4d\x53\x5c\x6d\x6e\xba\xa4\xfc\x83\x10\xd4\xc3\xc0\x11\x30\xa3\x87\x0d\xe6\x49\x50\xed\x9a\x76\xfb\x7b\xa0\x12\x04\x68\x0e\x9d\xe2\x3d\x51\xa0\x8e\x54\x6c\xfd\xac\xe8\x76\x15\x99\xa7\x71\xdb\xf4\x59\xbc\xf5\x39\x48\xae\x59\x26\xc0\xcc\xd5\x9d\x29\xcd\xa1\xf5\x69\x29\x72\xab\xbf\x32\xda\x26\x49\x5d\x06\x35\xd4\xfa\x48\xf7\xc3\x2d\x17\x85\x15\x3d\x0c\xb8\x2c\x8e\x26\x6f\x94\x8e\x89\x21\xfb\x31\x08\x61\x83\x07\x54\x5d\xaf\x1e\x53\x14\x17\x53\x49\x28\x5c\x5f\x5b\xb0\x8d\xef\x66\xd4\x13\xd9\xda\x8d\x56\xfd\xb2\x8b\x6d\x09\x5e\xb4\xb6\xf4\x22\x05\x79\x69\xbf\x87\x92\x92\xc8\x32\xef\x6c\x68\x50\x6a\x01\x0a\xe6\xe3\x41\x1f\x42\x52\xf9\xd8\xe5\x1e\xb0\xc6\xd6\xa1\xe8\xd3\x1d\xc2\x48\xed\xa8\xbe\x78\xf3\x72\x07\xcc\x02\xca\x2b\x6f\x89\x9a\x9a\x04\x81\x8a\xa7\xc5\x56\x53\xb1\x9e\xee\xd6\xd0\x6a\xc3\x90\xf5\xc3\x20\x15\x20\x75\x07\x15\xc0\x21\xf9\x2f\xb3\xf1\x11\xd3\x3f\x6e\xb8\x74\x57\xc0\x0a\x38\x4b\x5b\xf5\x61\xe0\x41\x1e\x5f\x88\x10\x0c\x70\x6c\xcd\xb9\x64\x7d\xe7\xc4\x39\x38\xb6\x8a\x08\x6f\xc6\x78\x00\x5c\xa4\xa9\x18\x05\xe8\xa3\xfd\x5b\x14\x7d\xff\x51\xf0\x31\xbe\x5c\x1b\xe4\x03\x52\x37\xe4\x84\xb6\xf5\x7b\xbc\xa4\x50\x06\xf5\x97\xdd\xb5\xab\x11\x29\x2e\x79\x5c\x5c\x8b\x5c\xc8\x14\xda\x15\xd7\x22\x96\xa5\x6e\x28\xe0\x7f\x6d\x94\x95\x7b\x11\xba\x12\xc1\x1b\x11\xf0\x71\x7c\xd3\x84\x7f\x6d\xe0\xfc\x8f\x53\xf8\x08\xee\xa6\x4a\xa5\x93\x0b\xa3\x07\xef\x61\x7b\xd5\x2c\x73\xf0\xb6\x75\xaf\xff\x1c\xf3\xd5\xcd\x5a\xde\x6a\x39\x89\x68\x1c\x1f\x28\x15\x33\x8e\x8f\x7c\xb2\x6d\x2f\x72\x6e\x3b\x5f\x50\x77\x79\x1b\x53\x5e\xf5\x15\x78\x25\xb6\x1e\xab\x07\x3c\x26\x3a\x0f\x4c\xcc\x2c\x68\xb7\x5c\xad\x6c\xdf\x9f\xaa\x97\x2a\xba\x79\xe5\x22\xa6\x9d\x57\x2d\x3e\x97\xf7\xfe\x24\x70\x63\xaa\xa4\xac\xd2\xab\xa8\x55\xc3\xae\xee\x64\x56\xa4\xb4\xdd\x4d\x4a\x67\x4d\xf5\x44\xab\x62\x57\x62\xfa\xb9\x34\xad\x94\x7f\x01\xcb\xfa\xf7\xc5\xfb\x77\x4e\xa6\xf8\x04\xb3\x15\xb6\x00\x1f\xfb\xa2\x4e\x0b\x94\xfd\xa5\x78\x21\x87\x54\x28\x19\x62\x8c\x7f\xef\x52\xb9\xa6\x53\x7f\x37\x72\xa1\xae\x01\x45\xba\xf7\xec\xe9\x9a\x4a\x9e\x29\x7e\xd8\xad\x76\x5d\x67\x8d\x16\x83\x89\x92\x11\xe2\xe5\x60\xa2\x62\xa9\xb5\x0f\x5d\xf6\x2d\xf2\x50\x0c\xb3\xea\x7f\xa8\xfd\xc0\x61\x98\x4c\x56\x5b\x82\xa0\xa6\xa5\x55\x0c\x00\xfa\x8a\xf7\xee\x6c\xf4\xdf\xea\xf8\x4d\x5f\x1f\xb7\x1f\xbf\xb5\x5e\x33\x6f\x71\x8f\xbf\x38\xda\x73\x20\x08\xc9\xf2\x71\xb5\x76\xe1\x23\x0d\x8f\xcc\x51\x99\x81\x4b\xea\xf2\x10\xd3\x39\xa0\x2e\x7a\xe4\xa4\x70\x5d\xf1\x76\xcf\x19\xf0\xe8\xfe\xb0\xe8\x2d\xe9\xf1\xe1\xa0\xa8\x76\xde\x5d\x26\x71\xe4\x51\x7a\x7b\xac\x4e\xb9\x75\xb3\xe2\x3c\xbb\x54\x83\x3e\xcf\xbe\x75\x37\x08\xfc\xcd\x66\x30\x29\x36\x83\x9b\x07\x35\x6d\x1e\x36\x22\xed\xe6\x37\x8f\x42\x37\x5f\x77\x1e\xe1\x71\xf3\xa9\xf6\x07\xbd\xd1\xea\x2e\x83\xd8\x17\x4b\x47\xc4\x98\xe6\xf0\x82\xbf\x9a\x21\x6f\xee\x93\xd0\xe8\x75\xc6\x6d\xee\x58\xbc\xe1\xf8\x1b\x7b\xa3\xb7\x70\x53\x79\xbb\x23\x42\xd3\x19\x2e\x6e\xd1\xbc\xf2\x07\x14\x27\x8d\x2e\x86\x45\xf3\x3d\xed\x7d\xbd\xf4\x3f\x7c\x7a\xea\x46\xbb\x38\x29\xbe\xee\x47\x23\x85\x4c\xc5\x91\x43\xf5\x89\x7f\xf9\xe2\xac\xef\x0c\x0e\x0f\x7b\xd1\xe7\x64\xde\xbb\x0a\x7c\x2e\x1c\x2c\x9a\x97\xfc\x8f\x66\xa3\x64\x24\xf7\x37\xd1\x86\x8d\xfe\x37\x00\x00\xff\xff\xf5\xbe\xac\x9f\x9c\x3a\x00\x00")

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

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 15004, mode: os.FileMode(420), modTime: time.Unix(1461078977, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x58\xdd\x8e\xdb\xba\x11\xbe\xb6\x9f\x82\xe1\xcd\xca\xb0\x25\x6f\x5a\xa0\x28\xd6\xd6\x16\x9b\xcd\x06\x27\x6d\xce\x26\xc8\x6e\x50\xb4\x49\xb0\xa0\x25\xca\x66\x96\x12\x55\x91\xb2\x8f\x1b\xf8\x9d\xfa\x0c\x7d\xb2\xce\x90\x94\x2c\xff\xc4\xc7\x07\x3d\x17\x86\xf9\x33\x9c\x9f\x6f\x86\x33\x43\x4d\x5f\xbc\x7e\x7f\xfb\xf8\x8f\x0f\x77\x64\x61\x72\x79\xdd\x9f\xbe\x08\x43\x72\xab\xca\x75\x25\xe6\x0b\x43\xfe\x70\xf9\xf2\x4f\xe4\x71\xc1\xc9\x5c\x85\x52\x6b\x43\x6e\x6a\xb3\x50\x95\x8e\xc8\x8d\x94\xc4\xd2\x68\x52\x71\xcd\xab\x25\x4f\xa3\x3e\x21\x70\xfa\x93\xe6\x44\x65\xc4\x2c\x84\x26\x5a\xd5\x55\xc2\x49\xa2\x52\x4e\x60\x3a\x57\x4b\x5e\x15\x3c\x25\xb3\x35\x61\xe4\xd5\xc3\xeb\x50\x9b\xb5\xe4\xee\x9c\x14\x09\x2f\xe0\xac\x59\x30\x43\x12\x56\x90\x19\x27\x99\xaa\x8b\x94\x88\x02\x16\x39\x79\xf7\xf6\xf6\xee\xfe\xe1\x8e\x64\x42\x72\x27\x0b\xf4\x75\x6a\x13\x32\x5d\x70\x96\xe2\x00\x86\x39\x37\x8c\x14\x2c\xe7\x31\x5d\x0a\xbe\x2a\x55\x65\x28\xe8\x50\x18\x5e\x98\x98\xae\x44\x6a\x16\x71\xca\x97\x20\x2f\xb4\x93\x11\xc9\x45\x21\xf2\x3a\x0f\x75\xc2\x24\x8f\x5f\x46\x97\x23\x10\x2a\x8c\x60\xb2\xbb\x54\x83\x9d\x76\xce\x66\xb0\xb4\xe6\x9a\x76\x05\x26\x0b\x56\x69\x0e\x02\x6a\x93\x85\x7f\x6e\xb6\x8c\x30\x92\x5f\xbf\xb9\x7d\x20\xef\x3e\xdc\x92\x47\xae\xcd\x8c\x17\xc9\x62\x3a\x76\x1b\x8e\x48\x27\x95\x28\x0d\xd1\x55\x12\xd3\x99\x5a\xf1\xea\x29\x51\x79\xa9\x0a\xd0\x57\x8f\x57\x7c\xb6\x9d\x7d\xdb\x9b\x87\x52\x18\x1e\x81\xfa\xd1\x37\xd0\x66\x3a\x76\x9c\x3c\x5b\x29\x8a\x67\xf0\x8e\x8c\xa9\xc8\x1d\x08\x8b\x8a\x67\x47\x44\x94\xac\x44\xcb\xd0\x17\xbb\x93\x30\x91\x4c\x6b\xae\x23\x84\x99\xfe\x0a\xdb\x2c\x01\x75\xca\x24\xcc\x95\x51\xd5\xce\x09\xcb\xcc\x8d\x89\x0d\x34\xf2\xdd\x4f\x08\xc1\x90\xc8\xa4\x5a\x85\xeb\x2b\xc2\x6a\xa3\x26\x7e\x67\xe3\xff\x67\x2a\x5d\x77\xc8\x33\x70\x63\x98\xb1\x5c\x48\xa0\xbf\xf8\xa8\x66\x20\xec\x62\x44\x2e\x7e\xe2\x72\xc9\x8d\x48\x18\xb9\xe7\x35\x87\x95\x76\x61\x44\x6e\x2a\x70\xe4\x88\x68\x56\xe8\x10\x5c\x28\xb2\xc9\x2e\xbb\x15\xc7\x40\xbe\x22\x7f\xbc\xbc\xdc\x95\x0e\x78\x36\x9a\x4f\xc7\x4d\x80\x4d\x1b\x8c\x29\xc4\x03\xd1\xa6\x12\x89\xa1\x93\xfe\x92\x55\x04\x10\x20\x31\x28\xdb\xc3\x48\xb9\x22\xf4\x7e\x7c\x43\x47\xfd\x9e\x51\xcf\xbc\x80\x29\x8e\x59\x9a\xe2\xce\xf7\xef\xd1\x0d\x8c\x36\x1b\x5c\x03\x8c\xcd\x53\x5d\xa6\xcc\xf0\xed\xa1\x0a\xc4\x81\x89\x19\x93\x9a\xc3\x14\xa2\x0b\x58\x00\xeb\x5e\x92\xa7\xfa\x8a\x14\xb5\x94\xb0\xdc\x83\x43\xac\x9d\x6d\x46\xfd\xcd\xa4\xdf\x0f\xb2\xba\x48\x8c\x50\x45\x30\xc0\x13\xa8\x99\x55\x01\x74\x9b\x73\x73\xab\xd4\xb3\xe0\x01\x85\x98\x7c\x7a\x7c\xff\xb7\xbb\x7b\x3a\x98\x38\x22\xd4\xfa\x90\xe6\xd3\xc3\xdd\x47\x4b\x02\xd6\x45\x9e\x04\xff\xfc\x4a\xc3\xd9\xfe\x4f\xfa\x9b\x60\x30\xe8\xf7\x2d\x69\x29\x15\x4b\x1f\x2c\x58\x6f\xe0\xca\x02\x4d\xab\x17\x5e\xe1\x56\x37\xb4\xd4\x72\x2d\xf8\x8a\x20\xe5\x47\xbb\x10\xa0\x4c\xb7\x17\xa9\x02\x79\x71\xc8\x05\x1d\x26\x96\x81\xe5\x00\x90\x38\xdc\x7b\x3d\xb3\x2e\x01\x44\x42\x13\x23\x11\xc5\x5e\x03\x7e\xab\xaa\x5d\xc4\xe4\x80\x54\x4e\xc7\x10\xae\x42\xce\x8a\x54\xbb\x13\x0e\x60\xaf\x56\x04\xd9\xad\x96\x06\x37\x36\xf0\x43\x36\xe8\x8a\x08\x89\x22\x0d\x1a\x05\x7f\x7d\x78\x7f\x1f\x61\x18\x14\x73\x91\xad\x03\xd8\x18\xa0\xe6\x9b\x56\x79\xfc\xbb\xd1\xaf\x44\xc1\xaa\xf5\x83\xa5\x73\xf6\x03\x58\x5d\xa4\x6e\xbd\x12\x5d\x13\xf9\x12\xee\x67\x0b\x54\xa6\xaa\x1c\x76\x3f\x28\xb9\xce\x81\x6f\xaa\x72\x4f\x10\x49\x05\x59\xe9\x91\x55\xe0\xba\xa8\x64\x15\xac\xdd\x49\x9e\xc3\x1f\x44\xc3\x1e\x42\x0e\xa0\x16\x9f\x63\xf0\x38\x74\x8e\x82\xe3\xb0\x41\x45\x22\xc3\x7f\x31\x20\x8a\x45\x4b\x26\x6b\x0c\xd1\x4d\xff\x7c\x78\xbc\xe5\x58\x3a\xcc\x69\xc3\x7f\x8b\xb5\x8e\x5f\xb0\xe5\x8f\x1a\x78\xf6\x3f\x86\x15\xaf\xd0\x6f\x05\xc7\x72\x88\x70\x8c\x86\x83\xbb\x91\x89\x43\x02\x58\xdd\xd7\xf9\x0c\x02\x38\x55\x49\x6d\xd5\x02\x45\xbd\x86\xaf\xd6\x6f\xd3\xc0\x66\x4a\x90\x11\xd2\xe1\x96\xcf\x90\x86\x5a\x0a\x88\x17\x3a\x70\x7c\x30\x88\xa0\x68\x69\x05\xd5\x4e\xaa\x79\x40\xd1\x18\x40\x91\x78\x87\x80\x9e\xc3\x3d\x70\x51\x09\x44\xf7\x1c\x47\x34\xb4\xad\x27\x32\x00\x6f\xf1\xb3\x2a\xba\x40\x59\x22\x40\xc6\x01\x05\xf1\x60\xd0\x4b\x9f\xa9\xe1\x39\x14\x09\x66\xea\x8a\xd3\x11\xa1\xa5\xd2\x02\xe9\x71\x5c\x95\xb9\xa6\x5f\x7d\x46\x81\x1e\xc0\x28\x3c\x62\xe1\x71\x33\xcc\x1b\xaa\x22\x01\xee\x0b\xd8\xba\x9c\xc0\xdf\xd4\x31\x8f\x24\x2f\xe6\x66\x01\x2b\xc3\xe1\xf6\x7a\x0b\x74\x9e\xdd\xff\x2c\x90\x73\xab\x0c\x32\x3e\x85\x71\xae\x0a\xc0\x58\xa4\x08\x65\x0f\x0f\x44\xa2\x28\x78\xf5\xd3\xe3\xcf\xef\xe0\xa8\x53\xe7\xb3\x48\xbf\xda\xeb\x8a\x40\x78\xbb\x3b\x29\x30\x41\xdf\x0c\x7c\x01\x42\xb9\x38\x87\xc3\x76\x9d\x0c\x09\x8d\xe9\xa4\xdd\x4b\x58\x57\xa3\xc4\x72\x88\x74\x09\x35\x3a\xb8\x98\x5c\x0c\x1c\x21\x18\xef\x6c\x8f\xd1\xf2\x69\xc2\xf6\x8d\x6e\x6a\x93\xe5\x88\xa2\x98\x35\xbb\x59\x5e\x2d\x30\x99\x06\x09\xfa\xb7\xba\x31\xc1\xe5\x20\x8e\x2f\xc8\xc5\xc0\xd1\x46\xba\x9e\x39\x27\x07\x2f\x07\xdb\x43\x22\xc3\x13\xa2\x48\xf9\x2f\xef\xb3\xc0\xd9\x14\x03\xf6\x03\x48\x73\xe0\xc5\x62\xe7\x20\x6e\x7b\xa5\x46\x89\x1f\x78\x5e\xae\x2e\xfa\x33\x94\xda\xf0\x59\x01\x5b\xb5\xf2\x39\xfa\x20\x41\xb7\x91\xe8\xaf\x19\x26\xf9\xbf\xf3\xd9\x83\x4a\x9e\xe1\xaa\xd2\x95\xbe\x1a\x8f\x29\x00\x89\x74\x58\x1c\x11\xd3\x31\xd2\xb6\x35\xa7\x8d\xe3\x73\x4f\x23\x2d\xdd\xbd\x06\x36\xfe\x54\xa1\x4a\x5e\xec\x69\xb8\x39\xa4\x4a\xa4\xd2\xfc\x80\xec\x90\x2e\xe7\x5a\xb3\x39\x3f\x9e\x59\x2c\x75\xa7\xb2\x7b\xe5\x5f\xc3\xd0\x56\x36\x57\x71\x9b\xad\x93\x61\xac\x0d\xdc\x34\x1d\x3a\x62\x0b\x4b\x0f\xfd\xd9\x1c\x8e\x6d\xf9\x77\x42\x7b\xe3\x31\xc9\x19\x34\xcc\x25\x28\x36\x22\x2c\x33\x50\x59\x93\x8a\x43\xed\xc4\x6e\x56\x93\x25\xf6\x3f\xd0\x15\xa1\xb6\x23\x52\xc0\x1d\x82\xd6\xba\xe2\x64\x0d\xb9\x34\x8a\x90\x83\x73\xee\xc4\x17\x3c\x27\x64\x7b\x6f\xe2\x3d\xb3\x20\x31\x7e\x7a\xbc\xf5\x65\xad\xb5\xcb\xfb\xda\xa6\x9c\x12\xdb\x62\x07\x8c\x05\xce\x12\xed\xa6\x9c\xa0\x5d\x07\xf5\x77\x72\x9e\x6d\x6f\x7e\x98\xe7\x7a\x3d\xec\x80\xce\x80\x0e\x62\x53\x14\x00\x5d\xc4\xc0\x4d\x4b\xde\x24\x24\xb7\x3c\x39\x93\x89\xed\xc8\x0e\x78\xd8\xd5\xb3\x59\x28\x63\x91\x0f\x53\x81\x79\xbe\x9b\x8d\x1c\x33\xbf\xff\x04\xfb\xdc\x46\xd4\xd9\x9c\x31\xed\x1e\xe3\x08\xeb\xe7\xf2\x60\xc5\x5c\xf2\x63\x4c\xec\x86\x75\x5b\x9b\xb8\x5d\xf2\xb2\xbb\x58\x0c\xda\xcc\xdd\x26\xee\xb3\x24\xe2\x51\xcc\xd0\x47\x64\x5a\xae\x90\xf8\x86\x94\xfc\xf7\x3f\xb7\xd4\xc5\xa3\x8f\x7c\x17\x3d\x80\x3b\x79\xd1\xf5\x82\x17\xbc\xdd\xdd\x77\xd1\x69\x9d\xb0\x22\x5b\xd2\xbd\x42\x8c\xf7\xbb\x61\x69\xb9\xb8\x9a\x64\x2b\x21\x4e\x7b\xdb\xd3\x1d\xf7\x3a\xf7\xb9\x8e\xb2\x43\x51\xe6\xfb\x4b\x16\xdc\x70\x5b\x42\x71\xf7\x6b\x2b\x27\x15\x1a\x5f\x9b\x98\x57\x5f\x1c\x6a\x51\x31\xa1\xed\x5e\x67\xcb\x42\x70\xa4\xc0\x82\xc6\x87\xe5\xb5\x53\x60\x61\xdf\x97\xd7\x1f\xe3\x24\xd2\x6e\x9b\xd2\xd1\xad\x19\x9e\x73\xbc\x9e\xe5\xc2\xfc\xff\xc7\x5b\xdb\xdd\xc0\x1e\xdd\xb8\xac\xb5\x39\xe8\x80\x7e\x35\xf5\x7b\xaa\xe3\xa9\xff\x80\xec\x74\xe6\x47\x48\xd5\xec\xdb\xc9\x04\xd8\xc2\x0e\x84\x91\xb0\xda\xef\x24\x3e\x88\x1f\x0e\x59\x26\xc5\xce\xfd\x48\xa7\x07\xa7\x06\x6d\x21\x80\xc9\x67\xca\xab\x8a\x7e\xc5\xfb\x40\xe9\xfe\x05\xfc\x57\xcd\xe1\xe5\xc1\x25\x84\x24\xf4\x1e\xbe\x33\x82\x21\xdc\x78\xc9\x2b\xef\x8a\x52\xb2\xf5\x5d\x55\x01\x01\xf2\x6e\xd2\xbf\x5e\x09\x93\x2c\xbc\x08\xb8\xb4\x73\x4e\xbf\x7a\xf6\x09\x03\xa4\xe8\x9c\x43\xb4\x2b\xc8\xad\x57\x27\x5d\x47\xb5\xc8\x43\x5d\xda\x5b\xde\x4d\xa1\xf6\x79\xeb\xbc\x7e\x68\x49\xa7\xac\x9d\xe2\x5c\x64\xb5\xc6\x4b\x87\xaa\xd0\x61\x33\x7d\x82\xb6\x6f\x27\xab\x1c\x22\xe8\x45\x0d\x9c\xfc\xcd\xae\x16\xb9\x9e\x1f\xd1\xc2\xbe\xa5\xe1\xe1\x03\xfc\xe8\xb4\xac\xf0\xfd\xbf\x96\x3c\xfe\x42\x71\x11\xe0\x14\x73\x78\x2f\x48\x9e\x99\x2f\xf4\x9a\x4e\xb6\x67\x80\x9b\xf7\x35\x8c\x7c\x4b\x48\xbf\x14\xd4\xcb\x76\x17\xb6\x73\x59\x91\x6a\xf7\xb2\x12\x47\xd8\xb3\xd2\x87\x31\x52\xc0\x75\xc5\x7e\x67\x3a\xab\x5a\x59\xce\x88\x96\x88\x4e\xc7\xa0\xe4\x35\xfd\x9d\x20\x44\xb6\x5d\xb0\x66\x90\x72\x9e\x27\xdb\x60\xc0\x16\xfb\xbc\x68\x68\x04\xda\x13\x27\x24\xba\xb8\x5b\xce\xa9\x4f\x4c\xfb\x12\xff\x2d\xca\xae\x40\x97\x31\x97\xa7\x1a\xaa\x56\x74\xc5\xf1\xe3\xd5\x9e\xf0\xc9\x96\x8d\xdc\xe9\xe5\x41\xae\xe1\x9e\x11\x5c\x21\x78\x00\x37\xbe\x4b\x65\x64\x2f\xf2\x96\xb3\x4b\x51\x21\x12\x85\xd8\x9c\x8a\xb4\xa5\xcc\xb9\x59\xa8\x34\xc6\x17\x93\xa1\xed\x2a\xb3\x29\x24\xa6\x0b\x63\x4a\x68\x68\xdb\x8f\x43\x63\xe8\xab\x6d\x57\xfd\x17\xfc\x60\x10\xc3\x73\xb1\xc0\x4f\xa9\x9f\x3e\xbe\xc5\x18\x8d\xf0\xb3\x5b\x47\x89\x0e\x6c\x10\x9b\xf6\x2b\xde\xac\x36\x06\x9e\x33\x3e\x55\x62\x7a\x13\xc9\x33\x44\xab\xd3\xf0\x0d\x28\x18\xe0\xfb\x68\x48\x07\x5f\x28\xb1\xdf\xfa\x60\x33\x51\x52\x55\x59\x2d\x61\x49\xa4\x48\xdc\xd8\xd5\x68\xb3\x8b\xd9\x90\x42\xac\xbf\xf6\x5b\xe4\x9f\x6f\x3f\x60\x3c\x06\x0f\x22\xaf\xa5\x2d\x86\x24\x24\x5e\xc6\x74\xdc\xd5\xaa\x89\x59\xf0\x57\xc4\xca\x12\x5f\xea\xf0\xb2\x49\x83\x54\x0e\x26\x64\xd7\xd9\x1b\xf7\x3e\xeb\x6d\xbf\x6f\xc2\x33\x04\xa6\xf6\xdb\x60\x5d\x40\x17\xa9\xa4\xcd\x95\xd6\x02\x0a\x99\x4c\xd5\x06\x5b\x5d\xfc\xfe\x27\x49\x02\x3e\x03\xb1\xee\x8f\x5e\x63\x85\x9c\x62\x94\x80\x75\x36\x19\x82\x74\x58\x05\x81\xd3\x9d\x4f\x98\xb8\x7e\x3d\x1d\x1f\xae\xe1\xf9\x31\x30\x40\x4e\xd3\x31\x2a\x71\xdd\x9f\x8e\xdd\x97\xe8\xff\x05\x00\x00\xff\xff\xae\xa3\x0c\x24\x51\x17\x00\x00")

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

	info := bindataFileInfo{name: "index.html", size: 5969, mode: os.FileMode(420), modTime: time.Unix(1461078438, 0)}
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
