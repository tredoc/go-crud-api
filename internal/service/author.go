package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-crud-api/internal/cache"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type AuthorService struct {
	repo  repository.Author
	cache cache.RCache
}

func NewAuthorService(repo repository.Author, cache cache.RCache) *AuthorService {
	return &AuthorService{
		repo:  repo,
		cache: cache,
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
	key := fmt.Sprintf("author:%d", id)
	var authorCache types.Author
	err := getFromCache(s.cache.Get, key, &authorCache)
	if err == nil {
		return &authorCache, nil
	}

	author, err := s.repo.GetAuthorByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	go setToCache(s.cache.Set, key, author, cache.EXPIRATION)
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
	key := "authors"
	var authorsCache []*types.Author
	err := getFromCache(s.cache.Get, key, &authorsCache)
	if err == nil {
		return authorsCache, nil
	}

	authors, err := s.repo.GetAllAuthors(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return authors, nil
		}

		return nil, err
	}

	go setToCache(s.cache.Set, key, authors, cache.EXPIRATION)
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
	go s.cache.Invalidate("authors")
	go s.cache.Invalidate(fmt.Sprintf("author:%d", id))
	return existingAuthor, err
}

func (s *AuthorService) DeleteAuthor(ctx context.Context, id int64) error {
	err := s.repo.DeleteAuthor(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrNotFound
		}

		return err
	}
	go s.cache.Invalidate("authors")
	go s.cache.Invalidate(fmt.Sprintf("author:%d", id))
	return nil
}
