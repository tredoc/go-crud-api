package service

import (
	"context"
	"errors"
	"github.com/pascaldekloe/jwt"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
	"log"
	"os"
	"strconv"
	"time"
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
	user, pwd, err := s.repo.GetUserByEmail(ctx, authUser.Email)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", ErrNotFound
		}
		return "", err
	}

	password := types.Password{
		Hash: []byte(pwd),
	}
	isMatch, err := password.Matches(authUser.Password)
	if err != nil {
		return "", err
	}

	if !isMatch {
		return "", ErrCredentialsMismatch
	}

	var claims jwt.Claims
	claims.Subject = strconv.FormatInt(user.ID, 10)
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(types.EXPIRATION))
	claims.Issuer = "go-crud-api"
	claims.Audiences = []string{"go-crud-api"}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Panic("secret is unavailable")
	}
	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(secret))
	if err != nil {
		return "", err
	}

	return types.AccessToken(jwtBytes), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int64) (*types.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return user, nil
}
