package main

import (
	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Book
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Category
	r.GET("/category", controllers.FindCategories)
	r.GET("/category/:id", controllers.FindCategory)
	r.POST("/category", controllers.CreateCategory)
	r.PATCH("/category/:id", controllers.UpdateCategory)
	r.DELETE("/category/:id", controllers.DeleteCategory)

	// Run the server
	r.Run()
}
