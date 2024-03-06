package handler

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tredoc/go-crud-api/internal/service"
	"net/http"
)

type Book interface {
	CreateBook(http.ResponseWriter, *http.Request, httprouter.Params)
	GetBookByID(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllBooks(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateBook(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteBook(http.ResponseWriter, *http.Request, httprouter.Params)
}

type Genre interface {
	CreateGenre(http.ResponseWriter, *http.Request, httprouter.Params)
	GetGenreByID(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllGenres(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateGenre(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteGenre(http.ResponseWriter, *http.Request, httprouter.Params)
}

type Author interface {
	CreateAuthor(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAuthorByID(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllAuthors(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateAuthor(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteAuthor(http.ResponseWriter, *http.Request, httprouter.Params)
}

type User interface {
	RegisterUser(http.ResponseWriter, *http.Request, httprouter.Params)
	LoginUser(http.ResponseWriter, *http.Request, httprouter.Params)
}

type Middlewares interface {
	authMiddleware(httprouter.Handle) httprouter.Handle
}

type Handler struct {
	book   Book
	genre  Genre
	author Author
	user   User
	mw     Middlewares
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		book:   NewBookHandler(services.Book),
		genre:  NewGenreHandler(services.Genre),
		author: NewAuthorHandler(services.Author),
		user:   NewUserHandler(services.User),
		mw:     NewMiddleware(services.User),
	}
}

func (h *Handler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(notAllowedResponse)

	router.POST("/api/v1/books", h.mw.authMiddleware(h.book.CreateBook))
	router.GET("/api/v1/books", h.book.GetAllBooks)
	router.GET("/api/v1/books/:id", h.book.GetBookByID)
	router.PATCH("/api/v1/books/:id", h.book.UpdateBook)
	router.DELETE("/api/v1/books/:id", h.book.DeleteBook)

	router.POST("/api/v1/genres", h.genre.CreateGenre)
	router.GET("/api/v1/genres", h.genre.GetAllGenres)
	router.GET("/api/v1/genres/:id", h.genre.GetGenreByID)
	router.PATCH("/api/v1/genres/:id", h.genre.UpdateGenre)
	router.DELETE("/api/v1/genres/:id", h.genre.DeleteGenre)

	router.POST("/api/v1/authors", h.author.CreateAuthor)
	router.GET("/api/v1/authors", h.author.GetAllAuthors)
	router.GET("/api/v1/authors/:id", h.author.GetAuthorByID)
	router.PATCH("/api/v1/authors/:id", h.author.UpdateAuthor)
	router.DELETE("/api/v1/authors/:id", h.author.DeleteAuthor)

	router.POST("/auth/register", h.user.RegisterUser)
	router.POST("/auth/login", h.user.LoginUser)

	return router
}
