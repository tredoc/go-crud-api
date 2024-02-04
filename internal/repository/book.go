package repository

import (
	"context"
	"database/sql"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *types.Book) (int64, error) {
	return 0, nil
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int64) (*types.Book, error) {
	return &types.Book{}, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*types.Book, error) {
	return []*types.Book{&types.Book{}}, nil
}

func (r *BookRepository) UpdateBook() (string, error) {
	return "update Book", nil
}

func (r *BookRepository) DeleteBook() (string, error) {
	return "delete book", nil
}
