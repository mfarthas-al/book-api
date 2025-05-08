package main

import (
	"book-api/database"

	"github.com/gofiber/fiber/v2"

	"book-api/handlers"
)

func main() {
	app := fiber.New()
	database.Connect()

	app.Post(("/books"), handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)
	app.Get("/books/search", handlers.SearchBooks)
	app.Get("/books/:id", handlers.GetBook)

	app.Listen(":3000")
}
