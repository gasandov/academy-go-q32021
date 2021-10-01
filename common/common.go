package common

import (
	"github.com/gasandov/academy-go-q32021/entities"
	"github.com/gasandov/academy-go-q32021/utils"
)

// Creates an array and a map of pokemons based on .csv input content
func BuildPokemonCollections(csvLines [][]string) (map[string]entities.Pokemon, []entities.Pokemon) {
	var pokemonsSlice []entities.Pokemon
	pokemonsMap := make(map[string]entities.Pokemon)
	
	for _, line := range csvLines {
		pokemon := utils.ParsePokemon(line)

		pokemonsMap[pokemon.Id] = pokemon
		pokemonsSlice = append(pokemonsSlice, pokemon)
	}

	return pokemonsMap, pokemonsSlice
}
