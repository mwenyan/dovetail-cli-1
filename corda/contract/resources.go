// Code generated by go-bindata.
// sources:
// resources/kotlin.concept.template
// resources/kotlin.contract.template
// resources/kotlin.contractimpl.template
// resources/kotlin.pom.xml
// resources/kotlin.state.template
// DO NOT EDIT!

package contract

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

var _resourcesKotlinConceptTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\x5d\x6f\xd3\x30\x14\x7d\xcf\xaf\xb8\x8d\xa6\x29\xd5\xa6\x08\x5e\x2b\xc2\x28\x65\x42\x08\x54\x90\xb6\x3f\xe0\xa5\xb7\xc3\xc3\xb5\x33\xfb\xa6\x52\x70\xfd\xdf\x91\x9d\xb4\xcd\xe7\x04\x48\xbc\x91\x87\x7c\xf9\xe4\xdc\xe3\x73\xcf\x4d\xc1\xf2\x1f\xec\x11\xc1\xda\x74\x7d\xe7\x5c\x14\xf1\x5d\xa1\x34\x81\x44\x4a\x73\xa5\x37\xcc\x9f\x31\x35\xa8\x39\x13\xfc\x27\x23\xae\x64\xba\xf2\x0b\x77\xc7\x57\x0f\x02\x27\xbe\xe2\x1b\x94\xc4\xa9\x4a\x97\x0f\x86\x34\xcb\xe9\x1b\xd3\x54\x1d\xc1\x4f\x6c\xcf\xd2\x92\xb8\x48\x57\xa5\xd6\x28\xf3\x2a\x8a\xde\x0d\xa9\xad\xdd\xe0\x96\x4b\x84\x78\x49\xa4\x3f\x60\x2e\x62\x70\x6e\xcf\x34\x58\x0b\xe9\x9a\xed\x10\x9c\x5b\xf8\x87\x8b\x03\x7c\x44\xfa\xac\x48\x70\x79\x5f\x15\xe8\x9c\xb5\x7c\x0b\xf8\x0c\xe9\x27\xf3\xb5\xf0\xd2\x99\x00\xd2\x25\x3a\x77\x63\x2d\xca\x8d\x47\x84\x4b\xa4\x0a\x94\x90\x0b\x66\x8c\xb7\x62\xe5\x6f\x9c\x4b\xac\xd5\x4c\x3e\x22\x5c\xf0\x6b\xb8\x60\x44\x1a\x16\x19\xa4\x5e\x07\x7f\x28\x09\x4d\x53\x41\x7a\x04\xbc\xf2\x4f\xf1\x75\x7c\x22\xb5\x96\x70\x57\x08\x46\x1d\xed\x81\xe7\x84\x99\x83\x8d\x00\x00\xc2\x49\xed\x51\x6b\xbe\x41\xd8\x96\x12\x48\xdd\x91\xe6\xf2\x31\x99\xc3\x02\xea\xdb\x1a\x1a\x0e\xbf\xff\x27\xa3\xe4\xa2\x5e\x81\x0c\x62\x1b\x9f\x97\x4f\xc2\x47\x45\xb7\x71\xb5\x41\x01\xd7\x76\x69\xcb\x84\xc1\x36\xd2\x17\x83\xac\xbe\x5c\x05\xb7\xfd\x27\x07\xb8\x57\xb5\xe1\x8d\x0e\xe7\xe0\x0a\xe2\xeb\x8e\x14\xec\x51\xf1\x6d\x62\x6d\x5d\xd1\xb7\xcf\x39\x98\x65\x20\x4b\x21\x82\x9a\xa3\x94\xa5\xd6\xac\x72\x0e\x2e\x2f\xa1\x8b\x9e\xcd\x52\x6e\xd6\x8a\x6e\x77\x05\x55\xc9\xfc\x68\x64\xcb\x9c\x49\xb9\xf0\x1b\x7a\x3b\xee\x74\x32\x32\x46\x9e\x6e\xb4\x2a\xbe\x30\x43\xc9\xeb\xf9\x94\x59\xb1\x6b\xd1\x6b\xa4\x52\xcb\xb0\x14\xd5\xe5\xa2\x7e\xe7\xf1\xb9\x64\xc2\x24\x8a\xbe\xa3\x5e\xc0\x52\x56\x37\xf3\x05\xbc\x57\x4a\x20\x93\x00\xb6\x63\x64\x00\x01\x6f\xa7\x76\x0e\x5d\x2b\xf6\x3e\xf4\x0a\x32\xa8\xb1\xcc\x40\xd2\x02\xf7\xa0\x9e\xeb\xd6\xd7\x3f\x57\xcc\xc2\xc8\x44\x83\x68\x4d\xce\x44\x87\x72\x7a\x04\x3b\xb0\x61\x26\xb2\x3a\x13\xbe\xff\xa4\xd2\xf1\xc5\xae\xfa\xc0\x53\xab\xef\x8b\xf6\x87\x8f\xa1\xaf\xf3\x47\x85\x9a\x64\xce\xe1\x70\x80\x89\xd4\xbe\xac\x70\x28\xb1\x49\x40\x98\xb0\xa1\x40\x3b\xc0\xf7\x67\x34\x0c\xc6\x98\x87\x8d\x8f\xb3\xc1\xb8\xe4\x4a\x12\xe3\xd2\x2c\x85\x48\x82\xd2\xa3\x48\x66\x60\xa5\x84\xc0\xdc\x37\xe6\xcd\x79\xa8\x3b\xbf\xd1\xb5\x6a\x46\xf1\xed\xc8\x66\x5e\xdc\x10\x8c\x8e\xff\x49\x29\x8c\x48\x6d\xb2\xdf\x56\xf9\x97\x55\xbb\x23\xdb\xb6\x78\x94\xee\x85\xe0\xf4\x03\x3d\xb1\xa1\x7f\xd0\xa7\x4e\xa2\xfe\x37\xeb\x2c\x67\xba\x59\x7d\xb2\x31\xf2\x46\x4d\xc3\x72\x5a\x72\x63\x13\x38\xaa\xdc\x35\xbf\x6e\xf7\x2b\x00\x00\xff\xff\xbc\x48\xb9\x08\x3f\x09\x00\x00")

func resourcesKotlinConceptTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinConceptTemplate,
		"resources/kotlin.concept.template",
	)
}

func resourcesKotlinConceptTemplate() (*asset, error) {
	bytes, err := resourcesKotlinConceptTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.concept.template", size: 2367, mode: os.FileMode(420), modTime: time.Unix(1542209137, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinContractTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xc1\x6e\xdb\x38\x10\xbd\xfb\x2b\x06\x42\x0e\x32\x60\x10\x8b\x05\x76\x0f\x5e\x04\x1b\xc3\x41\x5b\xa7\x45\x5a\xc4\x06\x7a\x0c\x68\x6a\x22\x4f\x2a\x91\x02\x39\x72\xe3\xaa\xfc\xf7\x82\x94\xe4\xc8\xb1\x5c\xa0\x39\x84\x34\x35\xf3\xe6\xbd\xc7\x19\x56\x52\x7d\x93\x39\x42\xd3\x88\xfb\xb5\xf7\x93\x09\x95\x95\xb1\x0c\xca\x94\x82\x69\xab\x8c\xc8\xcc\x1e\x59\x52\x21\x94\xd1\x2c\x49\xa3\x15\xca\xd8\x4c\x8a\x65\xf8\xbf\x34\x65\x29\x75\x76\x2b\x59\x7e\x25\xde\x85\xf5\x4f\x21\xde\x15\xe6\xfb\xd2\x68\xb6\x52\x71\x9f\xab\x91\xbb\x10\x65\x2c\xc6\xbc\xf0\xd9\x89\x3e\xf0\xbf\x0b\x91\x6c\xa5\x76\x52\x31\x19\xed\xc4\x27\xcc\x72\xb4\x9b\xd7\xa3\x3e\xe9\x59\xee\xa5\x20\x23\x56\xba\xaa\x79\xcd\x16\x65\x79\x01\xcf\xa1\x25\x59\xd0\x0f\x19\xd2\x5b\xbe\xeb\xfe\x68\x5b\xe0\x09\x60\xcd\x54\x88\x65\x6d\x2d\x6a\x75\xb8\x80\x47\x19\x6a\x26\x3e\x88\xc5\xd6\x45\x21\x5f\xa4\xe5\xc3\x6f\x2d\xb3\x28\x6c\xad\x99\xca\xa0\x8e\xf2\x1c\xad\x58\x6d\xda\xcd\x64\x72\x73\x4e\xa9\xaa\xb7\x05\x29\x50\x85\x74\x2e\xdc\x6b\x6f\xd9\x32\x1c\x78\x0f\x73\x38\xb3\x3d\x9d\xce\xa0\xdf\x43\x33\x01\x80\x40\xa5\x92\x9a\x8c\x06\xb3\x7d\xc6\xe3\x71\xf8\xbb\xb9\xdb\x97\x6b\x96\x4c\xea\x78\xb4\x97\x05\x34\xcd\xd5\xdb\x5a\x8f\xcb\xcf\xf7\x9b\x87\xc5\x72\xf3\xb8\xba\x85\x6b\x48\x42\x48\x68\x33\x31\x12\x9b\x44\x30\x0f\x71\x69\x1a\x2b\x75\x8e\x70\xc5\x2f\x1a\xe6\xd7\x20\xba\x3e\x73\xe0\x7d\x0c\x18\xd1\x1d\x59\x77\xa2\x63\xa2\xb8\x97\x25\x82\xf7\xe9\x11\x8d\x66\x70\x25\x99\x6d\x80\x8c\x11\x0b\x66\x4b\xdb\x9a\xd1\x79\xdf\x34\xf4\x04\x3a\x44\xc1\x5f\xe1\x57\x32\x4b\xc2\x82\x3a\xf3\xbe\x17\x18\x92\x23\xac\xf7\xf3\x58\x26\xa2\xfd\x84\xf7\xc8\x1f\x0d\x17\xa4\x37\x87\x2a\x94\x0c\x0a\x62\xe2\xb4\xb7\x7b\x64\x50\xd2\xe9\xc0\x54\xd2\x34\xf4\xf8\xc4\x84\x8b\x94\x4f\xc2\xab\x9a\x23\x6a\x72\xca\x33\x99\xbd\x21\x3e\xed\xa8\x8d\x27\xab\x96\x67\x32\x8b\xb7\x15\x5d\xec\x2e\xac\x77\xd4\xfb\x64\x7a\xcc\x6d\x51\x7c\x77\x69\xaf\xb8\x66\x8f\xd6\x52\x86\xf0\x54\x6b\xc8\x91\x1f\xd0\x99\xda\x2a\xfc\x20\xdd\x2e\x9d\xce\x61\xcd\x96\x74\x3e\x10\x6c\x91\x6b\xab\x43\x51\x10\xa1\x35\x5b\x53\xc3\x6e\xbd\x93\x7f\xff\xf3\x2f\x1c\x3b\x64\x32\x5a\x60\x30\xe5\x77\xce\xe8\x50\x63\x30\xdd\xe7\x85\x78\x47\x4e\x84\xb1\x8d\xed\x27\x06\x14\x17\xae\xcd\x49\x93\x93\xc7\xe4\xd9\x19\xdd\x09\x1f\xa3\xb0\x47\x4b\x4f\x87\x94\x5f\xe6\x70\xf6\xec\x0c\xef\xb9\x8d\x1b\x7c\x4c\xf9\xe5\x04\x75\x44\x59\x9c\xf5\x54\xcb\x12\x7b\xe3\xa6\xf3\xfe\x09\xf8\xff\x5c\x5a\xd3\x88\x45\x55\x75\x63\xb5\x2a\xab\x42\xbc\x41\xe9\xeb\xf9\x5f\x01\x00\x00\xff\xff\x44\xb5\x7d\x52\xfe\x05\x00\x00")

func resourcesKotlinContractTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinContractTemplate,
		"resources/kotlin.contract.template",
	)
}

func resourcesKotlinContractTemplate() (*asset, error) {
	bytes, err := resourcesKotlinContractTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.contract.template", size: 1534, mode: os.FileMode(420), modTime: time.Unix(1566930258, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinContractimplTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8e\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\x05\x93\xc9\x5e\xf8\x01\x06\x1e\x1e\x8c\x4e\x59\x32\x34\xdd\x8a\x0e\x8c\xca\x0a\x4a\x2d\x89\xa0\xe8\x2c\x86\xfe\xbd\xb0\x9b\xa0\x45\xa7\xa2\x9c\x78\x81\x7b\x70\x8f\x90\x7f\xa7\xc0\xb0\x2c\x78\x3c\xb5\xe6\x5c\x4c\x52\xd4\xc0\x97\x84\x16\xcf\xbe\xe0\x6b\xb9\xb2\x51\x9c\xd0\x17\x65\xd4\x39\x5b\x4c\x8c\xbe\x24\x89\x13\x6b\xc5\x51\xe4\xe1\x16\xfe\x04\xff\x16\x32\x8d\x21\xb0\xe2\xe1\xe9\xf3\x71\xae\x9c\x2f\xec\x6d\x55\x1f\x45\x8e\x94\xb8\xb5\x43\x92\x09\x16\xe7\x00\x00\xae\x34\x01\x89\xc0\x3f\xf8\xa6\x78\xdf\x1e\x45\xba\x9f\xe0\x30\xf8\x89\x6a\xc5\x0b\x5d\x09\x03\xdb\x23\xd7\x32\xab\xe7\xb1\x9e\x4c\x99\x52\xb7\x37\xa5\x5c\xc9\x5b\x2c\xb9\xe2\xa5\x96\xbc\xef\xfb\x6d\xeb\x6d\xce\x10\xd8\x6e\x6a\x5d\xa6\xc4\x03\x9c\x4c\x63\x0e\xfd\x00\x77\xe5\xff\xb0\x6c\xed\xf5\x94\x6d\xd6\xbc\x0a\xe2\x17\x58\xbb\x7e\xb7\x7b\x5e\xe9\x97\xad\xd8\x5c\xfb\x08\x00\x00\xff\xff\xe5\xa1\xe2\xe5\xa1\x01\x00\x00")

func resourcesKotlinContractimplTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinContractimplTemplate,
		"resources/kotlin.contractimpl.template",
	)
}

