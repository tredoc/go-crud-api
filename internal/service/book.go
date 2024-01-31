package service

import "github.com/tredoc/go-crud-api/internal/repository"

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook() (string, error) {
	return s.repo.CreateBook()
}

func (s *BookService) GetBookByID() (string, error) {
	return s.repo.GetBookByID()
}

func (s *BookService) GetAllBooks() (string, error) {
	return s.repo.GetAllBooks()
}

func (s *BookService) UpdateBook() (string, error) {
	return s.repo.UpdateBook()
}

func (s *BookService) DeleteBook() (string, error) {
	return s.repo.DeleteBook()
}
