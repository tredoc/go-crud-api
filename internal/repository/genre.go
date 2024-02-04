package repository

import (
	"context"
	"database/sql"
	"github.com/tredoc/go-crud-api/pkg/types"
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
	err := r.db.QueryRowContext(ctx, stmt, genre).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *GenreRepository) GetAllGenres(ctx context.Context) ([]*types.Genre, error) {
	stmt := `SELECT * FROM genres`
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
