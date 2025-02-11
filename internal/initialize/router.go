package initialize

import (
	c "github.com/anonydev/e-commerce-api/internal/controller"
	"github.com/anonydev/e-commerce-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// /v1/2025
	v1 := r.Group("/v1/2025")
	{
		// Invoke method GetUserByID from user controller
		v1.GET("/user/1", c.NewUserController().GetUserByID)
		v1.GET("/ping", c.NewPongController().Pong)
	}

	// /v2/2025
	v2 := r.Group("/v2/2025")
	{
		v2.GET("/ping", temp_api)
		v2.PUT("/ping", temp_api)
		v2.PATCH("/ping", temp_api)
		v2.DELETE("/ping", temp_api)
		v2.POST("/ping", temp_api)
		v2.HEAD("/ping", temp_api)
		v2.OPTIONS("/ping", temp_api)
	}
	return r
}

func temp_api(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "temp api",
	})
}
