package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/repository"
)

type BookService interface {
	CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	GetBook(ctx context.Context, id int) (*entity.Book, error)
	UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error)
	DeleteBook(ctx context.Context, id int) error
	ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	if book.Title == "" {
		return nil, fmt.Errorf("title are required")
	}

	if book.Author == "" {
		return nil, fmt.Errorf("author are required")
	}

	if len(book.Title) > 255 {
		return nil, fmt.Errorf("book title is too long (max 255 chars)")
	}

	if len(book.Author) > 255 {
		return nil, fmt.Errorf("author title is too long (max 255 chars)")
	}

	if book.PublishedAt == nil {
		now := time.Now()
		book.PublishedAt = &now
	}

	return s.repo.CreateBook(ctx, book)
}

func (s *bookService) GetBook(ctx context.Context, id int) (*entity.Book, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid book ID")
	}

	book, err := s.repo.GetBook(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}

	return book, nil
}

func (s *bookService) UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	if book.ID == 0 {
		return nil, fmt.Errorf("invalid book ID")
	}

	if book.Title == "" {
		return nil, fmt.Errorf("book title is required")
	}

	if book.Author == "" {
		return nil, fmt.Errorf("book author is required")
	}

	existingBook, err := s.repo.GetBook(ctx, book.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to find book: %w", err)
	}

	existingBook.Title = book.Title
	existingBook.Author = book.Author
	existingBook.Description = book.Description
	existingBook.PublishedAt = book.PublishedAt

	updatedBook, err := s.repo.UpdateBook(ctx, existingBook)
	if err != nil {
		return nil, fmt.Errorf("failed to update book: %w", err)
	}

	return updatedBook, nil
}

func (s *bookService) DeleteBook(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("invalid book ID")
	}

	if _, err := s.repo.GetBook(ctx, id); err != nil {
		return fmt.Errorf("failed to find book: %w", err)
	}

	if err := s.repo.DeleteBook(ctx, id); err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

func (s *bookService) ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	books, err := s.repo.ListBooks(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}

	if books == nil {
		books = []*entity.Book{}
	}

	return books, nil
}
