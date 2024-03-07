package handler

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"github.com/pascaldekloe/jwt"
	"github.com/tredoc/go-crud-api/internal/service"
	"github.com/tredoc/go-crud-api/pkg/types"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Middleware struct {
	service service.User
}

func NewMiddleware(service service.User) *Middleware {
	return &Middleware{service: service}
}

func (m *Middleware) authMW(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Vary", "Authorization")
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" {
			r = contextSetUser(r, types.AnonymousUser)
			next(w, r, ps)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			invalidAuthenticationTokenResponse(w, r)
			return
		}
		token := headerParts[1]

		claims, err := jwt.HMACCheck([]byte(token), []byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			invalidAuthenticationTokenResponse(w, r)
			return
		}

		if !claims.Valid(time.Now()) {
			invalidAuthenticationTokenResponse(w, r)
			return
		}

		if claims.Issuer != "go-crud-api" {
			invalidAuthenticationTokenResponse(w, r)
			return
		}

		if !claims.AcceptAudience("go-crud-api") {
			invalidAuthenticationTokenResponse(w, r)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject, 10, 64)
		if err != nil {
			serverErrorResponse(w, r, err)
			return
		}

		user, err := m.service.GetUserByID(r.Context(), userID)
		if err != nil {
			switch {
			case errors.Is(err, service.ErrNotFound):
				invalidAuthenticationTokenResponse(w, r)
			default:
				serverErrorResponse(w, r, err)
			}
			return
		}

		r = contextSetUser(r, user)
		next(w, r, ps)
	}
}

func (m *Middleware) adminOnlyMW(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		val := r.Context().Value(types.UserContextKey)
		if val == nil {
			insufficientPermissionsResponse(w, r)
			return
		}

		user := val.(*types.User)
		if user.Role != types.AdminRole {
			insufficientPermissionsResponse(w, r)
			return
		}
		next(w, r, ps)
	}
}
