package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) torrents(c *gin.Context) {
	c.JSON(http.StatusOK, s.tmg.Torrents())
}

func (s *Server) postMagnet(c *gin.Context) {
	var magnets []string
	err := c.BindJSON(&magnets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "magnet is empty",
		})
		return
	}

	for _, m := range magnets {
		if err := s.tmg.AddMagnet(m); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": magnets,
	})
}
