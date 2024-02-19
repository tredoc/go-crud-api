package types

import (
	"time"
)

type Book struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	ISBN        string    `json:"isbn"`
	Pages       uint16    `json:"pages"`
	Authors     []int64   `json:"authors"`
	Genres      []int64   `json:"genres"`
}

type BookWithDetails struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	ISBN        string    `json:"isbn"`
	Pages       uint16    `json:"pages"`
	Authors     []*Author `json:"authors"`
	Genres      []*Genre  `json:"genres"`
}

type CreateBook struct {
	Title       string  `json:"title"`
	PublishDate string  `json:"publish_date"`
	ISBN        string  `json:"isbn"`
	Pages       uint16  `json:"pages"`
	Authors     []int64 `json:"authors"`
	Genres      []int64 `json:"genres"`
}

type UpdateBook struct {
	Title       *string    `json:"title"`
	PublishDate *time.Time `json:"publish_date"`
	ISBN        *string    `json:"isbn"`
	Pages       *uint16    `json:"pages"`
	Authors     []int64    `json:"authors"`
	Genres      []int64    `json:"genres"`
}
