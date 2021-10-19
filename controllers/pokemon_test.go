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