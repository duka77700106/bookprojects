package services

import (
	"bookProject/internal/models"
	"bookProject/internal/repository"
)

type BookService interface {
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(id uint) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAll() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetByID(id uint) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) Create(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) Update(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) Delete(id uint) error {
	return s.repo.Delete(id)
}
