package service

import "github.com/tredoc/go-crud-api/internal/repository"

type Book interface {
	CreateBook() (string, error)
	GetBookByID() (string, error)
	GetAllBooks() (string, error)
	UpdateBook() (string, error)
	DeleteBook() (string, error)
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Book),
	}
}
