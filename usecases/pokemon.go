package usecases

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gasandov/academy-go-q32021/constants"
	"github.com/gasandov/academy-go-q32021/entities"
	"github.com/gasandov/academy-go-q32021/utils"
)

type PokemonService struct {
	repo fileManager
}

type fileManager interface {
	FileExists(name string) bool
	CreateFile(name string) (*os.File, error)
	ReadFile(name string) ([][]string, error)
	OpenFile(name, flag string) (*os.File, error)
	WriteFile(file *os.File, data []byte) (entities.APIResponse, error)
	ReadFileConcurrently(name, flag string, items, itemsWorker int64) ([][]string, error)
}

// receives limit and offset
// calls api endpoint with query params and returns response
// (body []byte)
func (ps *PokemonService) ConsumeAPI(limit, offset int64) ([]byte, error) {
	endpoint := fmt.Sprintf("%s?limit=%d&offset=%d", constants.APIUrl, limit, offset)

	data, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	response, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// receives (content []byte)
// creates file and writes content on it
// if file exists, new content is append it
// returns initial API response
func (ps *PokemonService) StoreData(content []byte) (entities.APIResponse, error) {
	if ps.repo.FileExists(constants.FileName) {
		file, err := ps.repo.OpenFile(constants.FileName, "append")
		if err != nil {
			return entities.APIResponse{}, err
		}

		response, err := ps.repo.WriteFile(file, content)
		if err != nil {
			return entities.APIResponse{}, err
		}

		return response, err
	}

	file, err := ps.repo.CreateFile(constants.FileName)
	if err != nil {
		return entities.APIResponse{}, err
	}

	response, err := ps.repo.WriteFile(file, content)
	if err != nil {
		return entities.APIResponse{}, err
	}

	return response, err
}

// if file exists returns map and slice of pokemons
func (ps *PokemonService) GetPokemonsData() (map[string]entities.Pokemon, []entities.Pokemon, error) {
	exists := ps.repo.FileExists(constants.FileName)
	if !exists {
		return nil, nil, errors.New("source not found")
	}

	content, err := ps.repo.ReadFile(constants.FileName)
	if err != nil {
		return nil, nil, err
	}

	pkMap, pkSlice := utils.BuildCollections(content)

	return pkMap, pkSlice, nil
}

func (ps* PokemonService) GetPokemonsDataConcurrently(flag string, items, itemsWorker int64) (map[string]entities.Pokemon, error) {
	exists := ps.repo.FileExists(constants.FileName)
	if !exists {
		return nil, errors.New("source not found")
	}

	content, err := ps.repo.ReadFileConcurrently(constants.FileName, flag, items, itemsWorker)
	if err != nil {
		return nil, err
	}

	pkMap, _ := utils.BuildCollections(content)

	return pkMap, nil
}

func NewPokemonService(repo fileManager) *PokemonService {
	return &PokemonService{repo}
}
