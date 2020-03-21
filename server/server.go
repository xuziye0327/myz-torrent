package server

import (
	"fmt"
	"myz-torrent/common"
	"myz-torrent/downloader"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server app context
type Server struct {
	r    *gin.Engine
	conf *common.Config
	dmg  *downloader.DownloadManager
}

// Run server
func (s *Server) Run() error {
	s.initRouter()

	if err := s.initConfig(); err != nil {
		return err
	}

	if err := s.initDownloader(); err != nil {
		return err
	}

	return s.r.Run(fmt.Sprintf("%v:%v", s.conf.ServerAddr, s.conf.ServerPortal))
}

func (s *Server) initRouter() {
	r := gin.Default()

	r.Static("/static/js", "./static/js")
	r.LoadHTMLGlob("./static/templates/*")

	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("download", s.listJob)
	r.POST("download", s.downloadJob)
	r.POST("download/:id", s.startJob)
	r.PUT("download/:id", s.pauseJob)
	r.DELETE("download/:id", s.deleteJob)

	r.GET("file", s.listFile)
	r.GET("file/:path", s.downloadFile)
	r.DELETE("file/:path", s.deleteFile)

	s.r = r
}

func (s *Server) initConfig() error {
	c, err := common.LoadConfig()
	if err != nil {
		return err
	}

	s.conf = c
	return nil
}

func (s *Server) initDownloader() error {
	dmg, err := downloader.Create(s.conf)
	if err != nil {
		return err
	}

	s.dmg = dmg
	return nil
}
