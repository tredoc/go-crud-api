package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/internal/validator"
	"github.com/tredoc/go-crud-api/pkg/types"
	"net/http"
	"strconv"
)

type AuthorHandler struct {
	service service.Author
}

func NewAuthorHandler(service service.Author) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var author types.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	types.ValidateAuthor(v, &author)
	if !v.IsValid() {
		resp, err := json.Marshal(v.Errors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(resp)
		return
	}

	ctx := r.Context()
	res, err := h.service.CreateAuthor(ctx, &author)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("already exists"))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, string(resp))
}

func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	author, err := h.service.GetAuthorByID(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := json.Marshal(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (h *AuthorHandler) GetAuthorByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	firstName := ps.ByName(firstNameParam)
	lastName := ps.ByName(lastNameParam)

	if firstName == "" || lastName == "" {
		http.Error(w, "missing first_name or last_name parameter", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	author, err := h.service.GetAuthorByName(ctx, firstName, lastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var authors []*types.Author
	ctx := r.Context()
	authors, err := h.service.GetAllAuthors(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if authors == nil {
		_, _ = fmt.Fprint(w, "[]")
		return
	}

	resp, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(resp)
}

func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var author types.UpdateAuthor
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	types.ValidateUpdateAuthor(v, &author)
	if !v.IsValid() {
		resp, err := json.Marshal(v.Errors)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(resp)
		return
	}

	ctx := r.Context()
	updatedAuthor, err := h.service.UpdateAuthor(ctx, id, &author)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "author not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(updatedAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(resp)
}

func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	err = h.service.DeleteAuthor(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "no author with such id", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
