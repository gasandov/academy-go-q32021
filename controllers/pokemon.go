package controllers

import (
	"net/http"

	"github.com/gasandov/academy-go-q32021/repositories"
	"github.com/gasandov/academy-go-q32021/usecases"

	"github.com/labstack/echo/v4"
)

const fileName = "pokemon_list.csv"

type pokemon struct {}

type PokemonController interface {
	GetPokemons(c echo.Context) error
	GetPokemonById(c echo.Context) error
}

// Reads csv file and send to the client an
// array of pokemons [{ id: #, name: string }]
func (p *pokemon) GetPokemons(c echo.Context) error {
	csvContent, err := repositories.ReadFile(fileName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, pokemonsSlice := usecases.BuildPokemonCollections(csvContent)

	return c.JSON(http.StatusOK, pokemonsSlice)
}

// Receives a param "id", reads from a csv file and sends
// to the client a single pokemon { id: #, name: string }
func (p *pokemon) GetPokemonById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "ID was not provided")
	}

	csvContent, err := repositories.ReadFile(fileName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pokemonsMap, _ := usecases.BuildPokemonCollections(csvContent)

	pokemon, exists := pokemonsMap[id]

	if !exists {
		return c.JSON(http.StatusNotFound, "Pokemon was not found")
	}

	return c.JSON(http.StatusOK, pokemon)
}

func NewPokemonController() PokemonController {
	return &pokemon{}
}