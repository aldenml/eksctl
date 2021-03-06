// Code generated by go-bindata. DO NOT EDIT.
// sources:
// assets/neuron-device-plugin.yaml (2.261kB)
// assets/vpc-admission-webhook-config.yaml (524B)
// assets/vpc-admission-webhook-csr.yaml (234B)
// assets/vpc-admission-webhook-dep.yaml (1.105kB)
// assets/vpc-admission-webhook.yaml (231B)
// assets/vpc-resource-controller-dep.yaml (1.101kB)
// assets/vpc-resource-controller.yaml (673B)

package addons

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _neuronDevicePluginYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xc1\x6e\xe3\x36\x10\xbd\xfb\x2b\x06\x5e\xe4\x56\x49\x49\xd1\x2d\x16\xea\x29\xd8\xcd\x21\x40\x36\x0d\xea\xb6\x97\xa2\x87\x31\x39\x96\x08\x53\x1c\x96\x33\x72\xec\x7e\x7d\x41\x59\xf1\x4a\x8e\x91\xa6\xe7\xfa\x62\x82\x7c\x6f\xde\xf0\x71\x66\xa0\x0f\xd0\xaa\x46\xa9\xab\x6a\xdb\xaf\x29\x05\x52\x92\xd2\x71\x65\xd9\x48\x65\x38\x18\x8a\x2a\x15\xed\x95\x82\x2d\xbe\x41\x2a\xc3\x5d\xec\x95\x0a\x51\x4e\xd8\x50\x11\x48\x2b\x4b\x3b\x67\xa8\x88\xbe\x6f\x5c\x90\x6a\x81\xd1\xfd\x4e\x49\x1c\x87\x1a\x30\x46\xa9\x76\x37\x8b\xad\x0b\xb6\x86\x2f\x48\x1d\x87\x15\xe9\xa2\x23\x45\x8b\x8a\xf5\x02\x20\x60\x47\x35\x04\xea\x13\x87\x62\x16\xac\xb0\x03\x41\x48\x47\x98\x44\x34\x54\x43\x4e\xa8\x90\x83\x28\x75\x0b\x89\x64\x72\x14\x21\x4f\x46\x39\xe5\x35\x40\x87\x6a\xda\x07\x5c\x93\x97\xe3\xc6\xdb\x32\xb2\x00\xe8\xa3\x45\xa5\x95\x26\x54\x6a\x0e\x47\x96\x1e\x22\xd5\xf0\x0b\x7b\xef\x42\xf3\xdb\x00\x58\x00\x28\x75\xd1\xa3\xd2\x28\x35\xb9\x4a\xfe\x61\x08\xac\xa8\x8e\xc3\x49\x1a\x40\x4c\x4b\xb6\xf7\x94\x4a\xf4\xb1\xc5\x72\xee\xba\x49\x4e\x9d\x41\x5f\x44\xb6\x35\x2c\x97\x23\xcd\xcf\xf2\xff\xf7\x1b\x00\xbc\x98\x31\xe4\xce\x9e\xd2\x79\x1e\x05\x6c\xe9\x50\xc3\xe7\x51\xf0\xd6\x5a\x0e\xf2\x73\xf0\x87\x13\x02\x80\x63\xe6\x71\xaa\xe1\x6e\xef\x44\xe5\x9c\x8c\xcf\x52\x62\x87\x7f\x73\x28\x0d\x77\xd5\x31\x9f\xf7\xf0\x01\x68\xb3\x21\xa3\x35\x3c\xf2\x6a\x34\x64\x3c\xfc\x00\x5f\x31\x6d\x41\x5b\x27\x10\xd9\x02\x0a\x20\xbc\xd8\x02\x68\x6d\xc1\xe1\x27\x78\x6e\x29\x00\x05\x5c\x7b\xb2\xdf\x81\xb6\x74\x0e\x39\x45\x3b\xf9\x0d\x89\x84\xd2\x8e\x24\x2f\xb8\x4f\x86\x04\x36\x9c\xce\x89\x59\x54\x40\x18\xb4\x45\xcd\x91\x0f\x60\xf0\x5b\xb8\x35\x65\xfa\x18\xd3\x02\x6e\x94\x12\x20\x6c\xd0\xf9\x3e\x51\x79\xc2\xad\x88\xde\x6a\x2c\x45\xd9\x4a\x85\xb6\x73\xc1\x89\x52\x2a\x8c\xef\xf3\x7f\xd5\xf4\x98\x30\x28\x91\x2d\x46\x15\x17\x9a\xe2\x54\x16\x98\x9f\x29\x17\x87\x54\xa3\x54\x4c\x8e\x93\xd3\xc3\x67\x8f\x22\x8f\x43\x5d\x2c\x8f\x0d\x51\x04\xb6\x74\xa2\xbe\x94\x12\x6e\x36\x2e\x38\x3d\x4c\x8a\x89\x2d\xdd\xbe\xda\x05\x48\xf4\x57\xef\x12\xd9\x2f\x7d\x72\xa1\x59\x9d\xb2\xb9\x6f\x02\x9f\xb6\xef\xf6\x64\xfa\x5c\x5a\x53\xe6\x31\xe6\x6a\xec\xc3\x5f\x29\x75\x32\x3f\xce\x15\x34\x34\xe6\xdd\x3e\x26\x12\x99\x97\xe6\x14\x35\xd4\xd9\x72\x4d\x7a\xde\x29\x2e\x88\x62\x30\x54\xe4\xce\x5c\x5e\xe0\x4e\xcb\xef\x3e\x5c\x04\xec\xd0\xf7\x74\x51\xf8\x28\xee\xc2\xe6\xa6\xdc\x7b\x4c\x0d\xbd\x8d\xf9\xfe\x3d\xa0\x1f\xdf\x15\xe9\x87\x8b\xa8\xff\x66\x57\x76\xff\x7f\x69\x97\xe1\xa0\xe8\x02\xa5\xd9\xa0\x73\x1d\x36\xb9\x2d\xae\xa4\xb4\xdb\x54\x92\x49\xe5\x95\x94\x57\x52\x5d\x1a\xa0\xf5\x4d\x79\x5d\x7e\xfa\xf8\xf1\xba\xbc\x9e\xda\x34\xc4\x78\xea\xbd\x7f\x62\xef\xcc\xa1\x86\xfb\xcd\x23\xeb\x53\x9e\x29\x41\x27\xb8\xe3\x64\xde\x7e\x92\xe2\xe2\x74\x36\x9a\x26\x60\x21\xd3\x0f\xbd\xcb\x41\x69\xaf\x73\x67\xd1\x7b\x7e\x7e\x4a\x6e\xe7\x3c\x35\x74\x27\x06\xfd\x30\xc3\x6b\xd8\xa0\x97\xb9\x2f\x06\x23\xae\x9d\x77\xea\x5e\xbf\x8f\x4d\x1c\x6b\xf8\x63\x79\xfb\xf0\xb0\xfc\x73\x72\xb6\x63\xdf\x77\xf4\x95\xfb\xa0\x67\x9c\x62\xbc\xc5\x2c\xf5\xb3\xa8\x5d\xe6\x3d\xa1\xb6\x35\x54\x3b\x4c\x95\x77\xeb\x61\xce\xf9\x57\x9f\x01\x8b\xa9\xdc\xec\x5d\xde\x56\x69\x59\x8e\x02\x33\xe5\xf8\x2e\xc9\x7f\x02\x00\x00\xff\xff\xe1\x52\xd7\x9e\xd5\x08\x00\x00")

