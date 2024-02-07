package repository

import (
	"context"
	"database/sql"
	"github.com/tredoc/go-crud-api/pkg/types"
	"time"
)

type Book interface {
	CreateBook(ctx context.Context, book *types.CreateBook) (int64, time.Time, error)
	GetBookByID(context.Context, int64) (*types.Book, error)
	GetAllBooks(context.Context) ([]*types.Book, error)
	UpdateBook() (string, error)
	DeleteBook(context.Context, int64) error
}

type Genre interface {
	CreateGenre(context.Context, *types.Genre) (int64, error)
	GetGenreByID(context.Context, int64) (*types.Genre, error)
	GetGenresByIDs(context.Context, []int64) ([]*types.Genre, error)
	GetAllGenres(context.Context) ([]*types.Genre, error)
	UpdateGenre(context.Context, int64, *types.Genre) error
	DeleteGenre(context.Context, int64) error
}

type Author interface {
	CreateAuthor(context.Context, *types.Author) (int64, error)
	GetAuthorByID(context.Context, int64) (*types.Author, error)
	GetAuthorsByIDs(context.Context, []int64) ([]*types.Author, error)
	GetAuthorByName(context.Context, string, string) (*types.Author, error)
	GetAllAuthors(context.Context) ([]*types.Author, error)
	UpdateAuthor(context.Context, int64, *types.Author) error
	DeleteAuthor(context.Context, int64) error
}

type Repository struct {
	Book
	Genre
	Author
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Book:   NewBookRepository(db),
		Genre:  NewGenreRepository(db),
		Author: NewAuthorRepository(db),
	}
}
