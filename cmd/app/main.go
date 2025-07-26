package main

import (
	"github.com/Nikita-Astafyev/bookshelf-api/internal/config"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Server.ReadTimeout = cfg.Server.ReadTimeout
	e.Server.WriteTimeout = cfg.Server.WriteTimeout

	if cfg.Server.Debug {
		e.Debug = true
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "BookShelf API is running")
	})

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
