package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/pkg/log"
	"github.com/tredoc/go-crud-api/pkg/types"
	"net/http"
	"strconv"
)

type envelope map[string]any

func writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(js)
	if err != nil {
		return err
	}

	return nil
}

func getIdParam(ps httprouter.Params) (int64, error) {
	id, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func logError(r *http.Request, err error) {
	log.Error(err.Error())
}

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	err := writeJSON(w, status, env, nil)
	if err != nil {
		logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	logError(r, err)

	message := "the server encountered a problem and could not process your request"
	errorResponse(w, r, http.StatusInternalServerError, message)
}

func badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponse(w, r, http.StatusNotFound, message)
}

func notAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func notValidResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	message := "invalid or missing authentication token"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func insufficientPermissionsResponse(w http.ResponseWriter, r *http.Request) {
	message := "insufficient permissions to perform the requested operation"
	errorResponse(w, r, http.StatusUnauthorized, message)
}

func contextSetUser(r *http.Request, user *types.User) *http.Request {
	ctx := context.WithValue(r.Context(), types.UserContextKey, user)
	return r.WithContext(ctx)
}