func resourcesKotlinContractimplTemplate() (*asset, error) {
	bytes, err := resourcesKotlinContractimplTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.contractimpl.template", size: 417, mode: os.FileMode(420), modTime: time.Unix(1566930225, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinPomXml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4d\x6f\xe3\x36\x13\x3e\xdb\xbf\x82\x6f\x90\x1c\xde\x60\x45\xda\x4e\x16\x5d\x18\xac\x2e\x4d\x0f\x5b\xec\xb6\xc5\xa6\x58\x14\xe8\xa1\xa0\xa5\xb1\x4c\x47\x22\x05\x92\xb2\xdd\x06\xf9\xef\x05\x45\x52\x5f\x96\x9c\xdd\x16\x9b\x8b\xc5\xf9\xe2\x33\x9c\x99\x87\x0c\x2d\x95\xdc\x43\x62\xd0\xa9\xc8\x85\xfe\xfe\x6a\x67\x4c\xb9\x26\xa4\x60\x07\x10\x98\x95\x2c\xd9\x01\x96\x2a\x23\xbf\xfe\xf2\x91\xdc\xe3\x05\x5e\x5c\x39\xcb\xf5\x49\xf3\xc6\xfa\x78\x3c\xe2\xe3\x5d\x6d\xb7\x5a\x2c\x96\xe4\xf7\x8f\x1f\x1e\x93\x1d\x14\x2c\xe2\x42\x1b\x26\x12\xb8\x42\x27\xcd\xd7\xba\x16\x7e\x90\x09\x33\x5c\x8a\x2f\xd8\x0c\x4d\x59\x9c\x74\xea\x84\x51\x6d\x87\x4f\x3a\xbd\x8a\xe7\x08\xd1\x42\xa6\x90\x7f\x06\xa5\xb9\x14\x71\xad\xa3\xa4\x27\xb3\x46\xa5\x92\x25\x28\xc3\x41\xdb\x25\x42\x74\xc3\x34\x3c\x70\x15\x13\x03\x45\x49\x49\x58\xce\x67\x33\x9a\x29\x59\x95\xef\xd3\xd8\xff\x52\x12\x04\x73\x84\x66\x94\x29\xc3\xb7\x2c\x31\xef\xd3\xb8\xfd\xa4\xa4\x23\xae\xcd\x0e\x7e\x73\xff\x4b\xc9\xa1\x45\x63\x41\xd7\xe9\x25\xb2\x28\x79\x0e\x0a\x1b\xa6\x32\x30\xf1\x12\xbf\xa3\x64\x5c\x35\xea\xa6\x65\xa5\x12\x18\x75\xf3\x2a\xe7\xf6\x24\x4d\xce\x05\x0e\x10\x96\x78\x85\xbf\x5b\x52\x32\x10\xdb\xdc\xbd\xa8\x09\x93\x33\x91\x55\x2c\x83\xcf\xad\x6b\xe3\x37\x69\x34\x16\x88\x95\xfc\x62\x8c\x8e\x7e\xcc\x7d\x7f\x28\x7e\xeb\x9c\xd1\xb4\x7a\x4e\x49\xbf\xd6\x4d\x39\xaf\x9f\xfd\xd7\x4b\xaf\xa2\xdd\x82\x5e\x3f\xb7\x8b\x97\x61\x4d\x9b\x92\x5e\x3f\xfb\xaf\x97\x5e\x55\x69\x0a\x25\x88\x14\x44\x52\xef\x3c\x9b\xb5\x82\xbf\x5c\x1d\xea\x5a\x84\xad\xa5\xca\xf0\x1e\xcc\x46\x31\x2e\x34\x76\x09\xf5\x80\x79\xfb\x0e\x08\x67\x14\x69\x93\xe6\x7c\x13\xed\xd3\xa7\x77\x43\x8c\xde\xa7\x45\xda\x2f\xf1\xcb\xb0\x0d\x49\x1f\xe2\x10\x72\x67\x18\x04\x18\x9c\x48\x95\xb2\xfe\x38\xf4\xe6\xa1\xd6\x47\x89\x54\x70\x36\x0f\xed\x40\xfc\x71\x8f\x17\x6f\xfe\xdf\x03\x32\xfb\x26\x38\xb6\x5c\x58\x1e\xba\x04\x65\x65\xd9\xe2\x3f\x80\x49\x64\x81\x0d\xdf\x24\x12\xa7\xf2\x00\x86\xf1\xfc\x02\xaa\x60\x12\x79\xf8\xd3\xb0\x16\x78\x89\xef\xdf\xa0\x57\x70\x75\x3b\x6e\x0c\x58\xc9\x72\x59\x70\x21\x73\xb6\xd1\xb8\x00\xa3\x78\xa2\x2f\xc0\xf3\x16\x91\x80\x63\xa4\x20\xe7\xc9\x25\x84\x4b\xbc\xc0\x6f\x5f\x3b\x38\x84\xd0\xcd\x0d\x9c\x0c\x28\xc1\xf2\x9b\x9b\x79\xaf\xdf\xc2\x74\x2a\x28\xa5\xe6\x46\xaa\x96\x9a\x1b\x51\x33\x36\x54\x0b\x56\xea\x9d\x34\xba\xd3\xe4\x20\xd8\x26\x87\x34\xde\xb2\x5c\x03\x25\x61\x19\x5c\xc8\x99\x0f\xe5\x69\x9c\x80\x30\x8a\xe5\x94\xf0\xd6\x52\xb0\x02\xe2\x1f\x9c\x02\x7d\x6a\x36\xa7\xa4\x56\x04\xab\x4a\xe5\xb1\xbd\x9b\xf4\x9a\x10\x8b\x10\x9f\xdd\x50\xb5\x60\x45\x89\xb5\xf4\xe3\xd5\x4f\xa5\xb3\x0e\xe9\x6f\x2a\x9e\xfb\x1b\x23\xe5\x0a\x92\xda\xf4\xfa\xd9\x5f\x47\x2f\x94\xb4\x52\x17\xd2\xf1\xfa\xc3\x88\x2d\xd1\x2a\x21\x05\xe3\x82\x04\x32\x19\xda\xba\x08\xb2\x32\x65\x65\x46\x23\xb8\xbb\xc6\xfb\x93\x24\x67\x5a\x83\xa6\x64\xe8\x11\xca\xe4\xe2\xb7\xe7\x1b\x24\x9d\x22\x8d\x25\xd5\x02\x6d\x42\x9c\xe5\xe9\x0e\xab\x77\x89\x91\xc1\x86\xd4\x0e\x78\xfe\xb3\xad\x91\x27\xc6\xeb\x67\xff\xb4\xc1\x1d\x1a\xef\x48\x5b\x16\x6c\x5d\x5d\xa8\x32\xaf\x32\x2e\x74\x6f\x35\xc1\xd9\xbe\xe0\xae\xfa\xde\xef\x15\xde\x76\xcf\x96\x3d\x53\x91\xb3\x7f\x85\xb5\x57\xf8\x0e\x2f\x07\x4c\x5d\x1b\x24\x52\x6c\x79\x56\xa9\xfa\x2d\xd5\xca\x5f\xa9\xea\x44\xfd\xdc\x99\x8e\x84\xa4\x24\xe4\x3f\x9b\x9d\x1d\xc5\xc4\x95\xe4\x32\xbc\x94\xdd\xbf\xb9\xfb\xbe\xea\x2e\x6b\x79\xe1\x04\x49\x65\xd3\xd1\x7d\x45\x5f\x79\xae\x6b\x28\xc2\x3d\x2a\xba\x14\x31\x30\xca\x24\xcb\x75\xec\x7e\x5b\xf3\x7a\x85\xdc\xef\xd8\xd6\x64\x62\xef\x8e\xa2\x7e\x36\xd8\xbf\xa9\x52\x5b\x95\x90\x47\xa6\x44\x6c\x54\x05\x94\xf8\x05\x42\xf4\x7f\x51\x84\x1e\xb8\xb6\x24\x88\xac\x8c\x8b\x4c\xa3\x28\xea\xfb\x32\x95\x0d\xa0\x39\x61\x1c\xed\xd9\x81\x45\x25\x53\xac\x00\x03\x4a\xdb\x22\x66\xb1\x8b\xfa\x63\xcd\xac\x48\xdb\xfb\xc1\x20\xfb\xb4\x46\x5b\xa9\xd0\x4f\x8f\x9f\xa2\xbb\xc5\x5b\xc4\x84\x90\xa6\xc6\x79\xbe\x1f\x71\x1b\xce\x7c\x56\x63\xed\x86\x7a\x2d\x17\x04\xdf\x66\x04\xf5\x8e\xa5\xf0\x65\x43\x78\x87\x57\xe3\x43\x38\xde\x5e\x53\x8d\x45\xcb\x1d\xd3\x10\x97\x2c\x79\x62\x19\x50\xe2\x96\xf3\x91\x76\x9a\x8f\x34\x59\x5c\x03\xf6\xad\x35\x68\x9a\x11\xaf\x4b\x04\xd1\x9e\xc6\x63\xf8\x6f\xa2\xa3\xe3\x22\xc9\xab\x14\xc6\xda\xd6\xab\x3a\xcf\xe7\xf5\xe0\x99\x1c\x2c\xa6\x7d\xcf\xdf\x49\xeb\xdb\x49\xbf\x46\x31\x44\xd3\x56\xec\x3c\x05\xba\xe5\xb9\xed\xdb\xb3\x60\x4e\x3e\x82\x2d\xc4\x9a\x00\xd7\xa8\xc7\x38\x64\xea\xb0\x1a\x9d\x0d\x4a\xea\xa0\x24\x04\xb5\xbd\x6f\x18\x17\xa0\x48\xfd\x00\x2c\x4b\x72\x7b\x4b\x6e\xed\xf4\x4f\x1c\x5f\x27\x5c\x08\xf2\xe7\xd7\xbb\x8e\x22\x51\x29\x23\x7b\x2d\x05\x49\x41\x83\xe2\x2c\xe7\x7f\x83\xba\x1c\xb4\x51\x9c\x9f\x31\x19\x3d\xe4\x20\x1e\x36\xe9\x14\x09\x4c\x50\x64\x9f\x1e\x83\x2c\xf0\x43\xf8\xb2\x2a\x4a\xfc\x6b\xaa\xfe\x17\xd0\xde\xf8\xf1\x3f\x01\x00\x00\xff\xff\x25\xbe\x99\x37\xec\x10\x00\x00")

func resourcesKotlinPomXmlBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinPomXml,
		"resources/kotlin.pom.xml",
	)
}

func resourcesKotlinPomXml() (*asset, error) {
	bytes, err := resourcesKotlinPomXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.pom.xml", size: 4332, mode: os.FileMode(420), modTime: time.Unix(1568744670, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _resourcesKotlinStateTemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcd\x72\xdb\x36\x10\x3e\x9b\x4f\xb1\xe6\x64\x3c\x64\xa2\xc1\xa4\x57\x35\x4a\xa2\xd8\x4e\xab\xc4\x75\x32\x91\x9a\x4e\x4f\x19\x98\x02\x65\xd8\x14\xc0\x00\x2b\xc5\x2a\x8c\x77\xef\x00\x24\xc5\x5f\x39\x69\x9a\xde\xca\x83\x45\x02\xfb\xf3\x61\x77\xbf\x5d\x38\xa7\xc9\x2d\x5d\x31\x30\x86\x5c\xce\xad\x0d\x02\xbe\xce\xa5\x42\x10\x0c\x49\x22\xd5\x92\xba\xbf\x8c\xf0\x25\x13\xc8\x71\x47\xa6\x57\x1a\x15\x4d\xf0\x3d\x55\xb8\xab\x84\x6f\xe8\x96\x92\x0d\xf2\x8c\x9c\x6e\x94\x62\x22\xd9\x1d\x30\xa3\x99\xe2\x34\xe3\x7f\x51\xe4\x52\x90\x53\xb7\x31\xaf\x96\xae\x32\x56\x69\x25\x72\x4d\x52\xaa\x91\xa9\xbb\x75\x46\x6e\x68\x72\xab\xa5\x20\x54\x08\x89\x85\xe6\x1b\x2d\xc5\x6c\x25\xa4\x62\x07\x1c\x25\x52\x78\x98\x9a\x3c\x0e\x8c\xe1\x29\xb0\xcf\x40\x66\x7a\x9e\x5c\xb3\xe5\x26\x73\xbe\x00\xd5\x86\x59\x7b\x40\xff\x8c\x65\x0c\xd9\x6b\xa9\xce\xde\x7c\xfc\xed\x80\x4c\x9a\xc9\x2f\x9a\xbc\xce\xe4\x97\x0b\xb9\xe2\xc9\x07\x96\xbe\xa6\x09\x4a\xb5\x0b\x8c\x61\x62\x59\xdb\x76\xc7\x41\x7e\x95\x48\xb2\x94\x5b\x86\x94\x67\x1e\x1f\xe5\x82\xa9\xd2\xe0\xe3\x20\x30\xe6\x11\x17\xf9\x06\x61\x3c\x01\x62\x6d\x60\xcc\x92\xa5\x5c\x30\x08\xe5\x96\x29\xc5\x97\x2c\xb4\x76\x7f\x96\x4b\xba\x66\x10\x66\x5c\x30\xaa\x66\xcb\xd0\xda\x4a\x08\xb6\x34\x33\x86\x65\x9a\x59\xbb\xa5\xaa\x84\x52\x21\xaa\x8d\xe6\x54\x31\x81\xa7\x19\xd5\xba\x69\xd7\xe7\xc4\xaf\x42\x18\x5a\x7b\x5a\xc6\x71\x8e\x14\x59\x61\x16\x06\x24\x0f\x46\xff\xc2\x03\xf4\xda\xa1\xb5\x8d\xaf\x3d\xae\x07\x72\x33\x6a\x2c\xb5\x95\x8a\xc3\xbc\xec\x17\xd0\xcb\x57\x2c\x93\x62\xa5\x17\xb2\x42\x1e\x19\x43\xaa\x77\x8f\xd6\xda\xf1\x38\x71\x2f\x71\x70\xe8\xd4\x81\xdf\x77\x9c\x28\x35\x22\x63\x14\x15\x2b\x06\x8f\xf8\x08\x1e\x51\x44\xe5\xb3\x34\x45\x54\xfc\x6a\x83\x4c\x97\x27\x11\x4e\x02\x9e\xba\xaf\x70\x14\xee\x91\xc2\x96\x2a\x30\xc6\x2b\xfa\xcc\x59\x3b\x06\x63\x0a\x4b\xf7\xf0\x0b\xc3\xb7\x12\x33\x2e\x16\xbb\x9c\xd5\x31\x29\xc4\x67\xfa\x5d\xee\x6a\x9e\x66\x65\x54\x5e\xb4\xa3\x10\x8f\x8d\x41\xb6\xce\x33\x8a\x9d\xac\x42\x51\x4f\xd6\x9a\xe0\x5f\x67\xee\x47\x47\xa4\x89\x79\x5f\xde\x85\x1d\x6b\x7f\x78\xa8\xa0\x8a\x15\x38\x5b\x5f\x89\x16\x98\x8a\xbf\x00\x10\x00\x00\x34\xa9\x05\x39\x55\xc8\x13\x9e\x53\x81\x1a\xc6\x70\xc1\x35\x3e\x6b\x35\xc5\xe7\x5e\xc7\x3d\x2f\x1b\x6d\xaa\x5a\x5b\x31\x8c\x62\x03\xfb\x6f\xf7\xf4\xcc\x4e\x60\xaa\x14\xdd\x0d\xd8\x8e\xe2\x96\xa6\x0b\x0c\xf8\xa8\xbc\x6f\xe8\x5b\xdb\x12\x52\x0c\x37\x4a\xb4\x3c\xec\xf7\x0b\xc9\x87\x3b\xa4\x3f\x4a\xab\x21\x76\xc2\x92\x6e\x04\x08\x76\x87\xa5\x2e\x5b\x4e\x13\xe4\x5b\x8e\xbb\x08\xaf\xb9\xf6\x25\xf4\x81\xa5\x63\xa8\xde\x46\x90\xf6\x9b\xe6\x18\x06\x3a\x69\x3c\x86\x9e\xd1\x17\x60\x82\x66\xe4\x50\xf1\x15\x4c\xa0\x4f\xf3\xd9\x3a\xcf\xc8\x8a\xe1\x42\xf1\xd5\x8a\xa9\x28\x34\x86\xf4\xac\x59\x1b\xd6\x21\xe5\x69\x54\x58\x9b\x80\xd8\x64\x59\x6c\x5a\x71\x9c\xef\x34\xb2\x35\x91\x1b\x24\xb9\xe2\x02\x33\x11\x85\xee\x20\x30\x6c\x17\xb8\x06\x21\x11\x52\xb9\x11\xcb\x91\x7b\xbd\xe6\x62\x05\xba\x92\x0c\xe3\xa1\x2c\x39\xbf\x9d\xec\xb8\xc7\xd3\xd7\x74\xaa\x46\x41\x82\x42\xc1\x04\x0a\x42\x57\x23\x25\xaa\xc7\xb1\x2f\x23\x4d\xa8\x76\xa5\xe4\xb3\x01\x54\x43\xab\xa9\xc7\x23\x80\xb0\xe6\xf5\xa7\xc1\x44\x76\xb0\x3a\xb7\x84\x2e\x97\x7b\x97\xef\x95\xcc\x99\xc2\x5d\x14\x0e\x24\x31\x1c\xcc\xf7\x37\x5b\xac\xaa\x26\x1c\x41\xb3\x9c\xbe\x59\xdf\x35\x11\x9d\xd3\x84\x85\x23\x7f\x50\x77\xcd\xe9\x9c\xa7\x1f\x57\xb9\x5e\x53\xb1\xac\x43\xeb\x3f\xcf\x28\xd2\x3f\x38\x5e\xbb\xdf\x0e\x11\x4b\x05\x92\x6f\xd0\xef\x86\xa8\xa8\xd0\x34\x71\xed\x68\xe6\xfa\x4a\x89\xfe\x2b\x5a\xe5\x42\x89\x74\xb8\x58\x7b\xad\x03\xef\xc4\x7c\x9b\x54\x58\x17\xb5\xe3\x39\x53\x5b\x9e\xb0\xc8\xd5\xd4\xa8\x72\xd6\x46\xd0\x7b\x5a\xbb\x8e\x0c\x84\x8b\xad\xbc\x65\x91\x8b\xf0\xa8\xf4\xd5\xb6\xe1\xdb\xd4\xbb\xab\x1b\x96\xe0\x73\x40\xaa\x6f\x5d\x03\xf3\x09\x59\x31\xdc\x27\x64\xaa\x77\x22\x59\xb8\xdd\xa8\x5d\xae\x64\x31\x9d\xbf\xfd\x34\x3f\xfd\xf5\xfc\xec\xf7\x8b\xf3\xb3\xe9\xe9\x62\xf6\x71\xb6\xf8\x33\xfe\xb9\xe5\xc4\x51\xb3\x30\x5d\x70\x13\xee\xef\x0b\x5f\x84\xeb\xf3\x75\x8e\xbb\x28\x8e\x3b\x04\x39\x3a\x6a\x70\xaa\x6d\xcd\x0e\x11\xea\xe8\xa8\x17\x70\xd0\x14\x26\x10\xf5\xd6\xe3\xc2\xb5\x6b\xe6\x4f\x3b\x40\xa1\xe6\xb2\xa6\x1d\xaf\x0d\x62\x07\xd0\xe8\xbe\xc5\x75\x06\xba\x3d\x15\xe5\x1c\x15\x17\xab\xc8\xcd\xad\xe2\xb5\x81\xd8\xd5\xe9\x8d\x96\x62\x5c\xec\xc0\x04\x42\x13\xd6\xdb\xfb\xf9\x3c\x38\x9b\x9b\x72\x07\x46\x68\x4a\xfd\x1d\xb2\x96\x74\xce\x60\x52\xfc\x3c\xd9\xcf\x64\xb8\x87\x85\x2c\x66\x72\x09\xc4\x5a\x78\x02\xe1\xa8\x85\x85\x75\x6c\xf1\x34\x6a\x0f\x79\x38\x2e\xf2\xea\xe1\x54\x58\x7c\xfb\xb2\x16\x4e\x4e\x3a\x57\x82\xe3\x63\xc2\xf5\xa5\xc4\x32\xf3\xd5\x78\x6f\xe7\xf3\xfb\xf1\x42\x2b\x3e\xad\x3b\xe7\x90\x75\xb2\x54\x32\xbf\xa0\x1a\xa3\x9f\xe2\x43\xe1\x0a\x6d\xc3\x7e\x59\x1f\x6e\x2b\x68\x94\x43\x2b\xf9\xec\xf3\x86\x66\x3a\x92\x78\xcd\xd4\x18\xa6\x62\xf7\x22\x1e\xc3\x2b\x29\x33\x46\x05\x34\x2b\x97\xa7\x85\x90\x9b\x37\x75\x1f\xef\x92\xc1\x37\x0a\x09\x13\x28\x64\xa9\x86\xa8\x21\xdc\x6b\x80\x5c\x9f\x3b\xff\xb5\xc7\x89\xbf\x0d\x04\xbd\xea\x3a\x78\xfb\xeb\xdc\x53\xf6\x17\x8c\xce\x15\xad\x4b\xf2\x4e\x55\x54\x6c\x3f\x39\x01\x94\x64\x78\xb3\xdf\xd1\x4a\xf4\x5d\xd0\x50\x0d\x52\x9e\x46\xff\xc8\x51\x59\x9b\xb1\x6b\x3a\x07\xea\xf6\x61\x84\x7d\x88\x65\x09\x78\x92\xf5\x01\x9a\x9e\x7c\x97\xa6\x9e\x1a\x43\x31\x2c\xe3\x78\xdc\x23\x4c\xf9\x2f\xa7\x9e\x66\x59\xe4\x91\x56\x20\xfd\x8d\x20\xcb\x98\x9f\x19\xcf\x1a\x3c\x69\x5d\xb6\x2f\x65\xc9\xc6\xe7\x03\xa7\x79\xf0\x44\x30\xd8\x01\xf6\x50\x61\x00\x6b\x59\xfc\x4d\x98\xdf\xe9\xb5\x4d\xda\x66\x8c\x07\xcd\x3d\x50\x39\xdd\x8a\x3e\x70\xa0\xff\x20\x51\xad\x92\xfa\x3f\x5b\x0d\x38\x87\xb3\xd5\x35\x36\x64\xbc\x44\x53\x5a\xa9\x87\xf3\x10\x07\x07\x91\xdb\xb2\x7b\xdb\xe0\xef\x00\x00\x00\xff\xff\x25\x95\x06\x23\x47\x13\x00\x00")

