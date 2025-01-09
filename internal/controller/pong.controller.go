package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

// Return pointer to PongController
func NewPongController() *PongController {
	return &PongController{}
}

// uc user controller
// us user services
func (uc *PongController) Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "Guest")
	nameStart := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "ping...`pong" + nameStart,
		"default": name,
		"uid":     uid,
		"user":    []string{"@nonydev", "cr7", "messi"},
	})
}
