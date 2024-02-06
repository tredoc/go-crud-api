package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/tredoc/go-crud-api/pkg/types"
	"strings"
)

type GenreRepository struct {
	db *sql.DB
}

func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (r *GenreRepository) CreateGenre(ctx context.Context, genre *types.Genre) (int64, error) {
	stmt := `INSERT INTO genres (name) VALUES ($1) RETURNING id`
	var id int64
	err := r.db.QueryRowContext(ctx, stmt, genre.Name).Scan(&id)
	return id, err
}

func (r *GenreRepository) GetGenreByID(ctx context.Context, id int64) (*types.Genre, error) {
	stmt := `SELECT id, name FROM genres WHERE id = $1`
	var genre types.Genre
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&genre.ID, &genre.Name)
	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (r *GenreRepository) GetGenresByIDs(ctx context.Context, ids []int64) ([]*types.Genre, error) {
	placeholders := make([]string, len(ids))
	for idx := range ids {
		placeholders[idx] = fmt.Sprintf("$%d", idx+1)
	}

	placeholder := strings.Join(placeholders, ",")
	stmt := fmt.Sprintf(`SELECT id, name FROM genres WHERE id IN (%s)`, placeholder)
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := r.db.QueryContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []*types.Genre
	for rows.Next() {
		var genre types.Genre
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &genre)
	}

	return genres, nil
}

func (r *GenreRepository) GetAllGenres(ctx context.Context) ([]*types.Genre, error) {
	stmt := `SELECT id, name FROM genres`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var genres []*types.Genre

	for rows.Next() {
		var genre types.Genre
		err := rows.Scan(&genre.ID, &genre.Name)
		if err != nil {
			return nil, err
		}
		genres = append(genres, &genre)
	}

	return genres, nil
}

func (r *GenreRepository) UpdateGenre(ctx context.Context, id int64, genre *types.Genre) error {
	stmt := `UPDATE genres SET name = $1 WHERE id = $2`
	res, err := r.db.ExecContext(ctx, stmt, genre.Name, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
