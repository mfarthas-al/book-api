package handlers

import (
	"book-api/models"
	"book-api/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	offset := (page - 1) * limit

	books, err := services.GetBooksPaginated(limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := services.GetBookByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	return c.JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := services.CreateBook(book); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create book"})
	}
	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	book, err := services.GetBookByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	return c.JSON(services.UpdateBook(book))
}

func DeleteBook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := services.DeleteBook(uint(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)

}

func SearchBooks(c *fiber.Ctx) error {
	query := c.Query("q")
	books, err := services.SearchBooks(query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	if len(books) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No books found"})
	}
	return c.JSON(books)
}
