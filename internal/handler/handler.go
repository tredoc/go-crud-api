package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"net/http"
)

type Book interface {
	CreateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
	GetBookByID(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
	GetAllBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
	UpdateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
	DeleteBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params)
}

type Handler struct {
	book Book
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		book: NewBookHandler(services.Book),
	}
}

func (h *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/v1/books", h.book.CreateBook)
	router.GET("/api/v1/books", h.book.GetAllBooks)
	router.GET("/api/v1/books/:id", h.book.GetBookByID)
	router.PUT("/api/v1/books/:id", h.book.UpdateBook)
	router.DELETE("/api/v1/books/:id", h.book.DeleteBook)

	return router
}
