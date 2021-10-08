package usecases

import (
	"testing"

	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/stretchr/testify/assert"
)

var pokemons = [][]string{
	{
		"charmander",
		"https://pokeapi.co/api/v2/pokemon/4/",
	},
	{
		"charizard",
		"https://pokeapi.co/api/v2/pokemon/6/",
	},
	{
		"metapod",
		"https://pokeapi.co/api/v2/pokemon/11/",
	},
}

var expectedPkMap = map[string]entities.Pokemon{
	"4": {
		Id: "4",
		Name: "charmander",
	},
	"6": {
		Id: "6",
		Name: "charizard",
	},
	"11": {
		Id: "11",
		Name: "metapod",
	},
}

var expectedPkSlice = []entities.Pokemon{
	{
		Id: "4",
		Name: "charmander",
	},
	{
		Id: "6",
		Name: "charizard",
	},
	{
		Id: "11",
		Name: "metapod",
	},
}

func TestCollectionService_BuildCollections(t *testing.T) {
	t.Run("build collections successfully", func(t *testing.T) {
		service := NewCollectionService()

		pkMap, pkSlice := service.BuildCollections(pokemons)

		assert.EqualValues(t, pkMap, expectedPkMap, "pokemon map should be built")
		assert.EqualValues(t, pkSlice, expectedPkSlice, "pokemon slice should be built")
	})

	t.Run("collections should be empty", func(t *testing.T) {
		service := NewCollectionService()
		var emptyContent [][]string

		pkMap, pkSlice := service.BuildCollections(emptyContent)

		assert.Empty(t, pkMap, "pokemon map should be empty")
		assert.Empty(t, pkSlice, "pokemon slice should be empty")
	})
}