// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/index.html
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xdf\x8a\x9c\x30\x14\xc6\xef\xfb\x14\xa7\x81\x42\x7b\xa1\xd1\x29\xb2\xb4\xc4\x5c\x94\x42\x7b\x51\x7a\xb3\x7d\x81\xa8\x67\xc7\xb0\x49\x94\x24\x16\xa6\xe2\xbb\x97\xac\x13\x66\x47\x9c\x19\x07\xd6\x2b\xcd\xf9\x9d\x3f\xdf\x17\x39\xac\xf5\x5a\xf1\x77\x00\x00\xac\x45\xd1\xcc\xaf\x2f\x9f\xef\x93\x04\xbe\x75\x9d\x77\xde\x8a\x1e\x92\xe4\x55\x48\x49\xf3\x0c\xad\xc5\xa7\x92\xd4\xce\xd1\x2a\x52\xa9\x96\x26\xad\x9d\x23\x60\x51\x95\xc4\xf9\x83\x42\xd7\x22\x7a\x72\x6c\x41\x4f\x3d\x58\xd5\x35\x87\x57\x35\x9f\x3a\xab\x41\xd4\x5e\x76\xa6\x24\xd4\x29\xc4\xfe\x8f\xd4\x68\xc9\x89\x99\xa7\xcc\x39\xfc\x1c\x10\x1e\x03\xe1\x03\x11\xca\xe6\x0b\x4a\x89\x0a\x15\xff\x25\xf7\xad\xff\x7a\x16\x79\x89\x3a\x54\x58\x7b\x30\x42\x63\x49\x54\x80\x08\x38\xf9\x0f\x4b\x92\x2f\xda\xc5\x67\x1c\xad\x30\x7b\x84\x74\x9a\x56\xe3\xac\xeb\xc3\xe4\x7c\x1c\xd3\xdf\x42\xe3\x34\x31\x7a\x3c\xb9\x50\x0e\x4d\xb3\x52\x8a\xd1\x79\xb4\x85\x1a\x3a\xcb\x39\x3f\xac\xec\xaa\xe6\xef\x83\x15\xa1\x31\x7c\xd4\xd2\x0c\x1e\x3f\xdd\xd2\xdf\x1c\x13\x6e\x59\x10\x25\xe6\xd7\xa5\x45\xac\xd8\x86\xe5\xd9\x26\x0e\xe6\x91\xb1\xe1\xbb\x6d\x09\xfc\xf3\x15\xee\x0d\x6c\x7e\xf4\xc2\x7a\x69\xf6\x50\xd9\xf0\xff\x18\x74\xee\x96\xd1\x27\x72\xa3\xd5\xf0\x57\xa8\x21\x60\xbb\x07\xc2\x8b\xec\xc3\x36\xa7\x62\xd2\x97\x8c\xf0\x87\xe2\xbe\xa4\x5d\x51\x10\x9e\x67\x5b\x5b\xc5\x4b\x89\xe9\x19\xe1\xcf\x88\xfd\xbd\xc6\xaf\xfa\x7c\xf9\x50\x9a\x7e\xf0\xe0\x0f\x3d\x96\xc4\x0d\x95\x96\x9e\xc4\x01\x7e\x74\xcb\x75\xb1\xbc\x53\x46\xc3\xa2\x89\xeb\x68\xde\x41\x8c\xce\x5b\xf0\x7f\x00\x00\x00\xff\xff\x8f\x38\xd9\x63\x0d\x05\x00\x00")

func assetsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexHtml,
		"assets/index.html",
	)
}

func assetsIndexHtml() (*asset, error) {
	bytes, err := assetsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.html", size: 1293, mode: os.FileMode(420), modTime: time.Unix(1602928584, 0)}
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
	"assets/index.html": assetsIndexHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
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
