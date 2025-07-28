package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dsn string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}

func (r *PostgresRepository) GetDB() *sql.DB {
	return r.db
}
