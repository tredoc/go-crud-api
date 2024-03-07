package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/tredoc/go-crud-api/pkg/types"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, email string, hash []byte) (int64, time.Time, error) {
	stmt := `SELECT id FROM users WHERE email = $1`
	var foundUserID int64
	err := r.db.QueryRowContext(ctx, stmt, email).Scan(&foundUserID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, time.Time{}, err
	}
	if foundUserID != 0 {
		return 0, time.Time{}, ErrEntityExists
	}

	var id int64
	var createdAt time.Time
	stmt = `INSERT INTO users(email, password) VALUES($1, $2) returning id, created_at`
	err = r.db.QueryRowContext(ctx, stmt, email, hash).Scan(&id, &createdAt)
	if err != nil {
		return 0, createdAt, err
	}

	return id, createdAt, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*types.User, string, error) {
	stmt := `SELECT id, email, created_at, password FROM users WHERE email = $1`
	var user types.User
	var password string
	err := r.db.QueryRowContext(ctx, stmt, email).Scan(&user.ID, &user.Email, &user.CreatedAt, &password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "", ErrNotFound
		}

		return nil, "", err
	}

	return &user, password, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*types.User, error) {
	stmt := `SELECT id, email, created_at, role FROM users WHERE id = $1`
	var user types.User
	err := r.db.QueryRowContext(ctx, stmt, id).Scan(&user.ID, &user.Email, &user.CreatedAt, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return &user, nil
}
