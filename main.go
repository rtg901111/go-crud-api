package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rtg901111/go-crud/go-crud-api/controllers"
	"github.com/rtg901111/go-crud/go-crud-api/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//r.GET("/", controllers.PostsCreate)
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	r.Run()
}
