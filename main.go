package main

import (
	"golangrestapi/config"
	"golangrestapi/controllers"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"

	// gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
)

// swagger embed files

func init() {
	config.LoadEnv()
	config.DatabaseInit()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:3000
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		posts := v1.Group("/posts")
		{
			posts.GET(":id", controllers.GetPost)
			posts.GET("", controllers.GetPosts)
			posts.POST("", controllers.CreatePost)
			posts.DELETE(":id", controllers.DeletePost)
			posts.PUT(":id", controllers.UpdatePost)
			// posts.POST(":id/images", c.UploadAccountImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
