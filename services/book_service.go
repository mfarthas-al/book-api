package services

import (
	"book-api/database"
	"book-api/models"
	"strings"
)

func CreateBook(book models.Book) error {
	return database.DB.Create(&book).Error
}

func GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

func GetBookByID(id uint) (models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	return book, result.Error
}

func UpdateBook(book models.Book) error {
	return database.DB.Save(&book).Error
}

func DeleteBook(id uint) error {
	return database.DB.Delete(&models.Book{}, id).Error
}

func GetBooksPaginated(limit int, offset int) ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Limit(limit).Offset(offset).Find(&books)
	return books, result.Error
}

func SearchBooks(query string) ([]models.Book, error) {
	var books []models.Book
	result := database.DB.
		Where("LOWER(title) LIKE ? OR LOWER(author) LIKE ?", "%"+strings.ToLower(query)+"%", "%"+strings.ToLower(query)+"%").
		Find(&books)
	return books, result.Error
}
