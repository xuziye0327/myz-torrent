package server

import (
	"encoding/base64"
	"fmt"
	"myz-torrent/common"
	"myz-torrent/util"
	"net/http"
	"os"
	"path/filepath"

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

	path := filepath.Join(s.conf.DownloadDir, string(decodePath))
	info, err := os.Stat(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fileName := info.Name()
	if info.IsDir() {
		fileName += ".zip"
	}

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
