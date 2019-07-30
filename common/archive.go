package common

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipWriter provide zip compress service
type ZipWriter struct {
	w          *zip.Writer
	tmpAbsPath string
}

// NewZipWriter create a NewZipWriter
func NewZipWriter(w io.Writer) *ZipWriter {
	return &ZipWriter{
		w: zip.NewWriter(w),
	}
}

// AddPath add all files on path to zip
func (z *ZipWriter) AddPath(path string) error {
	z.tmpAbsPath = path

	if err := filepath.Walk(path, z.zipWalker); err != nil {
		return err
	}

	return nil
}

// Close finish write zip
func (z *ZipWriter) Close() error {
	return z.w.Close()
}

func (z *ZipWriter) zipWalker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	rel, err := filepath.Rel(z.tmpAbsPath, path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		rel += string(filepath.Separator)
	}

	w, err := z.w.Create(rel)
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
