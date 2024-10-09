package main

import (
	"blog-crud/initializer"
	"blog-crud/models"
)

func init() {
	initializer.ConnectDB()

}
func main() {
	initializer.DB.AutoMigrate(&models.Post{})
}
