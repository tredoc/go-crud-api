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

type genreHandlerSuite struct {
	suite.Suite
	usecase       *mockservice.Genre
	handler       *GenreHandler
	testingServer *httptest.Server
}

func (s *genreHandlerSuite) SetupSuite() {
	usecase := new(mockservice.Genre)
	handler := NewGenreHandler(usecase)

	router := httprouter.New()
	router.POST("/api/v1/genres", handler.CreateGenre)
	router.GET("/api/v1/genres", handler.GetAllGenres)
	router.GET("/api/v1/genres/:id", handler.GetGenreByID)
	router.PATCH("/api/v1/genres/:id", handler.UpdateGenre)
	router.DELETE("/api/v1/genres/:id", handler.DeleteGenre)

	testingServer := httptest.NewServer(router)

	s.testingServer = testingServer
	s.usecase = usecase
	s.handler = handler
}

func (s *genreHandlerSuite) TearDownSuite() {
	s.usecase.AssertExpectations(s.T())
	defer s.testingServer.Close()
}

func (s *genreHandlerSuite) TestCreateGenre_Positive() {
	genre := types.Genre{
		Name: "any",
	}

	newGenre := types.Genre{ID: 1, Name: "any"}

	s.usecase.On("CreateGenre", mock.AnythingOfType("*context.cancelCtx"), &genre).Return(&newGenre, nil)

	requestBody, err := json.Marshal(&genre)
	s.NoError(err, "can`t marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/api/v1/genres", s.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"genre": newGenre,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusCreated, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *genreHandlerSuite) TestGetGenreByID_Positive() {
	id := int64(1)
	genre := types.Genre{ID: id, Name: "any"}

	s.usecase.On("GetGenreByID", mock.AnythingOfType("*context.cancelCtx"), genre.ID).Return(&genre, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/genres/%d", s.testingServer.URL, id))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"genre": genre,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *genreHandlerSuite) TestGetAllGenres_Positive() {
	genres := []*types.Genre{{ID: 1, Name: "medicine"}, {ID: 2, Name: "programming"}}

	s.usecase.On("GetAllGenres", mock.AnythingOfType("*context.cancelCtx")).Return(genres, nil)

	response, err := http.Get(fmt.Sprintf("%s/api/v1/genres", s.testingServer.URL))
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"genres": genres,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *genreHandlerSuite) TestUpdateGenre_Positive() {
	id := int64(1)
	genre := types.Genre{ID: id, Name: "updated"}

	s.usecase.On("UpdateGenre", mock.AnythingOfType("*context.cancelCtx"), id, &genre).Return(nil)

	requestBody, err := json.Marshal(&genre)
	s.NoError(err, "can`t marshal struct to json")

	request, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/genres/%d", s.testingServer.URL, id), bytes.NewBuffer(requestBody))
	s.NoError(err, "no error when preparing patch request")

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	result, err := io.ReadAll(response.Body)
	s.NoError(err, "can`t get string from response")

	expected, err := json.Marshal(map[string]any{
		"genre": genre,
	})
	s.NoError(err, "can`t convert expected map to json")

	s.Equal(http.StatusOK, response.StatusCode)
	s.Equal(string(result), string(expected))
}

func (s *genreHandlerSuite) TestDeleteGenre_Positive() {
	id := int64(1)
	s.usecase.On("DeleteGenre", mock.AnythingOfType("*context.cancelCtx"), id).Return(nil)

	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/v1/genres/%d", s.testingServer.URL, id), nil)
	s.NoError(err, "no error when preparing patch request")

	client := &http.Client{}
	response, err := client.Do(request)
	s.NoError(err, "no error when calling the endpoint")

	s.Equal(http.StatusNoContent, response.StatusCode)
}

func TestGenreHandler(t *testing.T) {
	suite.Run(t, new(genreHandlerSuite))
}
