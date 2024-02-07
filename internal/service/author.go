package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type AuthorService struct {
	repo repository.Author
}

func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *types.Author) (*types.Author, error) {
	id, err := s.repo.CreateAuthor(ctx, author)
	if err != nil {
		if errors.Is(err, repository.ErrEntityExists) {
			return nil, ErrEntityExists
		}

		return nil, err
	}

	author.ID = id
	return author, nil
}

func (s *AuthorService) GetAuthorByID(ctx context.Context, id int64) (*types.Author, error) {
	author, err := s.repo.GetAuthorByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return author, nil
}

func (s *AuthorService) GetAuthorsByIDs(ctx context.Context, ids []int64) ([]*types.Author, error) {
	authors, err := s.repo.GetAuthorsByIDs(ctx, ids)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return authors, nil
		}

		return nil, err
	}

	return authors, nil
}

func (s *AuthorService) GetAuthorByName(ctx context.Context, firstName string, lastName string) (*types.Author, error) {
	author, err := s.repo.GetAuthorByName(ctx, firstName, lastName)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return author, nil
}

func (s *AuthorService) GetAllAuthors(ctx context.Context) ([]*types.Author, error) {
	authors, err := s.repo.GetAllAuthors(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return authors, nil
		}

		return nil, err
	}

	return authors, nil
}

func (s *AuthorService) UpdateAuthor(ctx context.Context, id int64, author *types.UpdateAuthor) (*types.Author, error) {
	existingAuthor, err := s.repo.GetAuthorByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if author.FirstName != nil {
		existingAuthor.FirstName = *author.FirstName
	}

	if author.MiddleName != nil {
		existingAuthor.MiddleName = *author.MiddleName
	}

	if author.LastName != nil {
		existingAuthor.LastName = *author.LastName
	}

	err = s.repo.UpdateAuthor(ctx, id, existingAuthor)
	return existingAuthor, err
}
