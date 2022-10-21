package services

import (
	"simple-app/database"
	"simple-app/models"
)

type BookService struct{}

func (s *BookService) GetAll() []models.Book {
	var books []models.Book

	if err := database.DB.Find(&books).Error; err != nil {
		panic(err)
	}

	return books
}

func (s *BookService) GetByID(id string) models.Book {
	var book models.Book

	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		return models.Book{}
	}

	return book
}

func (s *BookService) Create(book models.Book) models.Book {

	if err := database.DB.Create(&book).Error; err != nil {
		panic(err)
	}

	var createdBook models.Book

	database.DB.Last(&createdBook)

	return createdBook
}

func (s *BookService) Update(id string, book models.Book) models.Book {
	var updatedBook models.Book = s.GetByID(id)

	updatedBook.Title = book.Title
	updatedBook.Publisher = book.Publisher

	if err := database.DB.Save(&updatedBook).Error; err != nil {
		panic(err)
	}

	return updatedBook
}

func (s *BookService) Delete(id string) bool {
	var book models.Book = s.GetByID(id)

	result := database.DB.Delete(&book)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
