package common

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// File file struct
type File struct {
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	IsDir    bool   `json:"is_dir"`
}

// Files struct slice
type Files []*File

func (a Files) Len() int {
	return len(a)
}

func (a Files) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Files) Less(i, j int) bool {
	if a[i].IsDir == a[j].IsDir {
		return a[i].Name < a[j].Name
	}
	return a[i].IsDir
}

// ListFiles list all file under root path
func ListFiles(root string) (Files, error) {
	fs, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	ret := Files{}
	for _, f := range fs {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		ret = append(ret, &File{
			FullPath: filepath.Join(root, f.Name()),
			Name:     f.Name(),
			Size:     f.Size(),
			IsDir:    f.IsDir(),
		})
	}

	sort.Sort(ret)
	return ret, nil
}

// DeleteFile delete file under root path
func DeleteFile(root string) error {
	return os.RemoveAll(root)
}
