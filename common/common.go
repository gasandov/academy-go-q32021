package common

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func BuildPokemonCollections(csvLines [][]string) (map[string]Pokemon, []Pokemon) {
	var pokemonsSlice []Pokemon
	pokemonsMap := make(map[string]Pokemon)
	
	for _, line := range csvLines {
		id := line[0]
		name := line[1]

		pokemon := Pokemon{id, name}
		pokemonsMap[id] = pokemon
		pokemonsSlice = append(pokemonsSlice, pokemon)
	}

	return pokemonsMap, pokemonsSlice
}