func resourcesKotlinStateTemplateBytes() ([]byte, error) {
	return bindataRead(
		_resourcesKotlinStateTemplate,
		"resources/kotlin.state.template",
	)
}

func resourcesKotlinStateTemplate() (*asset, error) {
	bytes, err := resourcesKotlinStateTemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "resources/kotlin.state.template", size: 4935, mode: os.FileMode(420), modTime: time.Unix(1565926679, 0)}
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
	"resources/kotlin.concept.template": resourcesKotlinConceptTemplate,
	"resources/kotlin.contract.template": resourcesKotlinContractTemplate,
	"resources/kotlin.contractimpl.template": resourcesKotlinContractimplTemplate,
	"resources/kotlin.pom.xml": resourcesKotlinPomXml,
	"resources/kotlin.state.template": resourcesKotlinStateTemplate,
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
	"resources": &bintree{nil, map[string]*bintree{
		"kotlin.concept.template": &bintree{resourcesKotlinConceptTemplate, map[string]*bintree{}},
		"kotlin.contract.template": &bintree{resourcesKotlinContractTemplate, map[string]*bintree{}},
		"kotlin.contractimpl.template": &bintree{resourcesKotlinContractimplTemplate, map[string]*bintree{}},
		"kotlin.pom.xml": &bintree{resourcesKotlinPomXml, map[string]*bintree{}},
		"kotlin.state.template": &bintree{resourcesKotlinStateTemplate, map[string]*bintree{}},
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

