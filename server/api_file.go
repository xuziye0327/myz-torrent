package server

import (
	"encoding/base64"
	"fmt"
	"myz-torrent/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (s *Server) getAllFiles(c *gin.Context) {
	root := s.conf.DownloadDir

	file, err := util.ListAllFiles(root)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, file)
}

func (s *Server) downloadFile(c *gin.Context) {
	base64Path := c.Param("file")

	decodePath, err := base64.StdEncoding.DecodeString(base64Path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	f, err := os.Open(string(decodePath))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	info, err := f.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	data := make([]byte, info.Size())
	l, err := f.Read(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	content := make([]byte, 512)
	if _, err := f.Read(content); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%v", f.Name()))
	c.Header("Content-Type", http.DetectContentType(content))
	c.Header("Accept-Length", fmt.Sprintf("%v", l))
	if _, err := c.Writer.Write(data); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
}
