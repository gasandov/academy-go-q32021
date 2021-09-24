package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"academy-go-q32021/common"
	"academy-go-q32021/services"
)

const FILE_NAME = "pokemon_list.csv"

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Server up and running!")
}

func GetPokemons(c echo.Context) error {
	csvContent, err := services.ReadFile(FILE_NAME)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, pokemonsSlice := common.BuildPokemonCollections(csvContent)

	return c.JSON(http.StatusOK, pokemonsSlice)
}

func GetPokemonById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "ID was not provided")
	}

	csvContent, err := services.ReadFile(FILE_NAME)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pokemonsMap, _ := common.BuildPokemonCollections(csvContent)

	pokemon, exists := pokemonsMap[id]

	if !exists {
		return c.JSON(http.StatusNotFound, "Pokemon was not found")
	}

	return c.JSON(http.StatusOK, pokemon)
}
