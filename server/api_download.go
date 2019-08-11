package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) listJob(c *gin.Context) {
	c.JSON(http.StatusOK, s.dmg.State())
}

func (s *Server) downloadJob(c *gin.Context) {
	var links []string
	err := c.BindJSON(&links)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "magnet is empty",
		})
		return
	}

	for _, m := range links {
		if err := s.dmg.New(m); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": links,
	})
}

func (s *Server) startJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Strart(id)
	c.JSON(http.StatusOK, nil)
}

func (s *Server) pauseJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Pause(id)
	c.JSON(http.StatusOK, nil)
}

func (s *Server) deleteJob(c *gin.Context) {
	id := c.Param("id")
	s.dmg.Delete(id)
	c.JSON(http.StatusOK, nil)
}
