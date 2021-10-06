package usecases

import (
	"strings"

	"github.com/gasandov/academy-go-q32021/entities"
)

// Creates an array and a map of pokemons based on .csv input content
func BuildPokemonCollections(csvLines [][]string) (map[string]entities.Pokemon, []entities.Pokemon) {
	var pokemonsSlice []entities.Pokemon
	pokemonsMap := make(map[string]entities.Pokemon)
	
	for _, line := range csvLines {
		pokemon := parsePokemon(line)

		pokemonsMap[pokemon.Id] = pokemon
		pokemonsSlice = append(pokemonsSlice, pokemon)
	}

	return pokemonsMap, pokemonsSlice
}

// Receive array row of pokemon [name, url], parse url and return Pokemon { id: <int>, name: <string>}
func parsePokemon(row []string) entities.Pokemon {
	name := row[0]
	url := row[1]
	splited := strings.Split(url, "/")

	id := splited[len(splited) - 2]

	return entities.Pokemon{Id: id, Name: name}
}
