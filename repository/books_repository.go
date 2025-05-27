package repository

import (
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"gorm.io/gorm"
)

type (
	BooksRepository interface {
		GetAllBooks() ([]*entity.Books, error)
		GetBookByID(id string) (*entity.Books, error)
	}

	bookRepository struct {
		db *gorm.DB
	}
)

func NewBooksRepository(db *gorm.DB) BooksRepository {
	return &bookRepository{
		db: db,
	}
}

func (r *bookRepository) GetAllBooks() ([]*entity.Books, error) {
	var books []*entity.Books
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *bookRepository) GetBookByID(id string) (*entity.Books, error) {
	var book entity.Books
	if err := r.db.First(&book, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}