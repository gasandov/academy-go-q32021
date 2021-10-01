package utils

import (
	"strings"

	"github.com/gasandov/academy-go-q32021/entities"
)

// Receive array row of pokemon [name, url], parse url and return Pokemon { id: <int>, name: <string>}
func ParsePokemon(row []string) entities.Pokemon {
	name := row[0]
	url := row[1]
	splited := strings.Split(url, "/")

	id := splited[len(splited) - 2]

	return entities.Pokemon{Id: id, Name: name}
}