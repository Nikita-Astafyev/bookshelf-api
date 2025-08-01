package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Nikita-Astafyev/bookshelf-api/internal/entity"
	"github.com/Nikita-Astafyev/bookshelf-api/internal/service/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookController_CreateBook(t *testing.T) {
	e := echo.New()

	mockService := new(mocks.BookService)
	testBook := &entity.Book{
		Title:  "Test",
		Author: "Author",
	}
	mockService.On("CreateBook", mock.Anything, testBook).
		Return(&entity.Book{
			ID:        1,
			Title:     "Test",
			Author:    "Author",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}, nil)

	controller := NewBookController(mockService)

	req := httptest.NewRequest(
		http.MethodPost,
		"/books",
		bytes.NewBufferString(`{"title":"Test","author":"Author"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := controller.CreateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	expectedJSON := `{
    "id": 1,
    "title": "Test",
    "author": "Author",
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
	}`
	assert.JSONEq(t, expectedJSON, rec.Body.String())
	mockService.AssertExpectations(t)
}
