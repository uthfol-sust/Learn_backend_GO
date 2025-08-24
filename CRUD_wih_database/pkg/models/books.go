package models

import (
	"github.com/jinzhu/gorm"
	"github.com/uthfol-sust/Learn_backend_GO/CRUD_wih_database/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// AutoMigrate ensures the table exists
func Migrate() {
	db := config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook inserts a new book record
func (b *Book) CreateBook() *Book {
	db := config.GetDB()
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks returns all books
func GetAllBooks() []Book {
	var books []Book
	db := config.GetDB()
	db.Find(&books)
	return books
}

// GetBookById returns a single book by ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := config.GetDB()
	result := db.Where("id = ?", Id).First(&getBook)
	return &getBook, result
}

// DeleteBook deletes a book by ID
func DeleteBook(ID int64) (*Book, error) {
	db := config.GetDB()

	var book Book
	result := db.First(&book, ID) // fetch book first
	if result.Error != nil {
		return nil, result.Error
	}

	db.Delete(&book) // delete it
	return &book, nil
}

