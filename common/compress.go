package common

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

// Compressor provide compress service
type Compressor struct {
	FilePath string

	zw *zip.Writer
}

// NewCompressor create a Compressor
func NewCompressor(filePath string) *Compressor {
	return &Compressor{
		FilePath: filePath,
	}
}

// Zip compress dir with zip
func (c *Compressor) Zip() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	c.zw = zip.NewWriter(buf)
	defer c.zw.Close()

	if err := filepath.Walk(c.FilePath, c.zipWalker); err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *Compressor) zipWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(c.FilePath, path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		rel += string(filepath.Separator)
	}

	w, err := c.zw.Create(rel)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(w, f); err != nil {
			return err
		}
	}

	return nil
}
