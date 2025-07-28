package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/config"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/controller"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/repository"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/router"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/service"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	repo, err := repository.NewPostgresRepository(cfg.Postgres.GetPostgresDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	if err := runMigrations(cfg.Postgres.GetPostgresDSN()); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize layers
	bookRepo := repository.NewBookRepository(repo.GetDB())
	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	// Create and start server
	e := router.NewRouter(bookController)
	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}

func runMigrations(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migrate driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
