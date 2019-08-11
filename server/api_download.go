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