func neuronDevicePluginYamlBytes() ([]byte, error) {
	return bindataRead(
		_neuronDevicePluginYaml,
		"neuron-device-plugin.yaml",
	)
}

func neuronDevicePluginYaml() (*asset, error) {
	bytes, err := neuronDevicePluginYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "neuron-device-plugin.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xfd, 0xe9, 0x45, 0xd9, 0x7c, 0xcc, 0xad, 0xbf, 0xcc, 0x62, 0xb6, 0x4f, 0x28, 0x25, 0x35, 0x93, 0x98, 0x74, 0xb4, 0x9f, 0x84, 0x4f, 0x12, 0x83, 0x47, 0xa7, 0xa4, 0x4b, 0xee, 0xe1, 0xc, 0xa8}}
	return a, nil
}

var _vpcAdmissionWebhookConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\x4f\x6b\xf3\x30\x0c\xc6\xef\xf9\x14\x22\xf7\xa4\xf4\xf6\xe2\xdb\x4b\x29\x63\x87\xc1\x18\x63\x3b\x8c\x1d\x14\x47\x4d\x45\x62\xcb\x58\x76\x4a\xf7\xe9\x47\xfe\xb4\xac\xb0\xd5\x17\xdb\x7a\xa4\xdf\x63\xc9\x18\xf8\x8d\xa2\xb2\x78\x03\xd8\x3a\xd6\xe9\x18\xa9\x63\x4d\x11\x13\x8b\xaf\xfb\x7f\x5a\xb3\x6c\xc6\x6d\x43\x09\xb7\x45\xcf\xbe\x35\xf0\x94\x13\x26\xf6\xdd\x3b\x35\x47\x91\x7e\x27\xfe\xc0\x5d\x5e\x2a\x0a\x47\x09\x5b\x4c\x68\x0a\x00\x8f\x8e\x0c\x8c\xc1\x56\x57\x7a\x75\x5a\x8a\x2a\x7b\xe8\xd6\x0c\x0d\x68\xc9\x40\x9f\x1b\xaa\xf4\xac\x89\x5c\x01\x30\x60\x43\x83\x4e\x10\x00\x0c\xe1\x0f\x4a\xb1\xee\x73\x62\x75\xcf\xaf\x46\x87\x5f\xe2\xf1\xa4\xb5\x15\x37\x63\xed\xc0\xe4\xd3\xf2\xfa\xc5\x08\x40\x29\x8e\x6c\xe9\x72\xbd\xdb\xc2\x4d\xce\xaf\x4d\x2c\x2b\x60\x3a\x1a\x28\x37\x6e\x1a\x1b\x95\x73\x3c\xe6\x81\xf4\xe2\x52\x81\x04\x5a\xc6\xa7\x06\x3e\xa0\xdc\xbd\xec\xff\xbf\xee\x4b\xf8\xbc\x32\x30\xf0\x43\x94\x1c\x26\xbd\x2c\x6f\xe2\xeb\x0f\xce\xca\xb8\xfd\xa1\x45\x52\xc9\xd1\xd2\xac\x04\x69\x75\xd5\x0e\xc8\x43\x8e\xf4\x2c\x03\xdb\xb3\x81\xc7\xce\x4b\xa4\xe2\x3b\x00\x00\xff\xff\x49\xee\x9e\x02\x0c\x02\x00\x00")

func vpcAdmissionWebhookConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookConfigYaml,
		"vpc-admission-webhook-config.yaml",
	)
}

func vpcAdmissionWebhookConfigYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7d, 0x11, 0x23, 0x17, 0x2f, 0xdc, 0x4e, 0x18, 0xaa, 0xc8, 0x66, 0xee, 0xf3, 0xc1, 0x85, 0x63, 0xb1, 0xe3, 0x53, 0x57, 0x80, 0x96, 0xe6, 0x54, 0x26, 0x46, 0x6b, 0x3f, 0x17, 0x20, 0x31, 0x8}}
	return a, nil
}

var _vpcAdmissionWebhookCsrYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8d\x4b\x4e\xc4\x30\x10\x44\xf7\x3e\x45\x5f\x20\x41\xb3\x43\xde\x72\x03\x90\xd8\x77\xec\xc2\x69\x19\x7f\x70\xb7\x83\xe6\xf6\x28\x13\xa4\xd9\x95\xaa\x9e\x5e\x71\x97\x4f\x0c\x95\x56\x3d\x05\x0c\x93\x2f\x09\x6c\xd0\x35\xbf\xea\x2a\xed\xe5\xb8\x6d\x30\xbe\xb9\x2c\x35\x7a\x7a\x7b\x12\x1f\x92\xaa\xd4\xf4\x8e\x9f\x09\x35\x57\x60\x1c\xd9\xd8\x3b\xa2\xca\x05\x9e\x8e\x1e\x16\x8e\x45\xf4\x94\x2f\xbf\xd8\xf6\xd6\xf2\x9a\xe7\x86\x45\xef\x6a\x28\x4e\x3b\xc2\xc9\xa7\xd1\x66\xd7\x33\x2d\x74\x4d\x9e\xa7\xed\xa8\xf6\x78\x8a\x8e\x68\x2a\x27\xfc\x23\x51\x92\x18\x7f\x93\x4a\xaa\x6c\x73\xe0\xd1\x66\xdc\x09\x35\x48\xdf\x31\x0a\xaa\x5d\x36\x8c\x03\x83\x4e\x9b\xfb\x0b\x00\x00\xff\xff\xf1\x7d\x42\x97\xea\x00\x00\x00")

func vpcAdmissionWebhookCsrYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookCsrYaml,
		"vpc-admission-webhook-csr.yaml",
	)
}

func vpcAdmissionWebhookCsrYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookCsrYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-csr.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0xf4, 0x54, 0x57, 0xf8, 0xbd, 0x8a, 0x1a, 0x67, 0xb2, 0x29, 0x18, 0xe2, 0x44, 0x8a, 0xc2, 0x7f, 0x44, 0x26, 0x4c, 0x52, 0xe0, 0x70, 0xe0, 0xc8, 0x5a, 0x74, 0x76, 0x2, 0xcd, 0x3e, 0xa}}
	return a, nil
}

var _vpcAdmissionWebhookDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x93\xcf\x8e\xd3\x40\x0c\xc6\xef\x7d\x0a\x5f\xca\x9e\x92\xd0\x05\x71\x18\xa9\x48\x08\xc4\x05\xd8\x56\x54\xe2\xee\xce\x58\xed\x28\xf3\x4f\x63\xa7\x25\x6f\x8f\xa6\xc9\xa2\x34\xdb\xee\xce\x21\x07\xdb\xbf\xcf\xe3\x2f\x1e\x4c\xf6\x0f\x65\xb6\x31\x28\xc0\x94\xb8\x39\xad\x16\xad\x0d\x46\xc1\x37\x4a\x2e\xf6\x9e\x82\x2c\x3c\x09\x1a\x14\x54\x0b\x80\x80\x9e\x14\x9c\x92\xae\xd0\x78\xcb\x85\xac\xce\xb4\x3f\xc6\xd8\x8e\x59\x4e\xa8\x49\x41\xdb\xed\xa9\xe2\x9e\x85\xfc\x02\xc0\xe1\x9e\x1c\x17\x01\x28\x7d\xee\x29\x70\x22\x5d\x8a\x32\x25\x67\x35\xb2\x82\xd5\x02\x80\x25\xa3\xd0\xa1\x1f\x70\xe9\x13\x29\xf8\x4d\x3a\x13\x0a\x95\x34\x39\xd2\x12\xf3\x90\xf6\x28\xfa\xf8\x73\xd2\xee\xd5\x86\x00\x42\x3e\x39\x14\x1a\xe9\xc9\xa8\xe5\xb8\x2b\xa1\x37\xa4\x00\x9e\xef\x5f\x8e\x8e\x41\xd0\x06\xca\x13\xbc\x7a\xc3\xbf\xff\x6d\xf2\x61\x42\x0d\x64\x25\x8e\xbf\x52\x96\xef\xd6\xd1\xba\x21\xd1\xcd\xc8\x35\x9a\xb2\xf0\xe5\x5b\xa7\x8b\xdb\x73\xec\x07\xf5\xf7\xa8\x96\xfa\x5b\xd0\x66\x77\xb1\x70\x37\x5a\xbb\x39\x51\xce\xd6\xd0\xfa\x6c\x83\x89\x67\x9e\x97\xa3\xe3\xe8\xe2\x41\x22\x8b\xa1\x9c\xe7\xe9\xd3\xfa\xe3\x2c\xf4\xf8\xf9\xdd\x6a\x12\xb2\x1e\x0f\xa4\xe0\x61\xc9\xb5\x69\x73\x4d\x3a\xd7\x4b\xae\x97\xdc\x50\xcb\xcd\x4d\xb3\xd4\xfb\xfa\xb1\xfe\xf0\x30\xd7\xd8\x76\xce\x6d\xa3\xb3\xba\x57\xf0\xc5\x9d\xb1\x9f\x5e\xf5\x14\x5d\xe7\xe9\x57\xec\x82\xbc\x70\x77\xf8\x2f\xa3\x78\x75\xf1\xe6\xaa\x02\xc0\x17\x6e\x8b\x72\x54\xf0\xd2\xc7\x59\x6d\x26\x34\x9b\xe0\x7a\x05\x92\x3b\x1a\x93\xc7\xc8\xf2\x44\x72\x8e\xb9\xbd\x8a\x87\x68\x68\x77\xb5\xc3\xe5\xec\x49\xb0\x2e\x8f\x28\x07\x12\xe2\xda\xc6\x26\xb2\x02\x67\x43\xf7\xf7\xb5\x22\xcc\xfa\xa8\x00\xbd\xf9\xf4\xec\xf9\x30\xf6\x8d\x2d\xbc\x37\x2d\x97\xb7\x25\xd7\x0e\x0d\xb1\xa7\xfb\xdb\x3b\xaa\xfc\x0b\x00\x00\xff\xff\x66\xca\x95\xcc\x51\x04\x00\x00")

func vpcAdmissionWebhookDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookDepYaml,
		"vpc-admission-webhook-dep.yaml",
	)
}

func vpcAdmissionWebhookDepYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xbd, 0xf, 0xf7, 0xec, 0xa2, 0xbe, 0x6, 0xda, 0xf0, 0xa0, 0x99, 0xda, 0xbd, 0x2a, 0xdb, 0xed, 0xa3, 0x83, 0xb5, 0x14, 0x30, 0x49, 0x3a, 0x42, 0xf, 0x46, 0x0, 0x16, 0x2d, 0x6d, 0xff, 0xf8}}
	return a, nil
}

var _vpcAdmissionWebhookYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8c\x31\xae\xc2\x40\x0c\x05\xfb\x3d\x85\x2f\xe0\xe2\xeb\xa7\xf2\x29\x90\x90\xe8\x9d\xcd\x13\xac\x92\xcd\x5a\x6b\x13\xc4\xed\x51\x22\x0a\x1a\x44\x67\xf9\xcd\x0c\x33\x27\xb5\x72\x41\xf7\xd2\x56\xa1\xed\x2f\xcd\x65\x9d\x84\xce\xe8\x5b\xc9\x48\x15\xa1\x93\x86\x4a\x22\x5a\xb5\x42\x68\xb3\xcc\x3a\xd5\xe2\xbb\xc1\x0f\x8c\xb7\xd6\xe6\xf7\xea\xa6\x19\x42\xf3\x7d\x04\xfb\xd3\x03\x35\x11\x2d\x3a\x62\xf1\x3d\x40\xa4\x66\xdf\x0a\x6e\xc8\x3b\x64\xad\xc7\x41\xf3\x71\x0a\x0d\xc3\xff\xe1\x86\xf6\x2b\xe2\xf4\xf1\x73\x2c\xc8\xd1\xfa\xcf\xf6\x2b\x00\x00\xff\xff\xbc\xa7\x78\x3e\xe7\x00\x00\x00")

func vpcAdmissionWebhookYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcAdmissionWebhookYaml,
		"vpc-admission-webhook.yaml",
	)
}

func vpcAdmissionWebhookYaml() (*asset, error) {
	bytes, err := vpcAdmissionWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-admission-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc0, 0x14, 0x3e, 0x68, 0x1c, 0x57, 0x26, 0x7e, 0x3b, 0xab, 0xf1, 0x55, 0x21, 0x61, 0xc3, 0xe1, 0xaf, 0x46, 0xb6, 0xf7, 0xdd, 0x11, 0x29, 0x41, 0x64, 0x57, 0xc8, 0xd4, 0xea, 0x97, 0xba, 0xff}}
	return a, nil
}

