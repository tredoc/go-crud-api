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

type GenreHandler struct {
	service service.Genre
}

func NewGenreHandler(service service.Genre) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

// CreateGenre godoc
// @Summary Create a new genre
// @Description Create a new genre with the input payload
// @Tags genres
// @ID create-genre
// @Accept  json
// @Produce  json
// @Param genre body types.Genre true "Genre object that needs to be added"
// @Security Bearer
// @Success 201 {object} types.Genre
// @Router /api/v1/genres [post]
func (h *GenreHandler) CreateGenre(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var genre types.Genre
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateGenre(v, &genre)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	newGenre, err := h.service.CreateGenre(r.Context(), &genre)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			badRequestResponse(w, r, fmt.Errorf("genre '%s' already exists", genre.Name))
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, envelope{"genre": newGenre}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// GetGenreByID godoc
// @Summary Get details of a genre
// @Description Get details of a genre by ID
// @Tags genres
// @ID get-genre-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Genre ID"
// @Success 200 {object} types.Genre
// @Router /api/v1/genres/{id} [get]
func (h *GenreHandler) GetGenreByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
	}

	genre, err := h.service.GetGenreByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"genre": genre}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// GetAllGenres godoc
// @Summary Get all genres
// @Description Get a list of all genres
// @Tags genres
// @ID get-all-genres
// @Accept  json
// @Produce  json
// @Success 200 {array} []types.Genre
// @Router /api/v1/genres [get]
func (h *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	genres, err := h.service.GetAllGenres(r.Context())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"genres": genres}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// UpdateGenre godoc
// @Summary Update a genre
// @Description Update a genre with a specific ID
// @Tags genres
// @ID update-genre
// @Accept  json
// @Produce  json
// @Param id path int true "Genre ID"
// @Param genre body types.Genre true "Genre object that needs to be updated"
// @Security Bearer
// @Success 200 {object} types.Genre
// @Router /api/v1/genres/{id} [put]
func (h *GenreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
	}

	var genre types.Genre
	err = json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateGenre(v, &genre)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	err = h.service.UpdateGenre(r.Context(), id, &genre)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	genre.ID = id
	err = writeJSON(w, http.StatusOK, envelope{"genre": genre}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// DeleteGenre godoc
// @Summary Delete a genre
// @Description Delete a genre with a specific ID
// @Tags genres
// @ID delete-genre
// @Accept  json
// @Produce  json
// @Param id path int true "Genre ID"
// @Security Bearer
// @Success 204 "No Content"
// @Router /api/v1/genres/{id} [delete]
func (h *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	err = h.service.DeleteGenre(r.Context(), id)
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
