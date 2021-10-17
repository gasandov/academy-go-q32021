package usecases

import (
	"errors"
	"os"
	"testing"

	"github.com/gasandov/academy-go-q32021/constants"
	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileManagerRepo struct {
	mock.Mock
}

func (fm MockFileManagerRepo) FileExists(name string) bool {
	args := fm.Called(name)
	return args.Bool(0)
}

func (fm MockFileManagerRepo) CreateFile(name string) (*os.File, error) {
	args := fm.Called(name)
	return args.Get(0).(*os.File), args.Error(1)
}

func (fm MockFileManagerRepo) ReadFile(name string) ([][]string, error) {
	args := fm.Called(name)
	return args.Get(0).([][]string), args.Error(1)
}

func (fm MockFileManagerRepo) OpenFile(name, flag string) (*os.File, error) {
	args := fm.Called(name, flag)
	return args.Get(0).(*os.File), args.Error(1)
}

func (fm MockFileManagerRepo) WriteFile(file *os.File, data []byte) (entities.APIResponse, error) {
	args := fm.Called(file, data)
	return args.Get(0).(entities.APIResponse), args.Error(1)
}

func (fm MockFileManagerRepo) ReadFileConcurrently(name, flag string, items, itemsWorker int64) ([][]string, error) {
	args := fm.Called(name, flag, items, itemsWorker)
	return args.Get(0).([][]string), args.Error(1)
}

var content = [][]string{
	{
		"unown",
		"https://pokeapi.co/api/v2/pokemon/201/",
	},
	{
		"wobbuffet",
		"https://pokeapi.co/api/v2/pokemon/202/",
	},
	{
		"girafarig",
		"https://pokeapi.co/api/v2/pokemon/203/",
	},
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

func TestPokemonService_GetPokemonsData(t *testing.T) {
	testCases := []struct {
		name string
		hasError bool
		fileName string
		fileExists bool
		fileError error
		content [][]string
		errMsg string
		expectedLength int
		expectedContent map[string]entities.Pokemon
	} {
		{
			"should return error on FileExists no existing",
			true,
			constants.FileName,
			false,
			nil,
			[][]string{},
			"source not found",
			0,
			nil,
		},
		{
			"should return error on ReadFile failing",
			true,
			constants.FileName,
			true,
			errors.New("unexpected error"),
			[][]string{},
			"unexpected error",
			0,
			nil,
		},
		{
			"should return pokemon map successfully",
			false,
			constants.FileName,
			true,
			nil,
			content,
			"",
			3,
			expectedMap,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := MockFileManagerRepo{}
			mock.On("FileExists", tc.fileName).Return(tc.fileExists)
			mock.On("ReadFile", tc.fileName).Return(tc.content, tc.fileError)

			service := NewPokemonService(mock)

			pkMap, _, err := service.GetPokemonsData()

			if tc.hasError {
				assert.EqualError(t, err, tc.errMsg)
			}

			assert.Len(t, pkMap, tc.expectedLength)
			assert.Equal(t, pkMap, tc.expectedContent)
		})
	}
}

