package server

import (
	"encoding/base64"
	"fmt"
	"myz-torrent/common"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAllFiles(c *gin.Context) {
	root := s.conf.DownloadDir
	path := c.Query("path")

	if len(path) != 0 {
		decodePath, err := pathDecoder(path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		path = string(decodePath)
	}
	fs, err := common.ListFiles(filepath.Join(root, path))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	for _, f := range fs {
		rel, err := filepath.Rel(root, f.FullPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		f.FullPath = rel
	}

	c.JSON(http.StatusOK, fs)
}

func (s *Server) downloadFile(c *gin.Context) {
	encodePath := c.Param("file")

	decodePath, err := pathDecoder(encodePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	path := filepath.Join(s.conf.DownloadDir, string(decodePath))
	info, err := os.Stat(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fileName := info.Name() + ".zip"

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
	c.Header("Content-Type", "application/zip")

	zip := common.NewZipWriter(c.Writer)
	if err := zip.AddPath(path); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	zip.Close()
}

// path => base64 => urlEncode => origin
func pathDecoder(path string) (string, error) {
	b64Decode, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return "", err
	}

	uDecode, err := url.Parse(string(b64Decode))
	if err != nil {
		return "", err
	}

	return uDecode.Path, nil
}
