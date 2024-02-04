package service

import (
	"context"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type Book interface {
	CreateBook(context.Context, *types.Book) (*types.Book, error)
	GetBookByID(context.Context, int64) (*types.Book, error)
	GetAllBooks(context.Context) ([]*types.Book, error)
	UpdateBook() (string, error)
	DeleteBook() (string, error)
}

type Author interface {
	CreateAuthor(ctx context.Context, author *types.Author) (*types.Author, error)
	GetAuthorByID(ctx context.Context, id int64) (*types.Author, error)
	GetAuthorByName(ctx context.Context, firstName string, lastName string) (*types.Author, error)
	GetAllAuthors(ctx context.Context) ([]*types.Author, error)
}

type Genre interface {
	CreateGenre(context.Context, *types.Genre) (*types.Genre, error)
	GetAllGenres(context.Context) ([]*types.Genre, error)
}

type Service struct {
	Book
	Author
	Genre
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book:   NewBookService(repos.Book),
		Author: NewAuthorService(repos.Author),
		Genre:  NewGenreService(repos.Genre),
	}
}
