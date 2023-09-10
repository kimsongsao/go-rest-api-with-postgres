package main

import (
	"golangrestapi/config"
	"golangrestapi/models"
)

func init() {
	config.LoadEnv()
	config.DatabaseInit()
}

func main() {
	config.DB.AutoMigrate(&models.Post{})
	config.DB.AutoMigrate(&models.User{})
}
