package repository

import (
	"bookProject/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}
