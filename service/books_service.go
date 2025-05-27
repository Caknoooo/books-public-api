package service

import (
	"github.com/Caknoooo/go-gin-clean-starter/entity"
	"github.com/Caknoooo/go-gin-clean-starter/repository"
)

type (
	BooksService interface {
		GetAllBooks() ([]*entity.Books, error)
		GetBookByID(id string) (*entity.Books, error)
	}

	bookService struct {
		booksRepository repository.BooksRepository
	}
)

func NewBooksService(booksRepository repository.BooksRepository) BooksService {
	return &bookService{
		booksRepository: booksRepository,
	}
}

func (s *bookService) GetAllBooks() ([]*entity.Books, error) {
	books, err := s.booksRepository.GetAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *bookService) GetBookByID(id string) (*entity.Books, error) {
	book, err := s.booksRepository.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
