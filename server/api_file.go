package server

import (
	"encoding/base64"
	"fmt"
	"log"
	"myz-torrent/common"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type path string

func (s *Server) getAllFiles(c *gin.Context) {
	root := s.conf.DownloadDir
	p := path(c.Query("path"))
	target, err := p.validate(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fs, err := common.ListFiles(target)
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
	root := s.conf.DownloadDir
	p := path(c.Param("path"))
	targat, err := p.validate(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	info, err := os.Stat(targat)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fileName := info.Name() + ".zip"

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v", fileName))
	c.Header("Content-Type", "application/zip")

	zip := common.NewZipWriter(c.Writer)
	defer zip.Close()

	if err := zip.AddPath(targat); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}

func (s *Server) deleteFile(c *gin.Context) {
	root := s.conf.DownloadDir
	p := path(c.Param("path"))
	targat, err := p.validate(root)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := common.DeleteFile(targat); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
	}
}

func (p path) validate(root string) (string, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("get root abd path %v error %v", root, err)
	}

	target, err := p.decode()
	if err != nil {
		return "", fmt.Errorf("path decode error %v", err)
	}

	res, err := filepath.Abs(filepath.Join(root, target))
	if err != nil {
		return "", fmt.Errorf("get abs path error %v", err)
	}

	if !strings.Contains(res, root) {
		return "", fmt.Errorf("invaild path %v", res)
	}

	return res, nil
}

// path => base64 => urlEncode => origin
func (p path) decode() (string, error) {
	b64Decode, err := base64.StdEncoding.DecodeString(string(p))
	if err != nil {
		return "", err
	}

	uDecode, err := url.Parse(string(b64Decode))
	if err != nil {
		return "", err
	}

	return uDecode.Path, nil
}