var _vpcResourceControllerDepYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x93\xcf\x6e\xdc\x3c\x0c\xc4\xef\x7e\x0a\x5e\x82\x9c\xec\xdd\x4d\xbe\x2f\x41\x05\xf4\x10\x34\x40\x7b\x28\x8a\x05\x12\xf4\xce\x95\x27\xb1\x60\x59\x12\x48\x7a\x93\xed\xd3\x17\x6e\xb3\x7f\x8c\xa6\x49\x79\x32\xe6\x27\x7a\x86\x82\x58\xd7\x75\xc5\x25\x7c\x87\x68\xc8\xc9\x11\x97\xa2\x8b\xed\xaa\xea\x43\x6a\x1d\xdd\xa2\xc4\xbc\x1b\x90\xac\x1a\x60\xdc\xb2\xb1\xab\x88\x12\x0f\x70\xb4\x2d\xbe\x16\x68\x1e\xc5\xa3\xf6\x39\x99\xe4\x18\x21\x2f\x5c\x0b\x7b\x38\xea\xc7\x0d\x6a\xdd\xa9\x61\xa8\xb4\xc0\x4f\xed\x82\x12\x83\x67\x75\xb4\xaa\x88\x14\x11\xde\xb2\x4c\x84\x68\x60\xf3\xdd\x57\xde\x20\xea\x6f\x81\xa6\x48\x6f\x99\x4d\x65\x01\xe2\x68\xc3\xbe\x47\x6a\xf7\x9a\xb0\xef\x1d\xa9\xf1\x26\xa2\x22\x32\x0c\x25\xb2\xe1\xc5\xe7\x64\x9c\xa9\xe2\xcc\xf2\x9f\x4c\x5f\xb7\xfd\xd3\x98\x68\x3f\xf8\xaf\x6f\xc8\x36\x78\xdc\x78\x9f\xc7\x64\xef\x79\x4c\x02\x87\x04\x39\x44\xab\xc9\xe7\x61\xe0\xd4\x1e\xb3\xd6\xb4\x78\x2f\x29\xcb\xa3\x9e\x36\xd4\x6a\x2d\x44\xac\x13\x68\x97\x63\xfb\x31\xa4\x87\x7c\xe0\x61\xe0\x47\x38\x3a\x3f\xd3\xa6\xed\xa5\x81\x97\xe6\x4c\x9b\x33\x5d\xa0\xd7\xbf\x59\xb9\x65\x73\xd1\x5c\x9e\xcf\xff\xb1\x1e\x63\x5c\xe7\x18\xfc\xce\xd1\x4d\x7c\xe2\x9d\x1e\x78\x0c\x5b\x24\xa8\xae\x25\x6f\x70\x8c\x46\xf4\xc0\x21\x8e\x82\xfb\x7d\x34\x47\xff\x9f\xd0\xce\xac\x7c\x86\x9d\x36\x10\x75\x59\xcd\xd1\xea\xe2\xba\x59\x36\xcb\x66\x35\x63\x85\xad\x73\xb4\xe8\xc0\xd1\xba\x1f\x73\x94\xc5\x1c\x5d\xad\xae\xaf\x3f\xcc\x74\xf5\x1d\xa6\x07\xfe\xe5\xfe\x7e\x7d\x02\x42\x0a\x16\x38\xde\x22\xf2\xee\x0e\x3e\xa7\x56\x1d\x5d\x2e\x4f\x4e\x14\x48\xc8\xed\xeb\xcc\xc2\x80\x3c\xda\x01\x1e\x87\x7a\x6f\x9b\xf6\xef\xc6\x8f\x12\x6c\xf7\x29\x27\xc3\xf3\xec\x02\x8a\x84\x6d\x88\x78\x44\xeb\xc8\x64\x44\x75\xbc\x95\x6f\xb0\xa7\x2c\xfd\x4c\x4f\xb9\xc5\xdd\x6c\xe9\xa6\xda\xc0\xb8\x99\xb6\x55\x12\x0c\xda\x84\xbc\xc8\xea\x28\x86\x34\x3e\xbf\x75\x88\xc5\x77\x8e\x78\x68\xaf\xfe\xab\x7e\x06\x00\x00\xff\xff\xb7\x98\xb6\xc5\x4d\x04\x00\x00")

func vpcResourceControllerDepYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerDepYaml,
		"vpc-resource-controller-dep.yaml",
	)
}

func vpcResourceControllerDepYaml() (*asset, error) {
	bytes, err := vpcResourceControllerDepYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller-dep.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xae, 0xfc, 0xdc, 0xe6, 0x6c, 0xf, 0x10, 0x8f, 0x2f, 0x99, 0x25, 0xf5, 0xc3, 0xe3, 0x44, 0x3e, 0xb0, 0x2d, 0xfc, 0x0, 0xb4, 0xf1, 0x5c, 0x7d, 0x39, 0xa1, 0x6f, 0x33, 0xa9, 0x87, 0x80, 0xbc}}
	return a, nil
}

