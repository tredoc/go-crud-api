package types

import (
	"fmt"
	"github.com/tredoc/go-crud-api/internal/validator"
	"strconv"
	"time"
)

const layout = time.DateOnly

type CustomDate struct {
	time.Time
}

func (cd *CustomDate) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	date, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	cd.Time = date
	return nil
}

func (cd *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, cd.Time.Format(layout))), nil
}

type Book struct {
	ID          int64      `json:"id,omitempty"`
	Title       string     `json:"title"`
	PublishDate CustomDate `json:"publish_date"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	ISBN        string     `json:"isbn"`
	Pages       uint16     `json:"pages"`
	Authors     []int64    `json:"authors"`
	Genres      []int64    `json:"genres"`
}

func ValidateBook(v *validator.Validator, book *Book) {
	v.Check(book.Title != "", "title", validator.CantBeEmpty)
	v.Check(book.PublishDate.Before(time.Now()), "publish_date", validator.OnlyInThePast)
	v.Check(book.ISBN != "", "isbn", validator.CantBeEmpty)
	v.Check(book.Pages > 0, "pages", validator.CantBeLessThanOne)
	v.Check(book.Pages <= 5000, "pages", validator.CantBeBiggerThan5k)
	v.Check(len(book.Authors) > 0, "authors", validator.CantBeEmpty)
	v.Check(len(book.Genres) > 0, "genres", validator.CantBeEmpty)
}

type BookWithDetails struct {
	ID          int64      `json:"id,omitempty"`
	Title       string     `json:"title"`
	PublishDate CustomDate `json:"publish_date"`
	CreatedAt   time.Time  `json:"created_at"`
	ISBN        string     `json:"isbn"`
	Pages       uint16     `json:"pages"`
	Authors     []*Author  `json:"authors"`
	Genres      []*Genre   `json:"genres"`
}

type UpdateBook struct {
	Title       *string     `json:"title"`
	PublishDate *CustomDate `json:"publish_date"`
	ISBN        *string     `json:"isbn"`
	Pages       *uint16     `json:"pages"`
	Authors     []int64     `json:"authors"`
	Genres      []int64     `json:"genres"`
}

func ValidateUpdateBook(v *validator.Validator, book *UpdateBook) {
	if book.Title != nil {
		v.Check(*book.Title != "", "title", validator.CantBeEmpty)
	}

	if book.PublishDate != nil {
		v.Check(book.PublishDate.Before(time.Now()), "publish_date", validator.OnlyInThePast)
	}

	if book.ISBN != nil {
		v.Check(*book.ISBN != "", "isbn", validator.CantBeEmpty)
	}

	if book.Pages != nil {
		v.Check(*book.Pages > 0, "pages", validator.CantBeLessThanOne)
		v.Check(*book.Pages <= 5000, "pages", validator.CantBeBiggerThan5k)
	}

	if book.Authors != nil {
		v.Check(len(book.Authors) > 0, "authors", validator.CantBeEmpty)
	}

	if book.Genres != nil {
		v.Check(len(book.Genres) > 0, "genres", validator.CantBeEmpty)
	}
}
