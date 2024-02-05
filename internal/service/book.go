package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) CreateBook(ctx context.Context, book *types.Book) (*types.Book, error) {
	id, err := s.repo.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	book.ID = id
	return book, nil
}

func (s *BookService) GetBookByID(ctx context.Context, id int64) (*types.Book, error) {
	book, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return book, nil
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*types.Book, error) {
	books, err := s.repo.GetAllBooks(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return books, nil
		}

		return nil, err
	}

	return books, nil
}

func (s *BookService) UpdateBook() (string, error) {
	return s.repo.UpdateBook()
}

func (s *BookService) DeleteBook(ctx context.Context, id int64) error {
	return s.repo.DeleteBook(ctx, id)
}