var _vpcResourceControllerYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x45\xd4\x3d\xad\xd8\xd0\x6d\xc0\xc0\x5e\x24\x76\x5f\xce\x6d\x4d\x73\x71\x64\x3b\x87\xe0\xe9\xd1\x35\x6d\x91\xa8\xa0\x20\xa6\xfb\xce\xf1\xef\x5f\xfa\x42\x08\x0e\x0a\x3d\xa3\x28\x71\xee\xbc\xf4\x10\x97\x50\x6d\xc7\x42\xef\x60\xc4\x79\xb9\xbf\xd5\x25\xf1\x6a\xba\x71\x7b\xca\x43\xe7\x1f\x52\x55\x43\x59\x73\x42\x37\xa2\xc1\x00\x06\x9d\xf3\x3e\xc3\x88\x9d\x9f\x4a\x0c\x82\xca\x55\x22\x86\xc8\xd9\x84\x53\x42\x71\x52\x13\x6a\xe7\x82\x87\x42\x8f\xc2\xb5\xe8\x9c\x09\x7e\xb1\x70\xde\x9f\x02\xc7\x59\xe6\x01\xf5\x93\x56\x6a\x60\xb5\x0d\x0a\x0f\x0d\x22\xe7\x0d\x6d\x47\x28\xf3\xef\x84\xd2\x1f\xb3\xb5\x0c\x60\x78\xc0\x2d\xda\xe1\x9b\x48\x1b\xbc\x82\xc5\x5d\x3b\x73\xa6\x28\x38\xef\xff\xcf\xc3\x3d\xe5\x81\xf2\xf6\x2f\x3a\x38\xe1\x1a\x37\xf3\xe2\x49\xc8\x0f\xa5\xce\xfb\x4b\xf7\xd7\x2a\xb4\xf6\x2f\x18\xed\x20\xbd\xa5\x9f\x50\x26\x8a\x78\x17\x23\xd7\x6c\x57\x0f\xb4\x77\x2d\x10\xb1\xf3\xfb\xda\x63\xd0\x37\x35\x1c\x2f\x64\x9d\x95\x7c\x29\xf8\xbd\x8e\x6f\xab\xdc\x47\x00\x00\x00\xff\xff\x16\x48\x7e\xad\xa1\x02\x00\x00")

func vpcResourceControllerYamlBytes() ([]byte, error) {
	return bindataRead(
		_vpcResourceControllerYaml,
		"vpc-resource-controller.yaml",
	)
}

func vpcResourceControllerYaml() (*asset, error) {
	bytes, err := vpcResourceControllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "vpc-resource-controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x7e, 0xd5, 0x2, 0x7a, 0xb4, 0xa2, 0x2b, 0x1, 0x1a, 0xb, 0xc2, 0x21, 0xc7, 0xa7, 0xae, 0x2, 0xa6, 0x5b, 0x95, 0xd3, 0xc5, 0xa6, 0xc7, 0x72, 0xfd, 0x35, 0xcf, 0xf0, 0x50, 0xe, 0xb, 0xdd}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"neuron-device-plugin.yaml":         neuronDevicePluginYaml,
	"vpc-admission-webhook-config.yaml": vpcAdmissionWebhookConfigYaml,
	"vpc-admission-webhook-csr.yaml":    vpcAdmissionWebhookCsrYaml,
	"vpc-admission-webhook-dep.yaml":    vpcAdmissionWebhookDepYaml,
	"vpc-admission-webhook.yaml":        vpcAdmissionWebhookYaml,
	"vpc-resource-controller-dep.yaml":  vpcResourceControllerDepYaml,
	"vpc-resource-controller.yaml":      vpcResourceControllerYaml,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"neuron-device-plugin.yaml":         &bintree{neuronDevicePluginYaml, map[string]*bintree{}},
	"vpc-admission-webhook-config.yaml": &bintree{vpcAdmissionWebhookConfigYaml, map[string]*bintree{}},
	"vpc-admission-webhook-csr.yaml":    &bintree{vpcAdmissionWebhookCsrYaml, map[string]*bintree{}},
	"vpc-admission-webhook-dep.yaml":    &bintree{vpcAdmissionWebhookDepYaml, map[string]*bintree{}},
	"vpc-admission-webhook.yaml":        &bintree{vpcAdmissionWebhookYaml, map[string]*bintree{}},
	"vpc-resource-controller-dep.yaml":  &bintree{vpcResourceControllerDepYaml, map[string]*bintree{}},
	"vpc-resource-controller.yaml":      &bintree{vpcResourceControllerYaml, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
