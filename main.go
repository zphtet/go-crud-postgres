package main

import (
	"blog-crud/controllers"
	"blog-crud/initializer"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.ConnectDB()
}
func main() {
	fmt.Println("MAIN GO")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ROUTES
	r.POST("/create-post", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)

	r.GET("/post/:id", controllers.GetPostById)

	r.DELETE("/post/:id", controllers.DeletePost)

	r.PATCH("/post/:id", controllers.UpdatePost)

	// RUN Server on
	r.Run("localhost:8000")

}
