package controllers

import (
	"blog-crud/initializer"
	"blog-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {

	var body struct {
		Title   string `json:"title"`
		Author  string `json:"author"`
		Content string `json:"content"`
	}

	ctx.Bind(&body)

	post := models.Post{Title: body.Title, Author: body.Author, Content: body.Content}
	result := initializer.DB.Create(&post)

	if result.Error != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message ": "Error creating post",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "create successfully",
		"result":  post,
	})
}

func GetAllPosts(ctx *gin.Context) {
	var posts []models.Post
	initializer.DB.Find(&posts)
	ctx.IndentedJSON(http.StatusOK, posts)
}

func GetPostById(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	initializer.DB.First(&post, id)
	ctx.IndentedJSON(http.StatusOK, post)
}

func DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	var post models.Post
	err := initializer.DB.Delete(&post, id).Error
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error deleting post",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Delete post successfully",
	})
}

func UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Post

	ctx.Bind(&post)

	err := initializer.DB.Model(&post).Where("id = ?", id).Updates(models.Post{Title: post.Title, Author: post.Author, Content: post.Content}).Error

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error updating post",
		})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Update post successfully",
	})
}
