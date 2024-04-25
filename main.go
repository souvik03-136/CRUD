package main

import (
	"github.com/gin-gonic/gin"
	"github.com/souvik03-136/CRUD/controllers"
	"github.com/souvik03-136/CRUD/initializers"
)

func init() {
	initializers.LoadEnvVariavles()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}
