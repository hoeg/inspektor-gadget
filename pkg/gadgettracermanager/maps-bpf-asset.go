// Code generated by go-bindata.
// sources:
// pkg/gadgettracermanager/maps-bpf-asset.o
// DO NOT EDIT!

package gadgettracermanager

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

var _mapsBpfAssetO = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xcf\x6b\x13\x41\x14\xfe\x36\xbf\x9a\x1c\xd4\x56\x82\x14\x04\xe9\xa9\xf4\xd2\x45\xbc\xe8\x41\xb0\x55\xa6\x25\x18\x62\x49\x42\x69\x8b\x30\x6e\xd2\x49\x5c\xdc\x8c\xcb\xce\xb4\x24\x5e\xf4\xe6\xd1\x8b\x47\x41\xff\x20\x0f\xfe\x1f\x9e\x3c\xe9\x4d\x99\xd9\xb7\xc9\x74\xcd\xc6\x83\x0f\x92\xf7\xde\xf7\xde\xf7\xde\xbc\x61\xf6\xbd\x65\xed\x83\x92\xe7\x21\x13\x0f\x3f\xb1\xf0\x16\xf2\xa0\xbc\xb0\xf7\xe8\xbf\x81\xba\xcd\xad\x03\xa8\x00\x78\x47\xf1\xc3\xa3\x36\x1a\x8d\x2a\x3e\x7f\xf7\xb0\x49\xd8\xa6\x97\xfe\x3e\xd9\x1e\x40\x15\xa8\x34\x01\x5c\x03\x50\xa2\x9c\x9b\x64\x6f\x01\xb8\x4d\xb6\xe9\xb5\x4d\xf6\x0b\x00\xf7\xc8\x36\xbd\x1e\xd2\x91\x2b\xc4\x69\xa5\xfe\x75\x8f\x7a\xf4\x28\x4e\x67\xcc\xc6\xb2\x93\x18\x4e\x8d\xf4\x49\xae\x4e\x44\x75\xaa\x54\x47\x5f\x8d\x7b\x53\x8a\xd7\x29\xfe\x3e\xd5\x37\x0c\xb8\x46\x4d\x4c\xee\x07\x07\xcf\xae\xcf\xcc\xfc\xd1\xc1\x1b\x4e\x3e\x06\xf1\x88\x4f\x82\x98\x9f\x8b\x11\xf4\x2c\x16\x78\x25\x66\x5c\x85\x6f\x04\x2e\x83\xe8\x42\xa4\xe6\x24\x98\x72\x21\x75\x12\x0a\x05\x93\x3c\x8a\x82\xb1\xc2\x85\x54\xe1\x58\x8a\xf3\xad\x50\x6a\x4c\xa4\x96\x8a\x2b\xa1\x31\x7c\x19\x24\xe0\x7c\xbf\xdb\xdd\x3f\xe5\xbd\xd6\x19\xe3\xfd\xd3\x23\xc6\x39\xda\xad\x27\xac\xd3\x63\x30\xe9\x4f\x59\xb7\xc3\xda\xfc\x98\x75\x7b\xad\x67\x1d\x44\xe1\x50\x48\x65\x1a\xc5\x0a\x97\x22\x51\xe1\x6b\xb9\xe4\x45\xa4\xb2\x6d\xcf\xfe\xe3\x77\x1e\x3f\x04\xb0\x91\x4e\xb5\x98\x0f\x00\xb7\x78\xf9\x2f\x7c\xcd\xe2\xa5\x39\xde\x24\xfd\xcb\xa9\x69\x58\xeb\xce\xe3\x34\xdc\x1d\xc7\xb7\x6c\x5f\x8b\xa9\x7b\x03\x57\xa6\x70\x47\xdb\x1d\xc4\xa3\xdd\x40\x29\xa1\xfd\x21\x7c\xa5\x13\x1d\x0c\xe0\xab\xd9\xc4\xe8\xdc\x95\xf8\x89\x88\xfc\xc7\xfd\x83\xf9\xb5\xfd\xb7\x64\x6f\xb4\x96\xc3\xf7\x0a\xf2\x2b\x39\x7f\x83\xf8\xe5\x02\x7e\xf3\x1f\xfc\x3b\x05\xfc\x7e\x41\x7e\x7e\x27\xdc\x2a\xe0\x9f\x14\xf0\xf3\xfe\xd9\x92\x9a\x46\x9e\x93\xfe\xea\xad\xee\xff\xc8\xd9\x1b\xae\xec\x10\x48\xdf\xa8\xdd\x4d\x25\xd2\x46\xb2\x7d\x74\xec\x7c\x7b\xae\x7c\x21\xfe\x5d\xf2\x6b\xe9\xae\x9a\xf3\xd7\x49\xdf\x5f\x32\xbb\x91\x6f\xc4\x1f\x63\xf5\xf9\xff\x04\x00\x00\xff\xff\x54\x64\x2e\xd7\x78\x05\x00\x00")

func mapsBpfAssetOBytes() ([]byte, error) {
	return bindataRead(
		_mapsBpfAssetO,
		"maps-bpf-asset.o",
	)
}

func mapsBpfAssetO() (*asset, error) {
	bytes, err := mapsBpfAssetOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "maps-bpf-asset.o", size: 1400, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
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
	"maps-bpf-asset.o": mapsBpfAssetO,
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
	"maps-bpf-asset.o": &bintree{mapsBpfAssetO, map[string]*bintree{}},
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
