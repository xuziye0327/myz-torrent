package common

import (
	"fmt"
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

func ListFiles(root string) (Files, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("Path: {%v} is not a dir ", root)
	}

	p, err := os.Open(root)
	if err != nil {
		return nil, err
	}

	fs, err := p.Readdir(-1)
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
