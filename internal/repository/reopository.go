package repository

import "database/sql"

type Book interface {
	CreateBook() (string, error)
	GetBookByID() (string, error)
	GetAllBooks() (string, error)
	UpdateBook() (string, error)
	DeleteBook() (string, error)
}

type Repository struct {
	Book Book
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book: NewBookRepository(db),
	}
}
