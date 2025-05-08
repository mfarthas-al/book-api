package main

import (
	"fmt"
	"log"

	"book-api/database"
	"book-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := "3000"
	app := fiber.New()

	database.Connect()

	// Routes
	app.Post("/books", handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Delete("/books/:id", handlers.DeleteBook)
	app.Get("/books/search", handlers.SearchBooks)
	app.Get("/books/:id", handlers.GetBook)

	// Start server and log
	fmt.Printf("✅ Server running on port %s\n", port)
	log.Fatal(app.Listen(":" + port))
}
