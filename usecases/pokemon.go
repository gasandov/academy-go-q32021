package usecases

import (
	"errors"

	"github.com/gasandov/academy-go-q32021/entities"
)

type PokemonService struct {
	repo csvIO
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

func NewPokemonService(repo csvIO) *PokemonService {
	return &PokemonService{repo}
}
