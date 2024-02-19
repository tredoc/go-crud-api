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

func (r *BookRepository) CreateBook(ctx context.Context, book *types.Book) (int64, time.Time, error) {
	var bookID int64
	var createdAt time.Time

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return bookID, createdAt, err
	}
	defer tx.Rollback()

	stmt := `INSERT INTO books(title, publish_date, isbn, pages) VALUES($1, $2, $3, $4) RETURNING id, created_at`
	err = tx.QueryRowContext(ctx, stmt, book.Title, book.PublishDate.Format(time.DateOnly), book.ISBN, book.Pages).Scan(&bookID, &createdAt)
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

		if rowsAffected == 0 {
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
	var customDate time.Time
	var book types.Book
	stmt := `SELECT title, publish_date, created_at, isbn, pages FROM books WHERE id=$1`
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&book.Title, &customDate, &book.CreatedAt, &book.ISBN, &book.Pages)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	stmt = `SELECT author_id FROM book_author WHERE book_id = $1`
	rows, err := r.db.QueryContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}

	var authors []int64
	for rows.Next() {
		var authorID int64
		err := rows.Scan(&authorID)
		if err != nil {
			return nil, err
		}
		authors = append(authors, authorID)
	}

	stmt = `SELECT genre_id FROM book_genre WHERE book_id=$1`
	rows, err = r.db.QueryContext(ctx, stmt, id)
	if err != nil {
		return nil, err
	}

	var genres []int64
	for rows.Next() {
		var genreID int64
		err := rows.Scan(&genreID)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genreID)
	}

	book.PublishDate = types.CustomDate{Time: customDate}
	book.Authors = authors
	book.Genres = genres

	return &book, nil
}

func (r *BookRepository) GetAllBooks(ctx context.Context) ([]*types.Book, error) {
	var books []*types.Book
	stmt := `
		SELECT b.id, b.title, b.publish_date, b.created_at, b.isbn, b.pages, 
		array_agg(DISTINCT ba.author_id) as authors, array_agg(DISTINCT bg.genre_id) as genres 
		FROM books AS b 
		LEFT JOIN book_author AS ba on b.id = ba.book_id 
		LEFT JOIN book_genre AS bg on b.id = bg.book_id
		GROUP BY b.id`

	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var customDate time.Time
		var book types.Book
		var authorsStr string
		var genresStr string
		err := rows.Scan(&book.ID, &book.Title, &customDate, &book.CreatedAt, &book.ISBN, &book.Pages, &authorsStr, &genresStr)
		if err != nil {
			return nil, err
		}

		authors, err := stringToInt64Slice(authorsStr)
		if err != nil {
			authors = []int64{}
		}

		genres, err := stringToInt64Slice(genresStr)
		if err != nil {
			genres = []int64{}
		}

		book.PublishDate = types.CustomDate{Time: customDate}
		book.Authors = authors
		book.Genres = genres
		books = append(books, &book)
	}

	return books, nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, id int64, book *types.Book) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `UPDATE books SET title = $1, publish_date = $2, isbn = $3, pages = $4 WHERE id = $5`
	_, err = tx.ExecContext(ctx, stmt, book.Title, book.PublishDate.Format(time.DateOnly), book.ISBN, book.Pages, id)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM book_author WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO book_author(book_id, author_id) VALUES($1, $2)`
	for _, authorID := range book.Authors {
		_, err = tx.ExecContext(ctx, stmt, id, authorID)
		if err != nil {
			return err
		}
	}

	stmt = `DELETE FROM book_genre WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO book_genre(book_id, genre_id) VALUES($1, $2)`
	for _, genreID := range book.Genres {
		_, err = tx.ExecContext(ctx, stmt, id, genreID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `DELETE FROM book_genre WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM book_author WHERE book_id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	stmt = `DELETE FROM books WHERE id = $1`
	_, err = tx.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
