package repository

import (
	"context"
	"database/sql"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type Book interface {
	CreateBook(ctx context.Context, book *types.Book) (int64, error)
	GetBookByID(context.Context, int64) (*types.Book, error)
	GetAllBooks(context.Context) ([]*types.Book, error)
	UpdateBook() (string, error)
	DeleteBook() (string, error)
}

type Author interface {
	CreateAuthor(context.Context, *types.Author) (int64, error)
	GetAuthorByID(context.Context, int64) (*types.Author, error)
	GetAuthorByName(context.Context, string, string) (*types.Author, error)
	GetAllAuthors(context.Context) ([]*types.Author, error)
}

type Genre interface {
	CreateGenre(context.Context, *types.Genre) (int64, error)
	GetGenreByID(context.Context, int64) (*types.Genre, error)
	GetAllGenres(context.Context) ([]*types.Genre, error)
}

type Repository struct {
	Book
	Author
	Genre
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book:   NewBookRepository(db),
		Author: NewAuthorRepository(db),
		Genre:  NewGenreRepository(db),
	}
}
