package main

import (
	"github.com/rtg901111/go-crud/go-crud-api/initializers"
	"github.com/rtg901111/go-crud/go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
