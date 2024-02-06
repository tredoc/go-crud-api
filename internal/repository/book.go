package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/tredoc/go-crud-api/pkg/types"
	"time"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateBook(ctx context.Context, book *types.CreateBook) (int64, time.Time, error) {
	var bookID int64
	var createdAt time.Time

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return bookID, createdAt, err
	}
	defer tx.Rollback()

	stmt := `INSERT INTO books(title, publish_date, isbn, pages) VALUES($1, $2, $3, $4) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, stmt, &book.Title, &book.PublishDate, &book.ISBN, &book.Pages).Scan(&bookID, &createdAt)
	if err != nil {
		return bookID, createdAt, err
	}

	stmt = `INSERT INTO book_author(book_id, author_id) VALUES($1, $2)`
	for _, authorID := range book.Authors {
		res, err := tx.ExecContext(ctx, stmt, bookID, authorID)
		if err != nil {
			return bookID, createdAt, err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return bookID, createdAt, err
		}

		if rowsAffected < 0 {
			return bookID, createdAt, errors.New("no rows affected on book_author relation create")
		}
	}

	stmt = `INSERT INTO book_genre(book_id, genre_id) VALUES($1, $2)`
	for _, genreID := range book.Genres {
		res, err := tx.ExecContext(ctx, stmt, bookID, genreID)
		if err != nil {
			return genreID, createdAt, err
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return genreID, createdAt, err
		}

		if rowsAffected < 0 {
			return genreID, createdAt, errors.New("no rows affected on book_author relation create")
		}
	}

	err = tx.Commit()
	return bookID, createdAt, err
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
