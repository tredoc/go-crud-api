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
)

type authorHandlerSuite struct {
	suite.Suite
	usecase       *mockservice.Author
	handler       *AuthorHandler
	testingServer *httptest.Server
}

func (s *authorHandlerSuite) SetupSuite() {
	usecase := new(mockservice.Author)
	handler := NewAuthorHandler(usecase)

	router := httprouter.New()
	router.POST("/api/v1/authors", handler.CreateAuthor)
	router.GET("/api/v1/authors", handler.GetAllAuthors)
	router.GET("/api/v1/authors/:id", handler.GetAuthorByID)
	router.PATCH("/api/v1/authors/:id", handler.UpdateAuthor)
	router.DELETE("/api/v1/authors/:id", handler.DeleteAuthor)

	testingServer := httptest.NewServer(router)

	s.testingServer = testingServer
	s.usecase = usecase
	s.handler = handler
}

func (s *authorHandlerSuite) TearDownSuite() {
	s.usecase.AssertExpectations(s.T())
	defer s.testingServer.Close()
}

func (s *authorHandlerSuite) TestCreateAuthor_Positive() {
	author := types.Author{
		FirstName:  "first",
		MiddleName: "",
		LastName:   "last",
	}

	newAuthor := types.Author{ID: 1, FirstName: author.FirstName, MiddleName: author.MiddleName, LastName: author.LastName}

	s.usecase.On("CreateAuthor", mock.AnythingOfType("*context.cancelCtx"), &author).Return(&newAuthor, nil)

	requestBody, err := json.Marshal(&author)
	s.NoError(err, "can`t marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/api/v1/authors", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"author": newAuthor,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusCreated, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *authorHandlerSuite) TestGetAuthorByID_Positive() {
	id := int64(1)
	author := types.Author{ID: id, FirstName: "first", MiddleName: "", LastName: "last"}

	s.usecase.On("GetAuthorByID", mock.AnythingOfType("*context.cancelCtx"), author.ID).Return(&author, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/authors/%d", s.testingServer.URL, id))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"author": author,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *authorHandlerSuite) TestGetAllAuthors_Positive() {
	authors := []*types.Author{
		{ID: 1, FirstName: "FirstNameOne", MiddleName: "MiddleNameOne", LastName: "LastNameOne"},
		{ID: 2, FirstName: "FirstNameTwo", MiddleName: "", LastName: "LastNameTwo"},
	}

	s.usecase.On("GetAllAuthors", mock.AnythingOfType("*context.cancelCtx")).Return(authors, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/authors", s.testingServer.URL))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"authors": authors,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *authorHandlerSuite) TestUpdateAuthor_Positive() {
	id := int64(1)
	lastName := "Updated"
	upd := types.UpdateAuthor{LastName: &lastName}
	author := types.Author{ID: id, FirstName: "first", MiddleName: "", LastName: lastName}

	s.usecase.On("UpdateAuthor", mock.AnythingOfType("*context.cancelCtx"), author.ID, &upd).Return(&author, nil)

	requestBody, err := json.Marshal(&upd)
	s.NoError(err, "can`t marshal struct to json")

	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/authors/%d", s.testingServer.URL, author.ID), bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when preparing patch request")

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"author": author,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *authorHandlerSuite) TestDeleteAuthor_Positive() {
	id := int64(1)
	s.usecase.On("DeleteAuthor", mock.AnythingOfType("*context.cancelCtx"), id).Return(nil)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/authors/%d", s.testingServer.URL, id), nil)
	s.NoError(err, "no error when preparing patch request")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")

	s.Equal(http.StatusNoContent, response.StatusCode)
}

func TestAuthorHandler(t *testing.T) {
	suite.Run(t, new(authorHandlerSuite))
}
