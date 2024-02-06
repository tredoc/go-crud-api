package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
	"strings"
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
	genres, err := s.repo.GetAllGenres(ctx)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	}

	for _, g := range genres {
		if g.Name == strings.ToLower(genre.Name) {
			return g, ErrEntityExists
		}
	}

	id, err := s.repo.CreateGenre(ctx, genre)
	if err != nil {
		return nil, err
	}

	genre.ID = id
	return genre, nil
}

func (s *GenreService) GetGenreByID(ctx context.Context, id int64) (*types.Genre, error) {
	genre, err := s.repo.GetGenreByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return genre, ErrNotFound
		}

		return nil, err
	}

	return genre, nil
}

func (s *GenreService) GetGenresByIDs(ctx context.Context, ids []int64) ([]*types.Genre, error) {
	genres, err := s.repo.GetGenresByIDs(ctx, ids)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return genres, nil
		}

		return nil, err
	}

	return genres, nil
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
