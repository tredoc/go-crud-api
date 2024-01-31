package handler

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"net/http"
)

type BookHandler struct {
	service service.Book
}

func NewBookHandler(service service.Book) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.CreateBook()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.GetBookByID()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.GetAllBooks()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.UpdateBook()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.DeleteBook()
	_, _ = fmt.Fprintf(w, res)
}
