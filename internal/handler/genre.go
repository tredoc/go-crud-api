package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/pkg/types"
	"net/http"
	"strconv"
)

type GenreHandler struct {
	service service.Genre
}

func NewGenreHandler(service service.Genre) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (h *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var genre types.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	res, err := h.service.CreateGenre(ctx, &genre)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			http.Error(w, "genre already exists", http.StatusBadRequest)
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

func (h *GenreHandler) GetGenreByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	genre, err := h.service.GetGenreByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, string(resp))
}

func (h *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var genres []*types.Genre
	ctx := r.Context()
	genres, err := h.service.GetAllGenres(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if genres == nil {
		_, _ = fmt.Fprint(w, "[]")
		return
	}

	resp, err := json.Marshal(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, string(resp))
}

func (h *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	var genre types.Genre
	err = json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.service.UpdateGenre(ctx, id, &genre)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "no genre with such id", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, _ = fmt.Fprint(w, string(resp))
}

func (h *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	err = h.service.DeleteGenre(ctx, id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			http.Error(w, "no genre with such id", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
