package service

import (
	"context"
	"github.com/tredoc/go-crud-api/internal/cache"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type Book interface {
	CreateBook(context.Context, *types.Book) (*types.BookWithDetails, error)
	GetBookByID(context.Context, int64) (*types.BookWithDetails, error)
	GetAllBooks(context.Context) ([]*types.Book, error)
	UpdateBook(context.Context, int64, *types.UpdateBook) (*types.Book, error)
	DeleteBook(context.Context, int64) error
}

type Genre interface {
	CreateGenre(context.Context, *types.Genre) (*types.Genre, error)
	GetGenreByID(context.Context, int64) (*types.Genre, error)
	GetAllGenres(context.Context) ([]*types.Genre, error)
	UpdateGenre(context.Context, int64, *types.Genre) error
	DeleteGenre(context.Context, int64) error
}

type Author interface {
	CreateAuthor(context.Context, *types.Author) (*types.Author, error)
	GetAuthorByID(context.Context, int64) (*types.Author, error)
	GetAuthorsByIDs(context.Context, []int64) ([]*types.Author, error)
	GetAuthorByName(context.Context, string, string) (*types.Author, error)
	GetAllAuthors(context.Context) ([]*types.Author, error)
	UpdateAuthor(context.Context, int64, *types.UpdateAuthor) (*types.Author, error)
	DeleteAuthor(context.Context, int64) error
}

type User interface {
	RegisterUser(context.Context, *types.AuthUser) (*types.User, error)
	LoginUser(context.Context, *types.AuthUser) (types.AccessToken, error)
	GetUserByID(context.Context, int64) (*types.User, error)
}

type Service struct {
	Book
	Author
	Genre
	User
}

func NewService(repos *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Book:   NewBookService(repos.Book, repos.Author, repos.Genre, cache.Redis),
		Genre:  NewGenreService(repos.Genre, cache.Redis),
		Author: NewAuthorService(repos.Author, cache.Redis),
		User:   NewUserService(repos.User),
	}
}
