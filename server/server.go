package server

import (
	"myz-torrent/common"

	"github.com/gin-gonic/gin"
)

// Server app context
type Server struct {
	r    *gin.Engine
	conf *common.Config
	tmg  *common.TorrentManager
}

// Run server
func (s *Server) Run() error {
	s.initRouter()

	if err := s.initConfig(); err != nil {
		return err
	}

	if err := s.initTorrent(); err != nil {
		return err
	}

	return s.r.Run()
}

func (s *Server) initRouter() {
	r := gin.Default()

	r.GET("torrent", s.torrents)
	r.POST("torrent/magnet", s.postMagnet)

	r.GET("file", s.getAllFiles)
	r.GET("file/:file", s.downloadFile)

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

func (s *Server) initTorrent() error {
	mg, err := common.InitTorrentManager(s.conf)
	if err != nil {
		return err
	}

	s.tmg = mg
	return nil
}
