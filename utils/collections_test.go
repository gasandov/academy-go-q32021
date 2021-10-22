package utils

import (
	"testing"

	"github.com/gasandov/academy-go-q32021/entities"
	"github.com/stretchr/testify/assert"
)

var content = [][]string{
	{
		"unown",
		"https://pokeapi.co/api/v2/pokemon/201/",
	},
	{
		"wobbuffet",
		"https://pokeapi.co/api/v2/pokemon/202/",
	},
	{
		"girafarig",
		"https://pokeapi.co/api/v2/pokemon/203/",
	},
}

var badContent = [][]string{
	{
		"https://pokeapi.co/api/v2/pokemon/201/",
		"unown",
	},
}

var expectedMap = map[string]entities.Pokemon{
	"201": {
		Id: "201",
		Name: "unown",
	},
	"202": {
		Id: "202",
		Name: "wobbuffet",
	},
	"203": {
		Id: "203",
		Name: "girafarig",
	},
}

func TestCollectionsUtils_BuildCollections(t *testing.T) {
	testCases := []struct {
		name string
		content [][]string
		expectedLength int
		expectedMap map[string]entities.Pokemon
		hasError bool
	} {
		{
			"should build pokemon map successfully",
			content,
			3,
			expectedMap,
			false,
		},
		{
			"should return an error on bad content",
			badContent,
			3,
			nil,
			true,
		},
		{
			"should return empty map on no content",
			[][]string{},
			0,
			map[string]entities.Pokemon{},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.hasError {
				assert.Panics(t, func() { BuildCollections(tc.content) })
			} else {
				pkMap, _ := BuildCollections(tc.content)
	
				assert.Len(t, pkMap, tc.expectedLength)
				assert.Equal(t, pkMap, tc.expectedMap)
			}
		})
	}
}
