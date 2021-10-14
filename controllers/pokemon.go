package controllers

import (
	"net/http"

	"github.com/gasandov/academy-go-q32021/constants"
	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type PokemonHandler struct {
	service pokemonService
}

type pokemonService interface {
	Get(fileName string) (map[string]entities.Pokemon, []entities.Pokemon, error)
	Save(fileName string, content []byte) (entities.API, error)
}

// Get all available pokemons
func (ph *PokemonHandler) GetPokemons(c echo.Context) error {
	_, pkSlice, err := ph.service.Get(constants.FileName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, pkSlice)
}

// Get pokemon by id (if any)
func (ph *PokemonHandler) GetPokemonById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "id was not provided")
	}

	pkMap, _, err := ph.service.Get(constants.FileName)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pokemon, exists := pkMap[id]

	if !exists {
		return c.JSON(http.StatusNotFound, "pokemon not found")
	}

 	return c.JSON(http.StatusOK, pokemon)
}

func NewPokemonController(service pokemonService) *PokemonHandler {
	return &PokemonHandler{service}
}
