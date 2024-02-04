package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type GenreService struct {
	repo repository.Genre
}

func NewGenreService(repo repository.Genre) *GenreService {
	return &GenreService{
		repo: repo,
	}
}

func (s *GenreService) CreateGenre(ctx context.Context, genre *types.Genre) (*types.Genre, error) {
	id, err := s.repo.CreateGenre(ctx, genre)
	if err != nil {
		if errors.Is(err, repository.ErrEntityExists) {
			return nil, ErrEntityExists
		}
		return nil, err
	}

	genre.ID = id
	return genre, nil
}

func (s *GenreService) GetAllGenres(ctx context.Context) ([]*types.Genre, error) {
	genres, err := s.repo.GetAllGenres(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return genres, nil
		}

		return nil, err
	}

	return genres, nil
}
