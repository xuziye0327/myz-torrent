package server

import (
	"myz-torrent/common"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *Server) index(c *gin.Context) {
	root := s.conf.DownloadDir
	fs, err := common.ListFiles(root)
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

	states := s.dmg.State()

	c.JSON(http.StatusOK, gin.H{
		"files":  fs,
		"states": states,
	})
}
