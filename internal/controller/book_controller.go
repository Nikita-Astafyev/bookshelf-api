package controller

import (
	"net/http"
	"strconv"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/service"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	service service.BookService
}

func NewBookController(service service.BookService) *BookController {
	return &BookController{service: service}
}

func (c *BookController) CreateBook(ctx echo.Context) error {
	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	createBook, err := c.service.CreateBook(ctx.Request().Context(), &book)
	if err != nil {
		return ctx.JSON(getStatusCode(err), map[string]interface{}{
			"error":   "Failed to create book",
			"details": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, createBook)
}

func (c *BookController) GetBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid ID book format",
		})
	}

	book, err := c.service.GetBook(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, book)
}

func (c *BookController) UpdateBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid book ID format",
		})
	}

	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}
	book.ID = id

	updatedBook, err := c.service.UpdateBook(ctx.Request().Context(), &book)
	if err != nil {
		return ctx.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedBook)
}

func (c *BookController) DeleteBook(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid book ID format",
		})
	}

	if err := c.service.DeleteBook(ctx.Request().Context(), id); err != nil {
		return ctx.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (c *BookController) ListBooks(ctx echo.Context) error {
	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	books, err := c.service.ListBooks(ctx.Request().Context(), limit, offset)
	if err != nil {
		return ctx.JSON(getStatusCode(err), map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, books)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err.Error() {
	case "book not found":
		return http.StatusNotFound
	case "invalid book ID", "book title is required", "book author is required":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
