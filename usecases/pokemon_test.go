package usecases

import (
	"errors"
	"os"
	"testing"

	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCSVRepo struct {
	mock.Mock
}

func (cv MockCSVRepo) FileExists(name string) bool {
	args := cv.Called(name)
	return args.Bool(0)
}

func (cv MockCSVRepo) ReadFile(name string) ([][]string, error) {
	args := cv.Called(name)
	return args.Get(0).([][]string), args.Error(1)
}

func (cv MockCSVRepo) CreateFile(name string) (*os.File, error) {
	args := cv.Called(name)
	return args.Get(0).(*os.File), args.Error(1)
}

func (cv MockCSVRepo) WriteFile(file *os.File, data []byte) (entities.API, error) {
	args := cv.Called(file, data)
	return args.Get(0).(entities.API), args.Error(1)
}

var fileName string = "pokemons_list.csv"

func TestPokemonService_Get(t *testing.T) {
	t.Run("file does not exists", func(t *testing.T) {
		mock := MockCSVRepo{}
		mock.On("FileExists", fileName).Return(false)

		service := NewPokemonService(mock)

		_, _, err := service.Get(fileName)

		assert.EqualError(t, err, "source not found")
	})
	
	t.Run("file could not be readed", func(t *testing.T) {
		mock := MockCSVRepo{}
		mock.On("FileExists", fileName).Return(true)
		mock.On("ReadFile", fileName).Return(pokemons, errors.New("source could not be readed"))

		service := NewPokemonService(mock)

		_, _, err := service.Get(fileName)

		assert.EqualError(t, err, "source could not be readed")
	})
}