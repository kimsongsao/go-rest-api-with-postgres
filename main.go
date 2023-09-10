package main

import (
	"golangrestapi/config"
	"golangrestapi/controllers"

	_ "golangrestapi/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// swagger embed files

func init() {
	config.LoadEnv()
	config.DatabaseInit()
}

// @title						Golang RESTful API with GIN, GORM & PostgreSQL
// @version					1.0
// @description				This is a sample server celler server.
// @termsOfService				http://swagger.io/terms/
// @contact.name				Kimsong SAO
// @contact.url				https://www.linkedin.com/in/kimsongsao/
// @contact.email				saokimsong@gmail.com
// @license.name				Apache 2.0
// @license.url				http://www.apache.org/licenses/LICENSE-2.0.html
// @host						localhost:3000
// @BasePath					/api/v1
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
