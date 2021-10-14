package usecases

import (
	"strings"

	"github.com/gasandov/academy-go-q32021/entities"
)

type collectionService struct {}

// Receives an [[name url]] and returns a map { '1': { Id: 1, Name: 'bulbasaur'} }
// and a slice of pokemons [{ id: 1, name: 'bulbasaur' }]
func (cs *collectionService) BuildCollections(content [][]string) (map[string]entities.Pokemon, []entities.Pokemon) {
	var pkSlice []entities.Pokemon
	pkMap := make(map[string]entities.Pokemon)

	for _, line := range content {
		pokemon := parsePokemonResponse(line)

		pkMap[pokemon.Id] = pokemon
		pkSlice = append(pkSlice, pokemon)
	}

	return pkMap, pkSlice
}

// Receives [bulbasaur https://pokeapi.co/api/v2/pokemon/3/]
// and returns a Pokemon{Id: 3, Name: 'bulbasaur'}
func parsePokemonResponse(line []string) entities.Pokemon {
	name := line[0]
	url := line[1]
	splited := strings.Split(url, "/")

	id := splited[len(splited) - 2]

	return entities.Pokemon{Id: id, Name: name}
}

func NewCollectionService() *collectionService {
	return &collectionService{}
}
