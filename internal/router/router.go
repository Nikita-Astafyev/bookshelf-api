package router

import (
	"github.com/Nikita-Astafyev/bookshelf-api/internal/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(bookController *controller.BookController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Gzip())

	bookGroup := e.Group("api/v1/books")
	{
		bookGroup.POST("", bookController.CreateBook)
		bookGroup.GET("/:id", bookController.GetBook)
		bookGroup.PUT("/:id", bookController.UpdateBook)
		bookGroup.DELETE("/:id", bookController.DeleteBook)
		bookGroup.GET("", bookController.ListBooks)
	}

	return e
}
