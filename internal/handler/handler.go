package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"net/http"
)

type Book interface {
	CreateBook(http.ResponseWriter, *http.Request, httprouter.Params)
	GetBookByID(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllBooks(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateBook(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteBook(http.ResponseWriter, *http.Request, httprouter.Params)
}

type Genre interface {
	CreateGenre(http.ResponseWriter, *http.Request, httprouter.Params)
	GetGenreByID(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllGenres(http.ResponseWriter, *http.Request, httprouter.Params)
}

type Handler struct {
	book  Book
	genre Genre
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		book:  NewBookHandler(services.Book),
		genre: NewGenreHandler(services.Genre),
	}
}

func (h *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/books", h.book.CreateBook)
	router.GET("/api/v1/books", h.book.GetAllBooks)
	router.GET("/api/v1/books/:id", h.book.GetBookByID)
	router.PUT("/api/v1/books/:id", h.book.UpdateBook)
	router.DELETE("/api/v1/books/:id", h.book.DeleteBook)

	router.GET("/api/v1/genres", h.genre.GetAllGenres)
	router.GET("/api/v1/genres/:id", h.genre.GetGenreByID)
	router.POST("/api/v1/genres", h.genre.CreateGenre)

	return router
}
