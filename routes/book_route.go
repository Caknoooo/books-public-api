package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
)

func Books(route *gin.Engine, booksController controller.BooksController) {
	books := route.Group("/api/books")
	{
		books.GET("/", booksController.GetAllBooks)
		books.GET("/:id", booksController.GetBookByID)
	}
}