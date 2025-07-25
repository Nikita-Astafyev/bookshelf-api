package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "BookShelf API is running")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
