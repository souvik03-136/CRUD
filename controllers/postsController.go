package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/CRUD/initializers"
	"github.com/souvik03-136/CRUD/models"
)

func PostsCreate(c *gin.Context) {

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

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id of URL
	id := c.Param("id")
	//Get the post
	var post models.Post
	initializers.DB.First(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//Get the id of the url
	id := c.Param(":id")

	//Get the data of the req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, Body: body.Body,
	})

	//Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//Respond
	c.Status(200)
}
