package service

import (
	"context"
	"errors"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repo: repository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, authUser *types.AuthUser) (*types.User, error) {
	var password types.Password
	err := password.Set(authUser.Password)
	if err != nil {
		return nil, ErrCantHandleCredentials
	}

	id, createdAt, err := s.repo.CreateUser(ctx, authUser.Email, password.Hash)
	if err != nil {
		if errors.Is(err, repository.ErrEntityExists) {
			return nil, ErrEntityExists
		}
		return nil, err
	}

	newUser := types.User{
		ID:        id,
		CreatedAt: createdAt,
		Email:     authUser.Email,
	}

	return &newUser, nil
}

func (s *UserService) LoginUser(ctx context.Context, authUser *types.AuthUser) (types.AccessToken, error) {
	_, pwd, err := s.repo.GetUserByEmail(ctx, authUser.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}
		return "", err
	}

	password := types.Password{
		Hash: []byte(pwd),
	}
	if err != nil {
		return "", ErrCantHandleCredentials
	}

	isMatch, err := password.Matches(authUser.Password)
	if err != nil {
		return "", err
	}

	if !isMatch {
		return "", ErrCredentialsMismatch
	}

	return "generated access token", nil
}
