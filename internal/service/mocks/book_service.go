package mocks

import (
	"context"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
	"github.com/stretchr/testify/mock"
)

type BookService struct {
	mock.Mock
}

func (m *BookService) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	args := m.Called(ctx, book)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *BookService) GetBook(ctx context.Context, id int) (*entity.Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *BookService) UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	args := m.Called(ctx, book)
	return args.Get(0).(*entity.Book), args.Error(1)
}

func (m *BookService) DeleteBook(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *BookService) ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*entity.Book), args.Error(1)
}
