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

type AuthorHandler struct {
	service service.Author
}

func NewAuthorHandler(service service.Author) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

// CreateAuthor godoc
// @Summary Create a new author
// @Description Create a new author with the input payload
// @Tags authors
// @ID create-author
// @Accept  json
// @Produce  json
// @Param author body types.Author true "Author object that needs to be added"
// @Security Bearer
// @Success 201 {object} types.Author
// @Router /api/v1/authors [post]
func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var author types.Author
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateAuthor(v, &author)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	newAuthor, err := h.service.CreateAuthor(r.Context(), &author)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			badRequestResponse(w, r, fmt.Errorf("author '%s %s %s' already exists", author.FirstName, author.MiddleName, author.LastName))
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, envelope{"author": newAuthor}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// GetAuthorByID godoc
// @Summary Get details of an author
// @Description Get details of an author by ID
// @Tags authors
// @ID get-author-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 {object} types.Author
// @Router /api/v1/authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	author, err := h.service.GetAuthorByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"author": author}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// GetAllAuthors godoc
// @Summary Get all authors
// @Description Get a list of all authors
// @Tags authors
// @ID get-all-authors
// @Accept  json
// @Produce  json
// @Success 200 {array} []types.Author
// @Router /api/v1/authors [get]
func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authors, err := h.service.GetAllAuthors(r.Context())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"authors": authors}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// UpdateAuthor godoc
// @Summary Update an author
// @Description Update an author with a specific ID
// @Tags authors
// @ID update-author
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Param author body types.UpdateAuthor true "Author object that needs to be updated"
// @Security Bearer
// @Success 200 {object} types.Author
// @Router /api/v1/authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	var author types.UpdateAuthor
	err = json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateUpdateAuthor(v, &author)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	updatedAuthor, err := h.service.UpdateAuthor(r.Context(), id, &author)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			notFoundResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"author": updatedAuthor}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// DeleteAuthor godoc
// @Summary Delete an author
// @Description Delete an author with a specific ID
// @Tags authors
// @ID delete-author
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Security Bearer
// @Success 204 "No Content"
// @Router /api/v1/authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := getIdParam(ps)
	if err != nil {
		badRequestResponse(w, r, err)
		return
	}

	err = h.service.DeleteAuthor(r.Context(), id)
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
