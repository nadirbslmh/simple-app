package controllers

import (
	"net/http"
	"simple-app/models"
	"simple-app/services"

	"github.com/labstack/echo/v4"
)

type BookController struct{}

var bookService services.BookService = services.BookService{}

func (bc *BookController) GetAll(c echo.Context) error {
	books := bookService.GetAll()

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "all books",
		"data":    books,
	})
}

func (bc *BookController) GetByID(c echo.Context) error {
	bookID := c.Param("id")

	book := bookService.GetByID(bookID)

	if book.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "book found",
		"data":    book,
	})
}

func (bc *BookController) Create(c echo.Context) error {
	var book models.Book

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "validation failed",
		})
	}

	createdBook := bookService.Create(book)

	return c.JSON(http.StatusCreated, map[string]any{
		"status":  "success",
		"message": "book created",
		"data":    createdBook,
	})
}

func (bc *BookController) Update(c echo.Context) error {
	bookID := c.Param("id")

	var book models.Book

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "validation failed",
		})
	}

	updatedBook := bookService.Update(bookID, book)

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "book updated",
		"data":    updatedBook,
	})
}

func (bc *BookController) Delete(c echo.Context) error {
	bookID := c.Param("id")

	isDeleted := bookService.Delete(bookID)

	if !isDeleted {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "book not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "book deleted",
	})
}
