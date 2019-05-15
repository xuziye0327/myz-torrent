package util

import (
	"fmt"
	"os"
	"sort"
)

// File file struct
type File struct {
	Name     string `json:"name"`
	FullPath string `json:"full_path"`
	IsDir    bool   `json:"is_dir"`
	Child    files  `json:"child"`
}

type files []*File

// ListAllFiles list all files
func ListAllFiles(file string) (*File, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	ret := &File{
		Name:     info.Name(),
		FullPath: file,
		IsDir:    info.IsDir(),
		Child:    nil,
	}

	if ret.IsDir {
		p, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		dir, err := p.Readdir(-1)
		if err != nil {
			return nil, err
		}

		child := []*File{}
		for _, f := range dir {
			c, err := ListAllFiles(fmt.Sprintf("%v/%v", file, f.Name()))
			if err != nil {
				return nil, err
			}
			child = append(child, c)
		}

		sort.Sort(files(child))
		ret.Child = child
	}

	return ret, nil
}

func (a files) Len() int {
	return len(a)
}

func (a files) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a files) Less(i, j int) bool {
	if a[i].IsDir == a[j].IsDir {
		return a[i].Name < a[j].Name
	}
	return a[i].IsDir
}
