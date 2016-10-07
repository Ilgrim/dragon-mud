// Code generated by go-bindata.
// sources:
// assets/raw/Gamefile.toml
// DO NOT EDIT!

package assets

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

var _gamefileToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x54\x51\x6b\xe4\x46\x0c\x7e\xdf\x5f\x21\x36\x14\xbf\x14\xb3\x81\x40\x42\x21\x0f\x6d\x12\x4a\xa0\x6d\x4a\x1a\x0a\x25\x84\xa0\xb5\x65\x7b\xd2\xf1\x8c\x19\x8d\x77\x63\x8e\xfb\xef\x27\xcd\xcc\x6e\x36\x07\x47\x20\x89\x25\xcd\x27\xe9\xd3\x27\x9d\xc1\x4f\x5f\x7a\x1c\xe9\x35\x9a\x68\xe9\x2b\x6c\x91\x09\x1a\xef\x3a\xd3\xcf\x01\xa3\xf1\x0e\x3a\x63\x69\xb5\x3a\x83\xdf\x25\x0c\x1c\xc5\xbd\x0f\xff\x7f\x0e\xa9\x57\xcf\x62\x7f\x59\xad\x00\xce\xe0\xde\x45\x0a\x1d\x36\x04\x86\x21\x0e\x04\xd8\xb6\x81\x98\x61\xf1\x33\xec\xd1\xc5\x64\x64\x0a\x3b\x0a\x10\x3d\x6c\x8d\x6b\xe5\x6f\x0d\xbf\x2d\xd0\x52\x87\xb3\x8d\x60\x62\xc5\x09\x6c\x6d\x7d\x83\x76\xf0\x1c\xd7\xb0\x9d\xd5\x01\x0d\x3a\x40\xcb\xf2\x50\x51\xa2\x42\xac\x37\x75\xfa\x59\x03\x32\xec\xc9\x5a\xf0\x01\xd0\x79\x49\x14\xe0\xfe\xef\x43\x05\x09\x31\xa6\x87\xd8\x34\x62\x30\x5b\x4b\xd0\x05\x3f\x82\x9f\x23\x9b\x96\x4e\x4a\xab\x25\xda\x1c\x5b\xb9\x3e\xad\x24\xf7\x79\x33\xa0\xeb\xf3\x8b\xc9\x87\xdc\x95\x32\x09\x61\x76\x0c\xde\x49\x50\xb2\x5f\xc3\xd5\xe6\x6a\xf3\xe3\x37\x85\x89\x29\x98\x1d\xc6\x62\x2f\x10\x35\x3c\x0d\x42\xe2\x27\x97\x51\x87\x5d\x12\x9c\x0c\xc1\x51\x13\xf1\xd8\x86\xe2\x8d\xd8\x0c\xc6\x11\x68\xa9\xc6\xf5\xa7\x39\x50\x98\x2e\x43\x51\x0c\x99\xc6\xa2\x7c\x30\xb9\x36\xe1\x55\xff\xdc\x3d\xfe\x7b\xf7\x58\xc1\x28\xe4\x60\x4f\xac\x5e\x14\x3a\x47\x1a\x69\x4b\x41\x52\x77\xc7\x3e\x6b\x78\x48\xfc\x46\xe9\x49\x7e\x19\xfe\x39\x8d\x2d\xe7\x13\x7f\x42\x94\x79\x94\x22\x53\x29\x3e\x7b\xa3\x14\x82\xa1\x4d\x0d\xd5\x29\xee\xe9\xe1\xf6\xe1\x17\xf8\x95\x59\x85\x22\xe1\x3a\x9c\xb8\x64\x5f\x69\xff\xf5\x83\xcd\x73\x95\xa3\xf5\xbd\x62\x47\x34\xc2\x95\x08\x41\x33\xc8\x3f\x13\x35\xa6\x33\xcd\x21\x97\x44\xf5\x52\x64\xe7\x43\x66\x3d\xf8\x37\x29\x06\x78\x6e\x06\x2d\x6e\xc4\x77\x33\xce\x63\x81\xb3\xb4\x23\x9b\x58\x12\x3d\x4c\x22\xb7\x88\xa1\xa7\xc8\xa2\x6f\x71\x27\x7d\xe7\x10\x91\x43\x4b\xdb\xb9\x2f\x52\xb8\xa5\x4e\x19\x57\x88\xf2\x22\x99\xff\xd0\xef\x65\x4a\xfb\xc4\x86\xa3\xb2\x57\x79\xae\x52\x86\x4a\xd7\xaa\xd2\x09\x53\x36\x6a\x24\x27\x75\x0f\xb8\x13\x85\x16\x28\x79\x94\x67\xc3\xb1\x95\xaa\xca\x63\xf9\xa0\x10\x2a\xd8\x0f\x82\x92\x3a\x3b\x46\xa7\xaf\x8c\x5e\x30\x79\xf0\xb3\x6d\x55\xf6\x64\x74\x64\x79\x32\x10\xc8\xca\xf6\x4a\xaa\x09\xe3\x70\xc2\x57\x5a\xf8\xb4\x41\x5b\xf6\x76\x8e\x39\xa0\x86\x3f\x65\x33\xcd\x64\x3f\xf5\x09\x18\xf2\xa0\x79\x9e\x74\x3e\xd4\xea\xcc\x9e\x95\xae\xba\x84\xbc\x24\xde\x20\x33\x21\xc4\x79\x5e\xe7\xef\x5c\xb0\x58\x72\x67\x6b\x1d\xea\x4d\x39\x2b\xb9\xa7\x83\x74\xe4\x08\x19\x27\x23\x1c\xf3\x41\x92\x52\xff\x22\x7f\xf1\x56\xc3\x7f\x72\x52\x94\xb1\x59\x8e\x16\xba\x05\xc8\xed\x4c\xf0\x6e\x24\x17\x05\xcc\xe9\x36\x1e\xaf\x0e\xe6\x0b\xa4\xe1\x59\x24\x8b\xb2\x27\x32\x38\x79\xa4\xd0\xf4\x4e\x8d\x36\xbd\x1f\xc8\xe9\x22\x3a\xd1\x95\x80\x9d\x5c\x86\xbc\x95\xa5\x38\x3e\x30\x77\xb8\x5d\x52\x4b\x48\x99\x75\x50\x13\x32\xcb\xbd\x6c\x75\x2e\xb9\xe6\xd5\x73\x8b\xb2\xb1\x72\x66\xeb\x56\xb5\xe4\x27\x4d\x9c\x38\xc2\x59\x70\x5c\x34\x4d\x6e\xf3\x1a\x62\x98\x95\x5d\xdd\xe4\xef\x2e\xd0\xf1\xb6\xac\x2f\x2f\x2e\x2f\xf4\xfb\x98\x56\x6c\x4e\x33\xa5\xa0\x43\xfa\x53\xe3\x07\xab\xaf\x22\x7f\x71\x9d\x6f\x56\xdf\x02\x00\x00\xff\xff\xd4\x17\x2e\x84\x07\x06\x00\x00")

func gamefileTomlBytes() ([]byte, error) {
	return bindataRead(
		_gamefileToml,
		"Gamefile.toml",
	)
}

func gamefileToml() (*asset, error) {
	bytes, err := gamefileTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Gamefile.toml", size: 1543, mode: os.FileMode(420), modTime: time.Unix(1475785686, 0)}
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
	"Gamefile.toml": gamefileToml,
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
	"Gamefile.toml": &bintree{gamefileToml, map[string]*bintree{}},
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

