package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-crud-api/internal/cache"
	"github.com/tredoc/go-crud-api/internal/repository"
	"github.com/tredoc/go-crud-api/pkg/types"
)

type BookService struct {
	repo       repository.Book
	authorRepo repository.Author
	genreRepo  repository.Genre
	cache      cache.RCache
}

func NewBookService(bookRepo repository.Book, authorRepo repository.Author, genreRepo repository.Genre, cache cache.RCache) *BookService {
	return &BookService{
		repo:       bookRepo,
		authorRepo: authorRepo,
		genreRepo:  genreRepo,
		cache:      cache,
	}
}

func (s *BookService) CreateBook(ctx context.Context, book *types.Book) (*types.BookWithDetails, error) {
	id, createdAt, err := s.repo.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	authors, err := s.authorRepo.GetAuthorsByIDs(ctx, book.Authors)
	if err != nil {
		return nil, err
	}

	genres, err := s.genreRepo.GetGenresByIDs(ctx, book.Genres)
	if err != nil {
		return nil, err
	}

	newBook := types.BookWithDetails{
		ID:          id,
		Title:       book.Title,
		PublishDate: book.PublishDate,
		CreatedAt:   createdAt,
		ISBN:        book.ISBN,
		Pages:       book.Pages,
		Authors:     authors,
		Genres:      genres,
	}
	go s.cache.Invalidate("books")
	return &newBook, nil
}

func (s *BookService) GetBookByID(ctx context.Context, id int64) (*types.BookWithDetails, error) {
	key := fmt.Sprintf("book:%d", id)
	var bookCache types.BookWithDetails
	err := getFromCache(s.cache.Get, key, &bookCache)
	if err == nil {
		return &bookCache, nil
	}

	book, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	authors, err := s.authorRepo.GetAuthorsByIDs(ctx, book.Authors)
	if err != nil {
		return nil, err
	}

	genres, err := s.genreRepo.GetGenresByIDs(ctx, book.Genres)
	if err != nil {
		return nil, err
	}

	bookWithDetails := types.BookWithDetails{
		ID:          id,
		Title:       book.Title,
		PublishDate: book.PublishDate,
		CreatedAt:   book.CreatedAt,
		ISBN:        book.ISBN,
		Pages:       book.Pages,
		Authors:     authors,
		Genres:      genres,
	}

	go setToCache(s.cache.Set, key, bookWithDetails, cache.EXPIRATION)
	return &bookWithDetails, nil
}

func (s *BookService) GetAllBooks(ctx context.Context) ([]*types.Book, error) {
	key := "books"
	var booksCache []*types.Book
	err := getFromCache(s.cache.Get, key, &booksCache)
	if err == nil {
		return booksCache, nil
	}

	books, err := s.repo.GetAllBooks(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return books, nil
		}

		return nil, err
	}

	go setToCache(s.cache.Set, key, books, cache.EXPIRATION)
	return books, nil
}

func (s *BookService) UpdateBook(ctx context.Context, id int64, book *types.UpdateBook) (*types.Book, error) {
	bookUPD, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	if book.Title != nil {
		bookUPD.Title = *book.Title
	}

	if book.PublishDate != nil {
		bookUPD.PublishDate = *book.PublishDate
	}

	if book.ISBN != nil {
		bookUPD.ISBN = *book.ISBN
	}

	if book.Pages != nil {
		bookUPD.Pages = *book.Pages
	}

	if book.Authors != nil {
		bookUPD.Authors = book.Authors
	}

	if book.Genres != nil {
		bookUPD.Genres = book.Genres
	}

	err = s.repo.UpdateBook(ctx, id, bookUPD)
	if err != nil {
		return nil, err
	}

	go s.cache.Invalidate("books")
	go s.cache.Invalidate(fmt.Sprintf("book:%d", id))
	return bookUPD, nil
}

func (s *BookService) DeleteBook(ctx context.Context, id int64) error {
	go s.cache.Invalidate("books")
	go s.cache.Invalidate(fmt.Sprintf("book:%d", id))
	return s.repo.DeleteBook(ctx, id)
}
