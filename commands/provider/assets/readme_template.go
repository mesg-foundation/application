// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// commands/provider/assets/readme_template.md
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

var _commandsProviderAssetsReadme_templateMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x5b\x8b\xeb\x36\x10\x7e\xd7\xaf\x18\xf0\x81\x26\xe1\xd8\x79\x0f\xed\xd3\xd9\xa5\x1c\xb6\xed\x42\x37\x2f\x25\x04\xac\xd8\x4a\x2c\xe2\x48\xa9\x24\xef\x12\xb4\xfa\xef\x65\x24\x5f\x14\xc7\x7b\x81\xb2\x2f\xc9\x68\x6e\xfa\xbe\xf9\xac\x49\xc0\xda\xec\x2f\x7a\x62\xce\x11\x62\x6d\x76\xc7\x74\xa1\xf8\xd9\x70\x29\xd0\x93\xc0\x0f\x29\x0c\x13\x46\x13\x92\xc2\xe6\xa7\xd0\x86\xd6\x35\xc5\xf0\x76\x96\xc4\xc7\x39\xc6\xef\xd8\x9e\x0b\x8e\x47\xbd\x9d\x25\xd1\x69\x4e\x00\xac\xe5\x7b\xc8\xee\x9f\xb1\x9b\x73\x29\x6c\x82\xb9\x9d\x25\xc1\x98\x5b\xab\xa8\x38\x30\xf8\x76\x64\x97\xef\xf0\x8d\xa1\x17\x56\xbf\x0d\x35\x04\x00\x20\x85\x8d\xb5\x52\xb5\x71\x8f\xdd\x57\x38\xb7\x9d\x25\x93\x11\x78\x05\x2a\x8a\x4a\xaa\x7b\x51\xc8\x92\x39\x37\xb7\x96\x89\xd2\xb9\xf6\xaf\xc7\xb6\xa6\xfa\x18\xa0\x79\x6b\x3b\x4b\xfc\xff\x18\x98\xa1\xfa\xe8\x71\xb5\xf9\xd7\xb0\x30\x3a\x89\xea\x3a\xf0\x21\x28\x92\x40\x3c\x5f\x42\x92\x04\xfe\xbc\x7f\xfa\x1d\x7e\x48\xc5\x08\x59\x57\x5c\x83\x66\xea\x99\x17\x0c\x14\xfb\xb7\xe1\x8a\x69\xd8\xf4\x19\xdb\x59\x65\xcc\x59\xaf\x96\xcb\x03\x37\x55\xb3\xcb\x0a\x79\x5a\x9e\x98\x3e\xa4\x7b\xd9\x88\xd2\xf7\x5c\x16\x52\xb1\x39\x18\x09\x3b\x06\x3c\x5c\xc6\x4a\xd8\x73\xa5\x4d\x46\xc8\x3f\xb2\x81\x82\x8a\x2e\x32\xdc\x0e\xbb\x0b\xa8\x46\x08\x2e\x0e\x60\x2a\x06\x7b\x59\xd7\xf2\x05\x4f\x85\x3c\x9d\xa8\x28\x41\x2a\xd8\x04\xaf\x4f\xe0\x11\x11\x38\x34\xbc\x8c\xe0\x95\xb2\xd0\x19\x02\xcb\x0c\x2b\xaa\xa5\x8f\x2e\xb5\xa1\xca\xa4\x15\x53\x6c\x19\xd7\x66\x95\x39\xd5\xf3\x8c\x90\x3c\xcf\x77\x54\x57\x04\x7f\xe0\xd7\x59\xd1\xa8\x1a\xd2\xbd\x7e\xfa\x03\xba\xb6\xbe\x23\x72\x6e\xeb\xe7\x58\xe3\x87\xf8\x14\x86\x86\x9f\x3b\x8a\xfe\x37\x3b\x4b\xcd\x8d\x54\x17\xe7\xd6\x12\x4a\x76\xae\xe5\x05\x4c\x34\xde\xef\x48\x76\x9a\xe8\xaa\x47\xe2\x47\x8b\xf3\xec\x45\x69\x3b\x59\x7b\x75\x85\x87\x61\x2d\xab\x35\x73\xee\x4e\xbe\x88\x5a\xd2\xd2\x37\xd7\xb2\x51\x05\x03\xfc\x1a\x40\xee\x47\x08\x70\xa8\xa6\x62\x28\x46\x48\x0e\x81\x5f\x34\x62\x2a\x99\xfa\x5f\x18\x3b\x4c\xfe\xbb\x4b\x20\x7a\xb7\xdd\x90\xba\x17\x98\x40\xb0\xc8\xc7\x6f\x35\xc1\xd5\x32\xf9\x4e\x09\xf1\x49\x70\x64\x97\x15\xe4\xd6\x06\x6f\x8e\x77\xb5\xc9\xa3\x2d\xe4\x31\x74\x21\x6a\xa8\x73\xaf\xb0\x58\x60\xc7\xc5\x02\xd0\x7c\x60\x97\xd6\x5a\x5f\xce\x9d\x33\x6a\x82\x1e\xf2\x0a\x69\x9a\xc2\xc4\xef\x40\xa6\xa4\x86\x3e\x78\x42\x68\x21\x9f\x9b\x5b\x03\x27\x0c\xb7\x94\xda\x1a\xe7\xfc\xb5\x48\xa7\xf7\xe4\x91\x23\x43\x64\xc1\x15\xe8\x78\xe7\xa3\x87\x47\x6b\xac\xce\x65\x7b\xc8\x17\x0b\xe8\x97\x41\xc8\xbb\x1a\x88\x47\x1c\xef\x8a\xd1\xe6\xb8\x5a\x66\x09\x78\x63\x2c\xd8\xcd\x0e\xeb\xe5\x1a\xef\x2f\x42\x30\x65\x52\x2c\x9f\x3a\xa9\x95\x8f\xfc\x14\xe7\xc6\x7f\x34\x09\xae\x32\xb4\x09\xf9\x22\xe1\x38\xb6\x0f\xca\x79\xd3\x4b\x77\x05\x62\xd0\xce\x27\xb4\x0c\xbb\xba\x48\xbd\xc1\x95\xc7\x9e\x1b\xfd\x82\xf7\x13\x02\x86\xc4\xf7\x15\x8c\xc7\xf6\xd8\x98\x61\x6e\xed\x81\x0c\x54\xa5\xf7\x04\xae\xc1\x1e\xc8\xc6\xa5\x9d\x9e\x21\xa7\xe5\xdb\x17\xa3\x54\x21\x7b\x50\x36\x0a\x06\x7d\xdb\xd2\x91\xc2\x5f\x24\xe1\x1b\xbc\x3a\x0c\xa3\x07\xf8\x16\xab\x5e\xc6\x98\xcc\x95\xeb\x46\xc8\xd6\xfd\x09\x25\xa7\xe6\x31\x48\xf9\xfe\xa3\xfc\x2f\x00\x00\xff\xff\x4b\xfe\xe3\xf0\x69\x09\x00\x00")

func commandsProviderAssetsReadme_templateMdBytes() ([]byte, error) {
	return bindataRead(
		_commandsProviderAssetsReadme_templateMd,
		"commands/provider/assets/readme_template.md",
	)
}

func commandsProviderAssetsReadme_templateMd() (*asset, error) {
	bytes, err := commandsProviderAssetsReadme_templateMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "commands/provider/assets/readme_template.md", size: 2409, mode: os.FileMode(420), modTime: time.Unix(1544093430, 0)}
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
	"commands/provider/assets/readme_template.md": commandsProviderAssetsReadme_templateMd,
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
	"commands": &bintree{nil, map[string]*bintree{
		"provider": &bintree{nil, map[string]*bintree{
			"assets": &bintree{nil, map[string]*bintree{
				"readme_template.md": &bintree{commandsProviderAssetsReadme_templateMd, map[string]*bintree{}},
			}},
		}},
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
