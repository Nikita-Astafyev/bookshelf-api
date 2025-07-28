package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	GetBook(ctx context.Context, id int) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id int) error
	ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error)
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	query := `
		INSERT INTO books (title, author, description, published_date)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(ctx, query,
		book.Title,
		book.Author,
		book.Description,
		book.PublishedAt,
	).Scan(&book.ID, &book.CreatedAt, &book.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to create book: %w", err)
	}

	return book, nil
}

func (r *bookRepository) GetBook(ctx context.Context, id int) (*entity.Book, error) {
	query := `
		SELECT id, title, author, description, published_date, created_at, updated_at
		FROM books
		WHERE id = $1 AND deleted_at IS NULL
	`

	var book entity.Book
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Description,
		&book.PublishedAt,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to get book: %w", err)
	}

	return &book, nil
}

func (r *bookRepository) UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	query := `
        UPDATE books
        SET 
            title = $1,
            author = $2,
            description = $3,
            published_date = $4,
            updated_at = NOW()
        WHERE id = $5 AND deleted_at IS NULL
        RETURNING updated_at
    `

	err := r.db.QueryRowContext(ctx, query,
		book.Title,
		book.Author,
		book.Description,
		book.PublishedAt,
		book.ID,
	).Scan(&book.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book not found")
		}
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return book, nil
}

func (r *bookRepository) DeleteBook(ctx context.Context, id int) error {
	query := `
        UPDATE books
        SET deleted_at = NOW()
        WHERE id = $1 AND deleted_at IS NULL
    `

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("book not found")
	}

	return nil
}

func (r *bookRepository) ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	query := `
        SELECT id, title, author, description, published_date, created_at, updated_at
        FROM books
        WHERE deleted_at IS NULL
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.Description,
			&book.PublishedAt,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return books, nil
}
