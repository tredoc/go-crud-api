package handler

import "github.com/julienschmidt/httprouter"

type Handler struct {
	book *BookHandler
}

func NewHandler() *Handler {
	return &Handler{
		book: NewBookHandler(),
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
