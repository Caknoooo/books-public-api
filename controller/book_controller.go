package controller

import (
	"net/http"

	"github.com/Caknoooo/go-gin-clean-starter/service"
	"github.com/Caknoooo/go-gin-clean-starter/utils"
	"github.com/gin-gonic/gin"
)

type (
	BooksController interface {
		GetAllBooks(ctx *gin.Context)
		GetBookByID(ctx *gin.Context)
	}

	booksController struct {
		booksService service.BooksService
	}
)

func NewBooksController(booksService service.BooksService) BooksController {
	return &booksController{
		booksService: booksService,
	}
}

const (
	MESSAGE_SUCCESS_GET_ALL_BOOKS = "Successfully retrieved all books"

	MESSAGE_FAILED_GET_ALL_BOOKS = "Failed to retrieve all books"
)

func (c *booksController) GetAllBooks(ctx *gin.Context) {
	books, err := c.booksService.GetAllBooks()
	if err != nil {
		res := utils.BuildResponseFailed(MESSAGE_FAILED_GET_ALL_BOOKS, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess(MESSAGE_SUCCESS_GET_ALL_BOOKS, books)
	ctx.JSON(http.StatusOK, res)
}

func (c *booksController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		res := utils.BuildResponseFailed("Book ID is required", "Book ID cannot be empty", nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	book, err := c.booksService.GetBookByID(id)
	if err != nil {
		res := utils.BuildResponseFailed("Failed to retrieve book", err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.BuildResponseSuccess("Successfully retrieved book", book)
	ctx.JSON(http.StatusOK, res)
}
