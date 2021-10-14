package usecases

import (
	"errors"
	"os"

	"github.com/gasandov/academy-go-q32021/entities"
)

type PokemonService struct {
	repo csvIO
}

type csvIO interface {
	ReadFile(name string) ([][]string, error)
	WriteFile(file *os.File, data []byte) (entities.API, error)
	CreateFile(name string) (*os.File, error)
	FileExists(name string) bool
}

// Receives fileName and reads from file (if exists)
// returns a map and slice of pokemons
func (ps *PokemonService) Get(fileName string) (map[string]entities.Pokemon, []entities.Pokemon, error) {
	fileExists := ps.repo.FileExists(fileName)

	if !fileExists {
		return nil, nil, errors.New("source not found")
	}

	content, err := ps.repo.ReadFile(fileName)

	if err != nil {
		return nil, nil, errors.New("source could not be readed")
	}

	pkMap, pkSlice := NewCollectionService().BuildCollections(content)

	return pkMap, pkSlice, nil
}

// Receives fileName and content []byte, creates file and writes the content on it
// returns api response
func (ps *PokemonService) Save(fileName string, content []byte) (entities.API, error) {
	file, err := ps.repo.CreateFile(fileName)

	if err != nil {
		return entities.API{}, err
	}

	response, err := ps.repo.WriteFile(file, content)

	if err != nil {
		return entities.API{}, err
	}

	return response, nil
}

func NewPokemonService(repo csvIO) *PokemonService {
	return &PokemonService{repo}
}
