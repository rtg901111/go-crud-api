package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rtg901111/go-crud/go-crud-api/initializers"
	"github.com/rtg901111/go-crud/go-crud-api/models"
)

// Create API - create a row in db from the HTTP Post request
func PostsCreate(c *gin.Context) {
	// Get data from request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Read API - read all data from the table in db
func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}
