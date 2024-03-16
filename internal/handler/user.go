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

type UserHandler struct {
	service service.User
}

func NewUserHandler(service service.User) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with the input payload
// @TAGS user
// @ID register-user
// @Accept  json
// @Produce  json
// @Param user body types.AuthUser true "User object that needs to be registered"
// @Success 201 {object} types.AuthUser
// @Router /api/v1/users/register [post]
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user types.AuthUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateRegisterUser(v, &user)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	newUser, err := h.service.RegisterUser(r.Context(), &user)
	if err != nil {
		if errors.Is(err, service.ErrEntityExists) {
			badRequestResponse(w, r, fmt.Errorf("user '%s' already exists", user.Email))
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusCreated, envelope{"user": newUser}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}

// LoginUser godoc
// @Summary Log in a user
// @Description Log in a user with the input payload
// @TAGS user
// @ID login-user
// @Accept  json
// @Produce  json
// @Param user body types.AuthUser true "User object that needs to log in"
// @Success 200 {object} map[string]string "{ \"access_token\": \"token\" }"
// @Router /api/v1/users/login [post]
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user types.AuthUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		badRequestResponse(w, r, errors.New("can't decode request"))
		return
	}

	v := validator.New()
	types.ValidateLoginUser(v, &user)
	if !v.IsValid() {
		notValidResponse(w, r, v.Errors)
		return
	}

	token, err := h.service.LoginUser(r.Context(), &user)
	if err != nil {
		if errors.Is(err, service.ErrNotFound) || errors.Is(err, service.ErrCredentialsMismatch) {
			invalidCredentialsResponse(w, r)
			return
		}
		serverErrorResponse(w, r, err)
		return
	}

	err = writeJSON(w, http.StatusOK, envelope{"access_token": token}, nil)
	if err != nil {
		log.Error(err.Error())
	}
}
