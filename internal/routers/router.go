package routers

import (
	"fmt"

	c "github.com/anonydev/e-commerce-api/internal/controller"
	"github.com/anonydev/e-commerce-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> AA")
		c.Next()
		fmt.Println("After -> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before -> BB")
		c.Next()
		fmt.Println("After -> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before -> CC")
	c.Next()
	fmt.Println("After -> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	// /v1/2025
	v1 := r.Group("/v1/2025")
	{

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
