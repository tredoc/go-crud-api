package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	mockservice "github.com/tredoc/go-crud-api/mocks/service"
	"github.com/tredoc/go-crud-api/pkg/types"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type bookHandlerSuite struct {
	suite.Suite
	usecase       *mockservice.Book
	handler       *BookHandler
	testingServer *httptest.Server
}

func (s *bookHandlerSuite) SetupSuite() {
	usecase := new(mockservice.Book)
	handler := NewBookHandler(usecase)

	router := httprouter.New()
	router.POST("/api/v1/books", handler.CreateBook)
	router.GET("/api/v1/books", handler.GetAllBooks)
	router.GET("/api/v1/books/:id", handler.GetBookByID)
	router.PATCH("/api/v1/books/:id", handler.UpdateBook)
	router.DELETE("/api/v1/books/:id", handler.DeleteBook)

	testingServer := httptest.NewServer(router)

	s.testingServer = testingServer
	s.usecase = usecase
	s.handler = handler
}

func (s *bookHandlerSuite) TearDownSuite() {
	s.usecase.AssertExpectations(s.T())
	defer s.testingServer.Close()
}

func (s *bookHandlerSuite) TestCreateBook_Positive() {
	parsedTime, _ := time.Parse(time.DateOnly, "2006-01-01")
	customDate := types.CustomDate{Time: parsedTime}

	book := types.Book{
		Title:       "Go mechanics",
		PublishDate: customDate,
		ISBN:        "11111100-09000",
		Pages:       499,
		Authors:     []int64{1, 2},
		Genres:      []int64{1, 2},
	}

	newBook := types.BookWithDetails{
		ID:          1,
		Title:       "Go mechanics",
		PublishDate: customDate,
		CreatedAt:   time.Now(),
		ISBN:        "11111100-09000",
		Pages:       499,
		Authors:     []*types.Author{{ID: 1, FirstName: "firstName", MiddleName: "middleName", LastName: "lastName"}},
		Genres:      []*types.Genre{{ID: 1, Name: "Programming"}},
	}

	s.usecase.On("CreateBook", mock.AnythingOfType("*context.cancelCtx"), &book).Return(&newBook, nil)

	requestBody, err := json.Marshal(&book)
	s.NoError(err, "can`t marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/api/v1/books", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"book": &newBook,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusCreated, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *bookHandlerSuite) TestGetBookByID_Positive() {
	id := int64(1)
	parsedTime, _ := time.Parse(time.DateOnly, "2006-01-01")
	book := types.BookWithDetails{
		ID:          id,
		Title:       "Go mechanics",
		PublishDate: types.CustomDate{Time: parsedTime},
		CreatedAt:   time.Now(),
		ISBN:        "11111100-09000",
		Pages:       499,
		Authors:     []*types.Author{{ID: 1, FirstName: "firstName", MiddleName: "middleName", LastName: "lastName"}},
		Genres:      []*types.Genre{{ID: 1, Name: "Programming"}},
	}

	s.usecase.On("GetBookByID", mock.AnythingOfType("*context.cancelCtx"), book.ID).Return(&book, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/books/%d", s.testingServer.URL, book.ID))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"book": &book,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *bookHandlerSuite) TestGetAllBooks_Positive() {
	parsedTime, _ := time.Parse(time.DateOnly, "2006-01-01")
	customDate := types.CustomDate{Time: parsedTime}
	books := []*types.Book{
		{
			ID:          1,
			Title:       "First book",
			PublishDate: customDate,
			CreatedAt:   time.Now(),
			ISBN:        "000000000000-00",
			Pages:       399,
			Authors:     []int64{1, 2},
			Genres:      []int64{1, 2},
		},
		{
			ID:          2,
			Title:       "Second book",
			PublishDate: customDate,
			CreatedAt:   time.Now(),
			ISBN:        "11111111111111-11",
			Pages:       499,
			Authors:     []int64{3, 4},
			Genres:      []int64{3, 4},
		},
	}

	s.usecase.On("GetAllBooks", mock.AnythingOfType("*context.cancelCtx")).Return(books, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/books", s.testingServer.URL))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"books": &books,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *bookHandlerSuite) TestUpdateBook_Positive() {
	id := int64(1)
	newTitle := "Update Title"
	newAuthors := []int64{1}
	upd := types.UpdateBook{
		Title:   &newTitle,
		Authors: []int64{1},
	}

	parsedTime, _ := time.Parse(time.DateOnly, "2006-01-01")
	book := types.Book{
		ID:          id,
		Title:       newTitle,
		PublishDate: types.CustomDate{Time: parsedTime},
		ISBN:        "11111100-09000",
		Pages:       499,
		Authors:     newAuthors,
		Genres:      []int64{1, 2},
	}

	s.usecase.On("UpdateBook", mock.AnythingOfType("*context.cancelCtx"), id, &upd).Return(&book, nil)

	requestBody, err := json.Marshal(&upd)
	s.NoError(err, "can`t marshal struct to json")

	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/books/%d", s.testingServer.URL, book.ID), bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when preparing patch request")

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"book": &book,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *bookHandlerSuite) TestDeleteBook_Positive() {
	id := int64(1)
	s.usecase.On("DeleteBook", mock.AnythingOfType("*context.cancelCtx"), id).Return(nil)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/books/%d", s.testingServer.URL, id), nil)
	s.NoError(err, "no error when preparing patch request")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")

	s.Equal(http.StatusNoContent, response.StatusCode)
}

func TestBookHandler(t *testing.T) {
	suite.Run(t, new(bookHandlerSuite))
}
