package repository

import (
	"context"
	"database/sql"
	"errors"
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
	var id int64
	stmt := `INSERT INTO books(title, publish_date, created_at, isbn, pages) VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRowContext(ctx, stmt, &book.ID, &book.Title, &book.PublishDate, &book.CreatedAt, &book.ISBN, &book.Pages).Scan(&id)
	return id, err
}

func (r *BookRepository) GetBookByID(ctx context.Context, id int64) (*types.Book, error) {
	var book types.Book
	stmt := `SELECT title, publish_date, created_at, isbn, pages FROM books WHERE id=$1`
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&book.ID, &book.Title, &book.PublishDate, &book.CreatedAt, &book.ISBN, &book.Pages)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &book, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*types.Book, error) {
	var books []*types.Book
	stmt := `SELECT title, publish_date, created_at, isbn, pages FROM books`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.ID, &book.Title, &book.PublishDate, &book.CreatedAt, &book.ISBN, &book.Pages)
		if err != nil {
			return nil, err
		}
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookRepository) UpdateBook() (string, error) {
	return "update Book", nil
}

func (r *BookRepository) DeleteBook(_ context.Context, _ int64) error {
	return nil
}
