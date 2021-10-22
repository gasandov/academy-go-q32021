package usecases

import (
	"errors"
	"io/ioutil"
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

var expectedAPIResponse = entities.APIResponse{
	Count: 3,
	Next: "next",
	Previous: "prev",
	Results: []entities.PokemonAPI{
		{
			Name: "one",
			Url: "url/one",
		},
		{
			Name: "two",
			Url: "url/two",
		},
		{
			Name: "three",
			Url: "url/three",
		},
	},
}

func TestPokemonService_StoreData(t *testing.T) {
	tempFile, err := ioutil.TempFile("", constants.FileName)
	if err != nil {
		return
	}

	testCases := []struct {
		name string
		fileExists bool
		tempFile *os.File
		createFileErr error
		writeFileErr error
		content []byte
		hasError bool
		expectedErrMsg string
		expectedLength int
		expectedResponse entities.APIResponse
	} {
		{
			"should return error on CreateFile fails",
			false,
			tempFile,
			errors.New("could not create file"),
			nil,
			[]byte{},
			true,
			"could not create file",
			0,
			entities.APIResponse{},
		},
		{
			"should return error on WriteFile fails",
			false,
			tempFile,
			nil,
			errors.New("coult not write in file"),
			[]byte{},
			true,
			"coult not write in file",
			0,
			entities.APIResponse{},
		},
		{
			"should return success response",
			false,
			tempFile,
			nil,
			nil,
			[]byte{},
			false,
			"",
			3,
			expectedAPIResponse,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := MockFileManagerRepo{}
			mock.On("FileExists", constants.FileName).Return(tc.fileExists)
			mock.On("CreateFile", constants.FileName).Return(tc.tempFile, tc.createFileErr)
			mock.On("WriteFile", tc.tempFile, tc.content).Return(tc.expectedResponse, tc.writeFileErr)

			service := NewPokemonService(mock)

			response, err := service.StoreData(tc.content)

			if tc.hasError {
				assert.EqualError(t, err, tc.expectedErrMsg)
			}

			assert.Len(t, response.Results, tc.expectedLength)
			assert.Equal(t, response, tc.expectedResponse)
		})
	}
}
