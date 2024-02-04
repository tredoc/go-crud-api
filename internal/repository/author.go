package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, author *types.Author) (int64, error) {
	stmt := `SELECT id, first_name, middle_name, last_name FROM authors WHERE first_name = $1 AND last_name = $2`
	row := r.db.QueryRowContext(ctx, stmt, author.FirstName, author.LastName)

	var foundAuthorID int64
	err := row.Scan(&foundAuthorID)
	if err != nil && !errors.Is(row.Err(), sql.ErrNoRows) {
		return 0, err
	}

	if foundAuthorID != 0 {
		return 0, ErrEntityExists
	}

	var id int64
	stmt = `INSERT INTO authors (first_name, middle_name, last_name) VALUES ($1, $2, $3) RETURNING id`
	err = r.db.QueryRowContext(ctx, stmt, author.FirstName, author.MiddleName, author.LastName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthorRepository) GetAuthorByID(ctx context.Context, id int64) (*types.Author, error) {
	stmt := `SELECT id, first_name, middle_name, last_name FROM authors WHERE id = $1`
	row := r.db.QueryRowContext(ctx, stmt, id)

	var author types.Author
	err := row.Scan(&author.ID, &author.FirstName, &author.MiddleName, &author.LastName)
	if err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) GetAuthorByName(ctx context.Context, firstName string, lastName string) (*types.Author, error) {
	stmt := `SELECT id, first_name, middle_name, last_name FROM authors WHERE first_name = $1 AND last_name = $2`
	row := r.db.QueryRowContext(ctx, stmt, firstName, lastName)

	var author types.Author
	err := row.Scan(&author.ID, &author.FirstName, &author.MiddleName, &author.LastName)
	if err != nil {
		if errors.Is(row.Err(), sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &author, nil
}

func (r *AuthorRepository) GetAllAuthors(ctx context.Context) ([]*types.Author, error) {
	stmt := `SELECT id, first_name, middle_name, last_name FROM authors`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*types.Author
	for rows.Next() {
		var author types.Author
		err := rows.Scan(&author.ID, &author.FirstName, &author.MiddleName, &author.LastName)
		if err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}
