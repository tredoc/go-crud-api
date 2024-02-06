package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/pkg/types"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service service.Book
}

func NewBookHandler(service service.Book) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newBook types.CreateBook
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	res, err := h.service.CreateBook(ctx, &newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName(idParam)
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	book, err := h.service.GetBookByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	books, err := h.service.GetAllBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(resp)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.UpdateBook()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName(idParam)
	if idStr == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id < 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.service.DeleteBook(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, _ = fmt.Fprintf(w, "success")
}
