package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/internal/validator"
	"github.com/tredoc/go-crud-api/pkg/log"
	"github.com/tredoc/go-crud-api/pkg/types"
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

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var book types.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateBook(v, &book)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	newBook, err := h.service.CreateBook(r.Context(), &book)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			badRequestResponse(w, r, fmt.Errorf("book with title '%s' already exists", book.Title))
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, envelope{"book": newBook}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
	}

	book, err := h.service.GetBookByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books, err := h.service.GetAllBooks(r.Context())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"books": books}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
	}

	var book types.UpdateBook
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		if err != nil {
			badRequestResponse(w, r, errors.New("can't decode request"))
			return
		}
	}

	v := validator.New()
	types.ValidateUpdateBook(v, &book)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	updatedBook, err := h.service.UpdateBook(r.Context(), id, &book)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"book": updatedBook}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
	}

	err = h.service.DeleteBook(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
