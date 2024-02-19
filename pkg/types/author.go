package types

import (
	"github.com/tredoc/go-crud-api/internal/validator"
	"regexp"
)

type Author struct {
	ID         int64  `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

type UpdateAuthor struct {
	FirstName  *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   *string `json:"last_name"`
}

func ValidateAuthor(v *validator.Validator, author *Author) {
	v.Check(author.FirstName != "", "first_name", validator.CantBeEmpty)
	v.Check(v.Matches(author.FirstName, regexp.MustCompile(`^[a-zA-Z]+$`)), "first_name", validator.OnlyLatinLetters)

	if author.MiddleName != "" {
		v.Check(v.Matches(author.MiddleName, regexp.MustCompile(`^[a-zA-Z]+$`)), "middle_name", validator.OnlyLatinLetters)
	}

	v.Check(author.LastName != "", "last_name", validator.CantBeEmpty)
	v.Check(v.Matches(author.LastName, regexp.MustCompile(`^[a-zA-Z]+$`)), "last_name", validator.OnlyLatinLetters)
}

func ValidateUpdateAuthor(v *validator.Validator, author *UpdateAuthor) {
	if author.FirstName != nil {
		v.Check(*author.FirstName != "", "first_name", validator.CantBeEmpty)
		v.Check(v.Matches(*author.FirstName, regexp.MustCompile(`^[a-zA-Z]+$`)), "first_name", validator.OnlyLatinLetters)
	}

	if author.MiddleName != nil && *author.MiddleName != "" {
		v.Check(v.Matches(*author.MiddleName, regexp.MustCompile(`^[a-zA-Z]+$`)), "middle_name", validator.OnlyLatinLetters)
	}

	if author.LastName != nil {
		v.Check(*author.LastName != "", "last_name", validator.CantBeEmpty)
		v.Check(v.Matches(*author.LastName, regexp.MustCompile(`^[a-zA-Z]+$`)), "last_name", validator.OnlyLatinLetters)
	}
}
