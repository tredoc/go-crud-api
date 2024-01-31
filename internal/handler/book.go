package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookHandler struct{}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("create book"))
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("get book by id"))
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("get all books"))
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("update book"))
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Write([]byte("delete book"))
}
