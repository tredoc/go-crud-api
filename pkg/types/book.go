package types

import "time"

type Book struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title"`
	PublishDate time.Time `json:"publish_date"`
	CreatedAt   time.Time `json:"created_at"`
	ISBN        string    `json:"isbn"`
	Pages       uint16    `json:"pages"`
	Authors     []Author  `json:"author"`
	Genres      []Genre   `json:"genre"`
}
