package types

import (
	"github.com/tredoc/go-crud-api/internal/validator"
	"regexp"
)

type Genre struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

func ValidateGenre(v *validator.Validator, genre *Genre) {
	v.Check(len(genre.Name) > 0, "name", "can't be empty")
	v.Check(v.Matches(genre.Name, regexp.MustCompile(`^[a-zA-Z]+$`)), "name", "should contain only latin letters")
}
