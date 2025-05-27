package routes

import (
	"github.com/Caknoooo/go-gin-clean-starter/controller"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Books(route *gin.Engine, injector *do.Injector) {
	booksController := do.MustInvoke[controller.BooksController](injector)

	books := route.Group("/api/books")
	{
		books.GET("/", booksController.GetAllBooks)
		books.GET("/:id", booksController.GetBookByID)
	}
}
