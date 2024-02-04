package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"net/http"
	"strconv"
)

var idParam string = "id"

type BookHandler struct {
	service service.Book
}

func NewBookHandler(service service.Book) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	//ctx := r.Context()
	//res, _ := h.service.CreateBook(ctx, &types.Book{})
	_, _ = fmt.Fprint(w, "create book")
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
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

	book, err := h.service.GetBookByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(res)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	books, err := h.service.GetAllBooks(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(res)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.UpdateBook()
	_, _ = fmt.Fprintf(w, res)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	res, _ := h.service.DeleteBook()
	_, _ = fmt.Fprintf(w, res)
}
