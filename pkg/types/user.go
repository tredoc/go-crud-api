package types

import (
	"errors"
	"github.com/tredoc/go-crud-api/internal/validator"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

const COST = 10

type User struct {
	ID        int64     `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Email     string    `json:"email"`
}

type Password struct {
	Plaintext *string
	Hash      []byte
}

func (p *Password) Set(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		return err
	}
	p.Plaintext = &password
	p.Hash = hash

	return nil
}

func (p *Password) Matches(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AccessToken string

func ValidateRegisterUser(v *validator.Validator, user *AuthUser) {
	v.Check(len(user.Email) > 0, "email", validator.CantBeEmpty)
	v.Check(v.Matches(user.Email, regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)), "email", "should look like example@example.com")
	v.Check(len(user.Password) >= 6, "password", validator.CantBeShorterThan6)
	v.Check(v.Matches(user.Password, regexp.MustCompile(`[A-Za-z]+`)), "password", "must contain at least one letter")
	v.Check(v.Matches(user.Password, regexp.MustCompile(`\d+`)), "password", "must contain at least one number")
	v.Check(v.Matches(user.Password, regexp.MustCompile(`[@$!%*#?&]+`)), "password", "must contain at least one special character")
}

func ValidateLoginUser(v *validator.Validator, user *AuthUser) {
	v.Check(len(user.Email) > 0, "email", validator.CantBeEmpty)
	v.Check(v.Matches(user.Email, regexp.MustCompile(`^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)), "email", "should look like example@example.com")
	v.Check(len(user.Password) > 0, "password", validator.CantBeEmpty)
	v.Check(len(user.Password) >= 6, "email", validator.CantBeShorterThan6)
}
