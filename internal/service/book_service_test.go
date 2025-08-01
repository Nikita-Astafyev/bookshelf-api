package service

import (
	"context"
	"testing"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
	"github.com/stretchr/testify/assert"
)

type FakeRepo struct {
	books map[int]*entity.Book
}

func (f *FakeRepo) CreateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	book.ID = len(f.books) + 1
	f.books[book.ID] = book
	return book, nil
}

func (f *FakeRepo) GetBook(ctx context.Context, id int) (*entity.Book, error) {
	return f.books[id], nil
}

func (f *FakeRepo) UpdateBook(ctx context.Context, book *entity.Book) (*entity.Book, error) {
	f.books[book.ID] = book
	return book, nil
}

func (f *FakeRepo) DeleteBook(ctx context.Context, id int) error {
	delete(f.books, id)
	return nil
}

func (f *FakeRepo) ListBooks(ctx context.Context, limit, offset int) ([]*entity.Book, error) {
	var result []*entity.Book
	for _, book := range f.books {
		result = append(result, book)
	}
	return result, nil
}

func TestBookService_CreateBook(t *testing.T) {
	fakeRepo := &FakeRepo{books: make(map[int]*entity.Book)}
	service := NewBookService(fakeRepo)

	book, err := service.CreateBook(context.Background(), &entity.Book{
		Title:  "Test Book",
		Author: "Test Author",
	})

	assert.NoError(t, err)
	assert.Equal(t, "Test Book", book.Title)
	assert.Equal(t, 1, book.ID)
}
