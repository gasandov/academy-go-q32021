package common

import "github.com/gasandov/academy-go-q32021/entities"

// Creates an array and a map of pokemons based on .csv input content
func BuildPokemonCollections(csvLines [][]string) (map[string]entities.Pokemon, []entities.Pokemon) {
	var pokemonsSlice []entities.Pokemon
	pokemonsMap := make(map[string]entities.Pokemon)
	
	for _, line := range csvLines {
		id := line[0]
		name := line[1]

		pokemon := entities.Pokemon{Id: id, Name: name}
		pokemonsMap[id] = pokemon
		pokemonsSlice = append(pokemonsSlice, pokemon)
	}

	return pokemonsMap, pokemonsSlice
}
