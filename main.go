package main

import (
	"golangrestapi/config"
	"golangrestapi/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.DatabaseInit()
}

func main() {
	r := gin.Default()
	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)
	r.PUT("/:id", controllers.UpdatePost)
	r.DELETE("/:id", controllers.DeletePost)
	r.Run()
}
