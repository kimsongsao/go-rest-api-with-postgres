package main

import (
	"fmt"
	"log"

	"kdgitial.com/rest-api-gorm/initializers"
	"github.com/kimsongsao/go-rest-api-with-postgres/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration complete")
}
