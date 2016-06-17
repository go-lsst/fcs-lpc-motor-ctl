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

var _fcsLpcMotorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x3c\xfd\x72\xdb\xb6\x93\x7f\xcb\x4f\x81\x30\x77\x53\xa7\x09\x29\x59\xb1\xaf\xa9\x22\xf9\x2e\x75\xd3\x69\x6f\x92\x34\x13\xe7\xae\xd7\x66\x32\x19\x88\x84\x24\xc6\x14\xc1\x90\x90\x6d\xc5\xe3\x77\xba\x67\xb8\x27\xbb\x5d\x7c\x90\x00\x4d\xea\x2b\x76\x72\xf7\x9b\x76\xa6\x16\xb8\x00\x76\x17\x8b\xc5\x62\x17\x58\x64\x78\xef\xe7\xdf\x4f\xde\xfe\xf9\xfa\x39\x99\x89\x79\x72\xbc\x37\xbc\xe7\xfb\xe4\x84\x67\xcb\x3c\x9e\xce\x04\xe9\xf7\x0e\xfe\x85\xbc\x9d\x31\x32\xe5\x7e\x52\x14\x82\x3c\x5b\x88\x19\xcf\x8b\x80\x3c\x4b\x12\x22\xdb\x14\x24\x67\x05\xcb\xcf\x59\x14\xec\x11\x02\xbd\xff\xa3\x60\x84\x4f\x88\x98\xc5\x05\x29\xf8\x22\x0f\x19\x09\x79\xc4\x08\x7c\x4e\xf9\x39\xcb\x53\x16\x91\xf1\x92\x50\xf2\xd3\xe9\xcf\x7e\x21\x96\x09\x53\xfd\x92\x38\x64\x29\xf4\x15\x33\x2a\x48\x48\x53\x32\x66\x64\xc2\x17\x69\x44\xe2\x14\x80\x8c\xbc\xf8\xed\xe4\xf9\xab\xd3\xe7\x64\x12\x27\x4c\xd1\x3a\xde\xdb\x1b\x26\x71\x7a\x06\x2c\x24\x23\x2f\x9e\x67\x3c\x17\x1e\x99\xe5\x6c\x32\xf2\xc6\xfc\x82\xe5\x1f\x42\x0e\xc0\x94\xa5\xa2\xe8\x66\x3c\x59\xce\x59\x6e\x7e\x03\x1c\xb0\x87\x18\x8a\x30\x8f\x33\x41\x8a\x3c\x6c\xe8\xf5\xf1\xd3\x82\xe5\xcb\x6e\x14\x17\x42\x97\x83\x8f\x85\x77\x3c\xec\xaa\x5e\x5b\x71\x10\xe7\x3c\xf5\x27\x09\xbb\xf4\x13\xba\xe4\x0b\x71\x03\xe0\x87\x09\x2d\x0a\x56\x54\xbc\x6d\x87\x9b\x2e\x04\x9f\xe6\xfc\xc2\x17\xec\x52\xd0\x9c\xd1\x16\xb0\xc1\xbf\x2d\xeb\x3c\x9f\x57\xa5\x1d\x91\xc4\x21\x4f\xed\xe2\x8e\x68\x32\x3a\x65\x76\x71\x07\x91\x65\x34\x63\xb9\x3f\x5e\x08\xc1\x53\xe7\x63\x7b\x96\x54\xef\x90\xe6\x91\x55\xdc\x15\x4d\x14\xd3\x84\x4f\x9d\x8f\x9d\x51\xe5\x14\xe0\x20\xa1\x94\x25\x0d\xa0\x5d\xd1\xe2\xc4\xb9\x82\xb3\x20\x3b\x23\x4d\x33\x58\x12\x56\xf9\x36\x10\xed\xae\xf1\x1a\x97\x60\x73\xab\xb8\x2b\x9a\x39\x15\x2c\x87\x99\xac\x7d\xee\x8a\x0e\x6c\x0f\x4f\x12\x7f\xc6\x68\x54\x9b\xdd\x86\x9a\x9d\x89\x24\x71\x84\xe6\xd2\xfa\xd8\x15\x95\xa0\x63\xbb\xb8\x33\x1a\x3e\x9d\x26\xcc\x55\x3c\x07\xb6\x3b\x62\x9e\x8c\x69\xee\x7e\x55\x16\x25\xe2\x73\x7f\xce\xa3\x45\x02\x9b\x58\x34\xf2\x26\x61\xe1\x27\x59\x08\x20\xc1\x73\x1f\x56\x67\x9c\x42\xb3\xce\x10\xf4\x23\x4b\x60\x66\xa1\xdc\x19\xca\x8d\xed\x18\x76\x29\x42\x06\x33\x0e\x3b\xe7\x95\x2c\x13\x02\x3b\x09\xb4\x5a\x0e\xc8\x38\xe1\xe1\xd9\x53\x0d\x1d\xf3\x4b\xbf\x88\x3f\xc7\xe9\x14\x2a\x78\x8e\x93\x07\x20\x53\x8b\x5a\xec\xd3\x24\x9e\xa6\x03\x02\x9b\x24\x28\x8f\xa9\x99\xd3\x1c\xc8\x0f\xc8\x51\x76\x59\x81\x2e\xfd\x8b\x38\x12\xb3\x01\xe9\x1f\xf5\x2c\x78\x9c\x96\xf0\x5e\x09\xbf\xde\xeb\x68\xd9\xa5\x5c\xed\xc6\xc0\x69\xa7\xe3\x5f\xb0\xf1\x59\x0c\x44\xb3\x8c\xd1\x9c\xa6\x21\x1b\x90\x14\x64\xf6\x14\xea\x14\x7f\x03\xd2\xc3\x8f\x8c\x46\x91\xe4\x5a\x7e\x8d\x69\x78\x06\x3b\x0d\x6c\xda\x03\x22\xa0\x5b\x91\xc1\xe2\x4b\x05\x54\x5d\xa3\x50\xba\x5a\x2a\x58\xc6\x2d\x84\xcc\x19\xb8\x13\x20\xd2\x0c\x24\xe4\x11\x1a\x8a\x98\xa7\x23\xaf\x5b\xca\x14\xda\xc5\xf3\x29\xa1\xe7\x14\x56\xb1\xda\xa1\x67\x42\x64\x83\x6e\xf7\xe2\xe2\x22\x40\x8f\x24\xe0\xf9\xb4\x5b\xc0\xda\x2c\xba\x11\x9b\xd0\x45\x22\xba\xe8\x1f\x14\x88\x83\x7f\x38\x08\xb2\x74\x0a\x88\x13\x31\xf2\x14\x12\xaf\xab\xd0\x5a\xa6\x82\xa4\x74\xce\x46\xde\x02\xbc\x18\x2c\x79\x24\xa1\x63\x54\x1d\x70\x64\x72\x8f\x9c\xd3\x64\x01\xb5\xe7\x31\x10\xe1\xf0\x9d\xb3\x4f\x8b\x38\x07\x1f\x06\x7c\x00\x0b\x49\x1b\xd6\x0c\x36\xf4\x0b\x10\x58\x89\xf5\x75\x09\x10\xcb\xcc\x69\x60\x30\x37\x21\x56\x73\x44\xa4\x7f\x30\xf2\xcc\x54\x19\x1c\xc5\x62\x3c\x8f\x05\x78\x25\xf6\x2e\x46\x72\x1a\x17\x80\xed\x54\x56\x0e\x9d\x1d\x0e\x48\xe8\x82\x9c\x16\x9c\x0b\xd4\xe0\x6e\xa5\xc2\xf0\x65\x3c\x9c\xce\x6b\xe5\x31\xed\xa3\x62\xc4\xc5\xa0\x71\x01\x3c\x82\x3a\x80\x7f\x50\xac\x7c\x90\xc0\x01\x99\x2c\x52\x39\xa7\xfb\xec\x1c\xd4\xe0\x81\x54\xad\x0e\xee\xfa\x1c\x5c\x38\x68\xb3\x8f\xb8\x3e\xcc\x68\x1a\x25\xcc\xf4\xf1\x1e\xaa\xc6\xa8\x4f\x9d\x73\x98\x76\xa9\x29\x23\x12\xf1\x70\x31\x87\x8a\x60\xca\xc4\xf3\x84\x61\xf1\xa7\xe5\x6f\x11\xa0\x80\x7a\x4f\x35\x77\x51\x03\x1c\xd1\xe1\xef\x43\x0f\x5d\x4c\xa5\x6e\x03\xf2\x4e\x01\x03\xf5\xfd\xd0\x7b\x8f\x95\x5a\xfb\x4c\x9d\xfa\x84\x3a\x8d\xda\xea\x00\xcc\x28\x8d\xad\x2a\x54\x6b\xac\xd0\xca\x7b\x93\x1d\xf4\x57\x8f\xbf\x9c\x05\x25\xe0\x7d\x05\xe9\x76\x73\x26\x16\x79\x4a\x40\x8c\x81\x12\xe3\x0b\x24\xbf\x5f\x89\xf0\xfa\x91\x5c\xa5\x4a\xc9\xfe\x13\x8c\x48\x44\x11\xeb\x6b\x2a\xc0\x8c\xa0\xb8\x83\xef\xff\xd5\xb3\xdb\x3c\xcf\x73\x9e\xbf\x64\x45\x01\x6e\x15\x54\x8f\x39\xe7\x58\x7f\x8d\xd8\x2a\xb7\x77\xd8\xad\xac\xe2\x26\x36\x12\xbc\xdb\xbf\x8d\xe4\x06\x46\xd2\xf1\x43\xd1\xcc\x20\xa1\x19\xc3\x48\x6b\x40\x0e\x0f\x82\xc7\x3f\xc8\x51\x74\x3a\x9a\x4b\xf2\xa4\xa7\xbe\x2b\xac\x7e\xc8\x13\x0e\xa4\x61\xed\xec\xfb\xbe\x42\x08\x98\xfc\xa3\x5e\x4f\x6a\x84\xae\xbe\x98\x81\xd5\x94\x44\xb7\x34\xcd\x7a\x2a\xbf\xc8\x2a\x99\x3e\xc0\x97\x77\xfc\x42\xe2\xbc\x53\x0b\x85\x3c\xaf\x55\xe2\x35\x5a\x0c\xd3\x2c\x55\x58\x29\x40\x0c\xdc\x87\x8b\x42\x40\x0f\x35\x6a\x64\xd4\x68\x71\xa7\x9c\x9f\x83\x5e\xef\x9f\xe5\x52\x35\x73\x58\x02\x8c\x56\xf6\x08\xc6\x67\x12\xd4\xa6\xe0\x50\xf5\x6f\x40\x3d\x59\xc2\x7c\xea\x48\x11\xa2\xf0\xf8\x33\x4f\x05\x4d\xd4\x2a\xdf\x93\x3a\x17\x2d\x15\xf1\x52\xe1\x0e\xb5\x76\xc8\xfa\x56\x17\x51\x77\xe2\xb0\xc5\xc1\x2c\x0f\xc0\x45\xe3\xc9\x42\x2a\x47\xa7\x23\x78\x36\x90\x6a\xdb\x91\xe1\xbe\x2e\x8f\x39\xcc\xcd\x5c\x2b\x74\x27\x61\x13\x61\xca\xab\xf5\x70\x9a\xb3\xa5\x0f\x2b\xeb\x11\xb9\xcf\x18\x7b\xb0\x01\x6b\x01\x18\x51\x01\xcb\xa3\x36\xb0\x27\x37\xc6\x85\x9e\xa5\x6a\x74\x93\x85\x04\x59\x9f\xe6\x74\x29\x59\x34\xcc\x60\x0f\xbf\x60\x09\x93\xca\xed\x83\xcf\x67\xda\x83\x5a\xd6\xb1\x2b\xa7\xb0\x8d\xc0\x38\xa1\x68\xa7\x1a\xbb\x04\x22\x16\xda\x62\x58\x93\x6e\x0f\x20\x00\x13\x10\x32\x8d\xbc\x3e\xd3\x78\x3e\x50\x09\x2a\x50\xa6\x4d\x35\x6d\x32\x79\xa6\xdd\x9c\xc6\xa9\x3d\xbb\x78\xf4\x32\x49\xf8\x85\x0f\x56\x55\xc9\xb9\x6c\xdb\xfd\x1e\x3b\x40\xbd\x88\x43\x9a\x80\x44\xd4\x66\x56\x63\xb8\x7f\x98\x5d\x56\x7f\x2a\x42\xf5\x7e\x3e\x4e\x18\x10\x6f\x1b\x8f\x69\xff\xc0\xd5\x6b\xc5\xbf\xff\x11\x96\x54\x3c\x89\x59\xa4\x87\xdc\xe9\x7c\xdf\xdd\xdb\xb3\xed\xd3\xd0\x36\x01\xc6\xe5\x6a\xd2\x9c\x49\x7c\x59\x19\x1a\x94\xa2\x27\x7b\x98\x2e\x7a\x7a\xa4\x25\x2b\x3d\xb7\x2a\x8c\x25\x58\x1e\x79\xe0\x66\x2c\x3c\xe2\x84\xce\x2a\xe4\xa8\xbc\xb4\xaa\x8f\xc6\x15\xc5\xe7\x86\xac\x9e\x58\xa5\x01\x6a\x88\x60\xef\x4e\x4f\xdf\x92\x5f\x4e\x4e\xc9\x4b\xb4\x2b\x05\x58\xa1\xf8\x5c\x77\x6d\x32\x5c\x40\xa9\x11\x2c\xc7\xd2\xad\x0d\x66\xaf\xc6\x80\x5e\x3d\x5e\x69\xcf\xa4\x5e\x94\x53\xe4\x39\xc3\xb7\xcf\x06\x54\x85\xc2\xa5\xe0\x25\x4a\x3d\x0c\xd5\x40\xb6\x38\x46\x5f\x79\x40\x86\x30\xdc\xb4\xa4\xb4\x90\xfe\xb3\x5c\x20\xda\xd0\xbf\xea\x3e\x83\x89\x84\x36\xc7\xd6\x98\x01\xc3\x2c\x2f\x8b\x3a\x40\x06\xa9\x1b\x50\x39\x37\x30\xef\xc7\x52\x62\xfe\x7f\x81\x19\xfd\xf1\x28\xe8\xf7\x0f\x82\x83\x83\x1f\x82\xfe\xe1\xd1\xe0\xa8\xd7\x3f\x68\x02\xf6\x87\x56\xf8\xbe\x02\xe5\x5f\x8d\x28\x1f\x37\xf5\x36\x30\xc5\xa4\x01\x4a\xd9\xca\xf9\x84\xfd\x80\x8a\x45\xa1\xa7\xcd\x85\x69\xa1\xca\xd1\x13\x4b\xc2\x38\x2b\x46\xbe\xd5\xca\xf5\x6a\x62\x91\x26\x4e\x19\x2c\x06\x42\xbe\xba\x32\xe5\xeb\x6b\x4f\xa9\x7b\x7d\x84\xd0\xe3\xf8\x84\xcf\xe7\xe0\x1a\x16\xc3\x2a\x04\x6f\x6a\xf6\xfc\x12\x4a\xa2\xa1\x91\x05\x2a\xaa\xf1\x56\x47\x6f\x6d\x1c\x55\x44\x2c\x7d\x54\x46\x80\x18\x23\xa0\xd7\x84\x5f\xd3\xa9\xc6\x3e\xd5\x96\xd7\xda\xcb\x99\x86\x50\x24\x44\x16\x47\xde\xa5\xe7\xce\x06\x54\xad\xef\xf4\x79\x75\xa7\x7b\xbe\xdf\x8c\x42\x79\x17\xb0\x4b\xd5\x34\xa0\x04\x57\xdd\xf0\x08\xbb\x11\x07\x4d\x40\x40\x6e\x7f\x05\xaa\x04\x64\xaf\x20\xf4\xc5\xf2\x5b\x15\xde\x9c\xa7\x18\xef\xb6\x09\x50\x57\x6f\xd6\xb9\x2e\xc8\xd5\x9d\xc1\xa9\x0e\xe9\xdc\xed\xa1\x61\x2d\x83\xaf\x7d\x39\x72\x51\x5a\x5d\x90\xdf\x5f\xbd\xf8\xf3\x36\xe5\xb3\xc5\x24\x7f\xc9\xbc\xba\x76\xd2\x3a\xef\x76\x2c\x89\xb2\xe1\x0d\x07\xbc\x7a\x8f\x90\x6d\xf6\xaa\x26\x0d\xbb\xe5\xce\xce\x35\x7a\xc5\x32\x7c\xcc\x39\x4a\x3a\x66\xc5\x40\xed\xfb\xc6\x12\xc0\x27\xc6\x03\x03\xf2\x6a\x31\x1f\xb3\xfc\x91\x3a\x51\x01\x8f\x42\x46\xa6\xf2\x4f\x0e\x8c\x2c\x5b\x4e\x0b\xf0\x04\x00\x77\x92\x95\x27\x00\x66\xb7\x51\x3e\x05\x16\x83\x38\x85\xed\xed\xd7\xb7\x2f\x5f\x40\x47\x8c\x8e\x11\x68\xe2\xe1\x5a\x18\xb0\x49\x30\xeb\x98\x71\x75\x6d\xd4\x16\x06\x38\x21\x1c\x38\xbc\x4c\xba\x52\x1b\xc7\x67\x6e\x78\x06\x3e\xd6\x4d\x7c\xef\x30\x1c\x3b\x67\xef\xd7\xe3\x95\xcd\x35\xe6\x12\x15\x5e\x4e\x60\xd7\x9a\x33\xad\x5c\x3d\x5f\x39\xf3\x0e\x48\x79\xfe\x2e\x0c\xe3\x02\x05\x01\xbc\x62\x86\xf8\x26\xb0\x7e\x60\xa9\xaa\xc6\x29\x44\x6c\x34\xc1\xda\xca\x77\xb3\xa3\xff\xa1\xc5\x0b\x2a\x22\x30\x32\xf2\x2a\xa7\x88\x9c\x2a\x49\x13\xd0\xa2\x73\xaa\x82\xcf\xbe\x5c\x8c\x8e\x77\x03\xbd\x7d\xe3\xe2\x28\x65\x87\x5d\x2a\x61\x7a\x45\x08\x44\x6c\x56\x8a\xc8\x8f\x01\x00\x6b\x0f\xfe\x98\x05\x05\x45\xed\x51\x34\x83\xff\xb2\xc0\x50\x34\xce\x62\xd7\x42\x3c\x14\x18\x76\xb9\x44\x88\x1c\xf0\xc8\xb3\xfc\x73\x94\xaa\x07\x5b\x71\x9a\xd6\x29\xb9\xe1\xb1\xd1\x37\xa5\x68\xfe\xa5\xcf\xd3\x04\x1c\x35\x88\xa9\xa5\xd7\x59\xe0\xd1\x08\x8e\x30\xaa\x05\xd2\xb2\xcf\x45\x2c\xc2\x19\x91\x73\xae\x9d\xad\x5a\x48\xbd\x05\xe5\xcf\x77\x44\xb9\x12\xe3\x5a\x69\x29\x15\xd8\x52\x5e\xd2\x96\x7c\x0b\x71\xdd\x09\xe1\x2d\xa4\xf5\x92\x47\x6c\x4b\x59\x81\x8d\xfb\x26\x9a\x75\x17\x74\xb7\x90\xd4\x9b\xd7\x2f\xb7\xd6\xaa\x6c\x0e\xb6\x48\x33\x58\x32\x6c\x5c\x08\xc0\xaa\x08\xdc\x86\x1a\xdd\x0a\xa5\x2d\xa4\xf1\x5a\x1f\xfb\x6c\x29\x11\x9a\x4e\xf1\xa8\xef\x6b\x88\xe4\x96\x48\x6d\x21\x93\xb7\x64\xff\x7f\xfe\xfb\xe4\xc1\x96\x22\xc1\xdd\xcd\xef\xdd\x64\xb4\x58\x4c\x20\x38\x1b\x79\x27\xde\x9d\x88\xe7\x0e\xc8\x7e\x1d\x51\x1d\x7c\x1b\x51\xdd\x2e\xd9\xaf\x23\xaa\xfe\xb7\x11\xd5\xed\x92\xfd\x3a\xa2\x7a\xfc\x6d\x44\x75\xbb\x64\x1d\x67\xb3\x74\x30\xa1\x5c\x7a\xb7\x26\xe0\x7f\x41\x0b\x41\x14\x2b\x64\x91\x45\xe0\x5c\x0f\x4c\x95\x7b\x1c\xa7\xd9\x55\x6d\x14\x79\x79\x14\xb7\x57\x45\x91\x43\x2b\x7b\x09\x6f\x47\xac\x78\xb0\x0c\x07\xad\x68\xd0\x04\x83\x6e\x6c\x04\xf1\x55\x3d\x16\x94\x0d\x06\xe0\xd2\xe7\xe0\xe5\xef\x61\x08\x86\x11\xd8\x8a\x00\xec\x9e\xef\x3f\xbc\xbd\xff\xf6\x3a\x18\x57\x28\x2e\x09\x46\x0b\x10\xfd\x16\x7b\x2a\x7d\xb0\x21\xd4\x53\xb3\xe1\x87\xf3\xa8\xe5\xc6\xb2\x3d\xe8\x53\x69\x32\x38\x66\x13\x80\x29\x88\xba\xca\x57\xc2\xa8\x2e\x86\x64\x38\x75\x7d\x23\xc6\xc3\x56\x37\x6f\xfc\x6a\x8d\xde\x29\x65\x7b\x1f\xc8\xa0\x6f\xb2\x48\x9a\x63\xc1\xfb\x87\xfd\x27\x47\x93\xc3\x2a\xac\xbc\x3f\x99\x4c\x9e\xae\xc4\x6f\x9f\x81\x5b\x03\xb1\xc0\x30\x9a\x33\x7f\x75\x98\x6a\x39\x6f\x16\x8e\x9b\xdd\xed\x58\x35\x43\x38\xed\xab\x28\xb8\xd6\x65\x02\x33\xe0\x9f\xb1\xe5\x98\x63\xc4\x37\xe1\x20\x7d\x25\xcd\xd5\xe1\xaf\x83\xd2\x89\xac\xc9\x3d\x95\x19\x44\xd5\xc5\xeb\x4d\x8a\x4a\xbe\x5f\x44\xd3\xa5\xb1\x9a\xfc\x17\x8b\xbf\x76\x87\x6b\x8e\x85\x64\xe1\xea\x4a\x5e\x40\x5c\x5f\x57\xc7\x45\x2b\x8f\xc0\xdc\x5b\x01\xad\xd4\x66\x79\x84\x22\xf1\xaf\xae\xe4\x6a\xba\xbe\x86\x12\xac\x13\xf8\x55\x8d\x3c\x02\xa3\x37\x29\x33\x57\x57\xb2\x80\x67\xd3\xf3\x38\xc5\x6f\xf8\x91\x5f\xf4\x52\x7e\xd1\x4b\xfc\x32\xe6\x72\xe4\xbd\x7b\x67\xca\xef\xdf\x43\xe0\x1f\xc5\xca\xda\x0d\x9d\x1c\x34\x87\x35\xe7\x4e\x19\x11\xe8\x55\xf1\xbe\x1d\x6d\x79\x2f\xa2\x96\x8d\x47\x40\xa0\x82\x66\x72\x6c\xa7\x2c\x8d\xf4\x69\xb9\xb7\x6e\xbc\xfa\x72\xbb\x31\xbb\xc6\x3a\x75\xb3\x8e\xdf\x6c\x7b\xda\x69\xb1\xa7\xa4\x6e\x81\xd6\x59\x53\x3c\x3e\x83\x76\x03\x42\x2c\x80\x19\xaf\x3e\x89\x53\xc7\x6f\x3f\x71\x9e\x30\x9a\x3e\x52\x87\x6a\xf2\x08\x6e\x42\x93\x82\x55\xa7\x70\x52\x76\xdb\xf5\xd1\x40\x7d\xb4\x87\xcc\xe1\x65\xa1\xf5\x49\x2f\xad\x4f\xec\xe2\x8a\xd9\x3a\xf4\x53\xe7\x7d\x78\x54\x57\x54\xf5\xfb\x4a\xf0\x03\xcc\x26\x0f\x64\xf1\x91\x4c\xac\x52\x00\x9c\x0a\x5c\xd9\x7a\x43\xd9\xf2\x48\x4f\x4b\x59\x75\xf9\xdb\xd4\xff\x6d\xea\xff\x01\x4c\xfd\x86\x26\xbd\x31\x17\x67\xad\x49\xb4\xf7\x91\xbb\xb0\x77\x7a\x25\xb6\x9b\x3c\xcb\xc4\x49\x46\x1c\x88\xb6\x82\x15\x64\x03\x6b\x83\x97\x0b\xd0\x8f\x8c\x74\xb2\x83\x22\x63\x19\x9b\xca\x14\x7a\xb0\x0f\x78\xea\x9b\x9f\x31\xcc\x6d\x04\x4b\x25\x8b\x12\x28\x8d\x12\x21\xc6\x2c\x49\x03\xf9\x54\xdb\xb3\x70\x46\x53\x04\x2a\xcb\xb6\xff\xef\xa7\xbf\xbf\x0a\x0a\xc9\x65\x3c\x59\xee\x43\xc5\x83\xcd\x8c\xd8\xaa\x8b\x09\xe4\xee\x2b\x9b\xb0\xdb\xbf\x46\xa8\x29\xf6\xca\x1b\x82\x6a\x4f\x26\x27\xda\xa5\x6f\xb8\x2c\x58\x73\x5b\xe0\x5c\x17\xb8\x47\xf9\xd5\x85\xc1\x0d\x0d\x35\xf7\x9a\x25\x07\x1e\x2a\x11\x4c\x47\x0c\x66\x62\xc6\x31\x85\x58\xaa\x27\x70\x0b\x10\xf2\x2b\x42\xf4\x9d\x9f\x83\xc7\x89\x33\x3b\xdb\xd1\xca\x41\x9d\xf9\xdc\x93\x89\x8f\x8a\x16\x9e\xff\x92\x37\x0a\xbc\x8e\x9a\x1d\x54\x3a\x51\x65\x75\x8f\x5c\xde\x6f\x0f\x5d\x8f\xa4\xc5\x2f\xd2\xa7\xdf\xda\xa2\xb4\x31\xad\x1a\x69\xff\xb0\xa7\xfd\xc2\x9e\xf6\x08\x1f\x1e\x78\xd5\x58\xdc\x19\x56\x1d\xdd\x51\x01\xc6\x1a\xcb\xe6\x36\x7d\x43\x8e\xb3\x79\xe5\x25\xb6\x72\x8c\x8d\x5a\xf8\x7d\xdc\xeb\xf5\x4a\x8e\xf1\x4c\xb9\x95\xc1\x8d\xf8\x91\xa7\x9c\xbe\x49\xc2\x5b\xcf\x5a\xbd\x7d\x8d\x4b\xff\xc7\x52\xae\x3f\x56\x6c\x9a\xc3\x5e\xb2\x1f\x31\x3c\x55\x2f\x1e\xb4\x30\x6d\x99\x73\xfb\x4c\xa0\xed\x92\xb8\xc9\xaa\x5b\x96\x69\xe3\x33\x81\x1d\x6f\x65\xcd\xad\x7b\x6b\x8a\x71\x99\x98\xd9\x96\x62\x79\x87\x56\xed\xf6\x7c\x39\x40\xa6\xde\x16\x98\x3c\x69\x27\xc9\x74\x75\x92\x74\xe5\x75\x68\x0c\x9d\x32\xc9\x3a\x96\x17\x7d\xbe\xce\xb5\x5e\x2f\x22\xf3\xa8\x6a\x95\x3c\xcb\x04\xe8\x83\xec\x92\x14\x1c\xf6\x15\x99\xf1\xa9\xc6\xd0\xf8\x28\x11\xb1\xb9\xa9\xd7\xab\x38\x71\x79\x90\xa4\x6e\x52\xda\x66\x3f\x29\xc8\x69\xa9\x45\x3b\xec\x23\x55\x9a\xc6\x70\x76\x78\xfc\x9b\x94\xb1\x42\x48\x7e\x89\x13\x36\xec\x02\x54\xd7\xcb\xf4\xea\x32\x49\x2e\x4b\x38\x8d\xb4\x06\xfb\xf8\x74\x45\x3e\xae\x74\xb2\xeb\xdc\xc4\x3b\xd2\xe6\xdb\x61\x52\x96\xa4\xab\xd2\xaf\x11\x97\x8e\x62\xa1\x24\x39\xf2\xcc\x53\x17\x45\x54\xb5\x00\x01\xb3\x51\xbf\x07\x66\xa0\x7a\x6f\xb2\x8d\xa7\x18\x26\x71\x78\x36\xf2\x3e\xac\x8d\x88\x3b\x6e\x3a\x4b\x99\xd4\xed\xc2\xab\x0d\xc8\x15\x69\x95\xc5\x56\x97\x24\x30\x54\x3e\x4a\xf5\xea\x82\x0d\x75\xb7\x9a\x54\x6b\x8a\xac\xdf\xe6\x48\xa3\xac\x3b\x48\xe3\xe9\x83\x82\x16\xa3\x43\xa2\x7f\x68\x72\x41\x97\x80\x0a\x30\x0b\x5f\xf6\x21\xb0\x82\x42\x36\xe3\x09\x28\xe0\xc8\x9b\x32\x41\x7a\xc1\x93\xe0\xe0\xc8\x2b\x4f\x2d\x0c\x8d\x6d\xe5\xaa\xc5\x2a\x53\x57\xe4\x50\x8c\x00\x74\x76\xcc\x1a\x41\x6f\x4f\x05\x9f\x71\x8b\x1b\x44\xde\x20\xb4\x91\x46\xf3\xfc\xdd\xde\xbe\x51\x19\xf5\x86\xdd\x43\xfb\xe0\x3a\xb3\x48\xbb\xfc\x26\xb3\xc8\xf3\xae\xb5\xfb\xaf\x95\xb2\xd1\xef\x87\x35\x8e\xcf\x8a\xa4\xc7\xfe\x4f\x41\xb9\x46\x8c\xd3\xae\x84\xae\x96\x30\xae\xe0\x7d\x6c\x2f\x9b\x15\x01\x66\x83\xee\xf7\x5a\xbc\xf6\x86\xb3\xec\xda\x01\xb4\x14\x31\x61\xf8\xce\xa6\x58\xf5\x9e\xa0\xf5\x84\xda\xce\x1c\x6b\xdd\xf2\x74\x36\x98\x7c\x3f\x6c\xd9\xd6\x23\xf9\xfc\xc0\x58\xe7\xbe\xb1\x99\x25\x6c\x4d\x76\xd2\xea\x80\x58\xe6\xd9\xeb\xe6\x56\x2e\x53\x33\x2e\xc7\x3e\x97\x2a\xab\x19\xae\x0d\x55\x09\xab\x5a\x55\xaa\xd5\xf1\x36\x41\xa6\x23\xb4\xdd\x55\x4a\xef\x9a\xf2\x95\x94\xa5\x57\x7c\xfc\xb1\x52\xad\x9c\x7d\x02\xcd\xaa\xc5\x79\xd0\x02\xd6\xd8\x27\x79\x3c\x21\xf5\x2f\xc7\x9c\x36\x84\x42\xc9\x00\x53\x74\xe2\xad\x4c\x37\xf9\x2f\x0e\x9c\xca\x4c\x3a\x9e\xef\x7f\x77\xff\x86\x48\xbe\x93\xf8\xb0\x9b\x93\xf1\xe6\x0d\x67\xfd\x63\xc9\x23\xd8\xcb\xfe\xb1\xb4\xa5\xde\x43\xe8\xf2\xd0\x53\x2b\x14\xcd\xac\xfc\x1f\x6a\xdf\x30\x20\x53\x08\xbb\x25\x30\x6a\x5a\x7a\x25\x01\x90\x57\xba\xbf\xb1\xd2\xdf\xd5\x05\x8e\xce\x1e\x6d\xbe\xc0\x69\xcc\x32\x5d\xb1\x3c\xfe\x8f\x7b\x7b\x01\x18\x21\x51\xbd\x38\xd3\x4b\xf8\x50\xbb\x47\xe6\x64\xcf\xb8\x4b\x32\xff\x0e\xb7\x73\xf0\xba\xd4\x13\x0d\xe9\xd7\x95\x2f\x8f\x82\x3e\x9b\x6f\xeb\x16\x59\x41\xd8\x4b\x25\xd0\x2f\xf7\x8e\xde\xc2\x74\xb0\x9c\x8a\x45\xee\xf8\x45\x88\xa3\x5a\xf6\xa9\x15\x1a\x89\xaa\x43\x79\x80\x5f\x49\x46\x5f\xa5\x5a\x89\xb0\x26\xdf\xf5\x06\xe5\x67\xe9\x74\x91\xc0\x32\xcb\xca\x9c\x97\x4d\xc8\x57\x31\xd6\x97\xd0\x56\x59\x47\x9b\xd0\x53\x29\x40\x1b\xd3\xba\xdd\x6d\x57\x2f\x9b\xcd\x43\xb6\xbd\x8b\x38\x8d\xf8\x45\xc0\x53\xdc\x30\x31\xdb\xd6\xde\x6b\xaf\xb7\xd9\x1a\x55\x9a\xf7\xba\x85\x5d\x26\x83\xff\x3f\x5e\xd7\xe1\x8c\xe6\x62\xfd\x92\x86\xa6\x13\x9c\xff\xb2\xb9\xf5\x3e\xf5\x49\xad\x8b\x41\x51\x7f\x57\xb8\x6d\x18\xf4\x07\x1b\x9f\xd0\xf9\x2e\xab\x1c\x9f\xea\x1b\x8d\x56\x93\x04\x8a\x0b\xae\x2f\x20\x93\x4f\xf7\xf5\x43\xfc\x3f\x64\x15\x01\x65\x36\x8f\xf1\x6f\x4d\x79\x15\x55\x6f\x6b\xc5\xac\x69\xe6\xff\x06\x00\x00\xff\xff\x3e\x17\x58\x2a\x16\x49\x00\x00")

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

	info := bindataFileInfo{name: "fcs-lpc-motor.html", size: 18710, mode: os.FileMode(420), modTime: time.Unix(1466153351, 0)}
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

	info := bindataFileInfo{name: "index.html", size: 7160, mode: os.FileMode(420), modTime: time.Unix(1466153595, 0)}
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
