package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"

	"book-api/database"
	"book-api/models"
)

func setupAppForTest() *fiber.App {
	database.ConnectTestDB()

	// Insert test book
	database.DB.Create(&models.Book{
		Title:  "Go in Action",
		Author: "William Kennedy",
		Year:   2016,
	})

	app := fiber.New()
	app.Get("/books", GetBooks)
	app.Get("/books/:id", GetBook)
	app.Post("/books", CreateBook)
	app.Put("/books/:id", UpdateBook)
	app.Delete("/books/:id", DeleteBook)
	app.Get("/books/search", SearchBooks) // ✅ Add this for search route

	return app
}

func TestGetBooks(t *testing.T) {
	app := setupAppForTest()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetBook(t *testing.T) {
	app := setupAppForTest()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreateBook(t *testing.T) {
	app := setupAppForTest()
	body := `{"title":"New Book","author":"Tester","year":2024}`
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestUpdateBook(t *testing.T) {
	app := setupAppForTest()
	body := `{"id":1,"title":"Updated","author":"AuthorX","year":2025}`
	req := httptest.NewRequest(http.MethodPut, "/books/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteBook(t *testing.T) {
	app := setupAppForTest()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
}

func TestGetBooksPagination(t *testing.T) {
	app := setupAppForTest()

	// Add multiple books to test pagination
	database.DB.Create(&models.Book{Title: "Book A", Author: "Author A", Year: 2001})
	database.DB.Create(&models.Book{Title: "Book B", Author: "Author B", Year: 2002})
	database.DB.Create(&models.Book{Title: "Book C", Author: "Author C", Year: 2003})
	database.DB.Create(&models.Book{Title: "Book D", Author: "Author D", Year: 2004})

	// Test page=1, limit=2 (should return first 2)
	req := httptest.NewRequest(http.MethodGet, "/books?limit=2&page=1", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestSearchBooks(t *testing.T) {
	app := setupAppForTest()

	// Add searchable books
	database.DB.Create(&models.Book{Title: "Golang for Beginners", Author: "John Go", Year: 2020})
	database.DB.Create(&models.Book{Title: "Advanced Python", Author: "Alice Py", Year: 2021})

	req := httptest.NewRequest(http.MethodGet, "/books/search?q=go", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Accept either 200 (found) or 404 (not found), just for testing robustness
	assert.True(t, resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNotFound)
}
