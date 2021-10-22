package utils

import (
	"strings"

	"github.com/gasandov/academy-go-q32021/entities"
)

// receives a [[name url]]
// returns map { 1: { Id: 1, Name: "charmander" } }
// returns slice [{ id: 1, name: "charmander" }]
func BuildCollections(content [][]string) (map[string]entities.Pokemon, []entities.Pokemon) {
	var pkSlice []entities.Pokemon
	pkMap := make(map[string]entities.Pokemon)

	for _, line := range content {
		pokemon := parsePokemon(line)

		pkMap[pokemon.Id] = pokemon
		pkSlice = append(pkSlice, pokemon)
	}
	
	return pkMap, pkSlice
}

// receives [charmander https://pokeapi.co/api/v2/pokemon/1]
// returns a Pokemon{Id: 1, Name: "charmander"}
func parsePokemon(line []string) entities.Pokemon {
	name := line[0]
	url := line[1]
	splited := strings.Split(url, "/")

	id := splited[len(splited) - 2]

	return entities.Pokemon{Id: id, Name: name}
}
