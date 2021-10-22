package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPokemonController struct {
	mock.Mock
}

func (pc MockPokemonController) ConsumeAPI(limit, offset int64) ([]byte, error) {
	args := pc.Called(limit, offset)
	return args.Get(0).([]byte), args.Error(1)
}

func (pc MockPokemonController) StoreData(content []byte) (entities.APIResponse, error) {
	args := pc.Called(content)
	return args.Get(0).(entities.APIResponse), args.Error(1)
}

func (pc MockPokemonController) GetPokemonsData() (map[string]entities.Pokemon, []entities.Pokemon, error) {
	args := pc.Called()
	return args.Get(0).(map[string]entities.Pokemon), args.Get(1).([]entities.Pokemon), args.Error(2)
}

func (pc MockPokemonController) GetPokemonsDataConcurrently(flag string, items, itemsWorker int64) (map[string]entities.Pokemon, error) {
	args := pc.Called(flag, items, itemsWorker)
	return args.Get(0).(map[string]entities.Pokemon), args.Error(1)
}

var expectedMap = map[string]entities.Pokemon{
	"201": {
		Id: "201",
		Name: "unown",
	},
	"202": {
		Id: "202",
		Name: "wobbuffet",
	},
	"203": {
		Id: "203",
		Name: "girafarig",
	},
}
var expectedPokemon = "{\"id\":\"201\",\"name\":\"unown\"}\n"

func TestPokemonController_GetPokemons(t *testing.T) {
	testCases := []struct {
		name string
		hasError bool
		expectedMapResponse map[string]entities.Pokemon
		expectedSliceResponse []entities.Pokemon
		expectedError error
	} {
		{
			"should return bad response on error",
			true,
			map[string]entities.Pokemon{},
			[]entities.Pokemon{},
			errors.New("generic error"),
		},
		{
			"should return success response",
			false,
			map[string]entities.Pokemon{},
			[]entities.Pokemon{},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/pokemons", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
		
			mock := MockPokemonController{}
			mock.On("GetPokemonsData").Return(tc.expectedMapResponse, tc.expectedSliceResponse, tc.expectedError)
			
			controller := NewPokemonController(mock)
			
			if tc.hasError {
				controller.GetPokemons(c)
				assert.Equal(t, http.StatusBadRequest, rec.Code)
			} else {
				if assert.NoError(t, controller.GetPokemons(c)) {
					assert.Equal(t, http.StatusOK, rec.Code)
				}
			}
		})
	}
}

func TestPokemonController_GetPokemonById(t *testing.T) {
	testCases := []struct {
		name string
		hasError bool
		idParam string
		expectedMapResponse map[string]entities.Pokemon
		expectedSliceResponse []entities.Pokemon
		expectedError error
		expectedErrMsg string
		expectedBadHTTP int
		expectedPokemon string
	} {
		{
			"should return bad request on id not provided",
			true,
			"",
			map[string]entities.Pokemon{},
			[]entities.Pokemon{},
			nil,
			"\"id was not provided\"\n",
			400,
			"",
		},
		{
			"should return bad request on pokemon not found",
			true,
			"199",
			expectedMap,
			[]entities.Pokemon{},
			nil,
			"\"pokemon not found\"\n",
			404,
			"",
		},
		{
			"should return success response on pokemon found",
			false,
			"201",
			expectedMap,
			[]entities.Pokemon{},
			nil,
			"",
			0,
			expectedPokemon,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/api/pokemons", nil)
			rec := httptest.NewRecorder()
		
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.idParam)

			mock := MockPokemonController{}
			mock.On("GetPokemonsData").Return(tc.expectedMapResponse, tc.expectedSliceResponse, tc.expectedError)

			controller := NewPokemonController(mock)

			controller.GetPokemonById(c)

			if tc.hasError {
				assert.Equal(t, tc.expectedBadHTTP, rec.Code)
				assert.Equal(t, tc.expectedErrMsg, rec.Body.String())
			} else {
				assert.Equal(t, http.StatusOK, rec.Code)
				assert.Equal(t, rec.Body.String(), tc.expectedPokemon)
			}
		})
	}
}
