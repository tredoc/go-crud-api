package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
	"time"
)

type BookService struct {
	repo       repository.Book
	authorRepo repository.Author
	genreRepo  repository.Genre
}

func NewBookService(bookRepo repository.Book, authorRepo repository.Author, genreRepo repository.Genre) *BookService {
	return &BookService{
		repo:       bookRepo,
		authorRepo: authorRepo,
		genreRepo:  genreRepo,
	}
}

func (s *BookService) CreateBook(ctx context.Context, book *types.CreateBook) (*types.BookWithDetails, error) {
	id, createdAt, err := s.repo.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	authors, err := s.authorRepo.GetAuthorsByIDs(ctx, book.Authors)
	if err != nil {
		return nil, err
	}

	genres, err := s.genreRepo.GetGenresByIDs(ctx, book.Genres)
	if err != nil {
		return nil, err
	}

	convertedTime, err := time.Parse("2006-01-02", book.PublishDate)
	if err != nil {
		return nil, err
	}
	newBook := types.BookWithDetails{
		ID:          id,
		Title:       book.Title,
		PublishDate: convertedTime,
		CreatedAt:   createdAt,
		ISBN:        book.ISBN,
		Pages:       book.Pages,
		Authors:     authors,
		Genres:      genres,
	}

	return &newBook, nil
}

func (s *BookService) GetBookByID(ctx context.Context, id int64) (*types.BookWithDetails, error) {
	book, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	authors, err := s.authorRepo.GetAuthorsByIDs(ctx, book.Authors)
	if err != nil {
		return nil, err
	}

	genres, err := s.genreRepo.GetGenresByIDs(ctx, book.Genres)
	if err != nil {
		return nil, err
	}

	newBook := types.BookWithDetails{
		ID:          id,
		Title:       book.Title,
		PublishDate: book.PublishDate,
		CreatedAt:   book.CreatedAt,
		ISBN:        book.ISBN,
		Pages:       book.Pages,
		Authors:     authors,
		Genres:      genres,
	}

	return &newBook, nil
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
