package main

import (
	"encoding/csv"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/health-check", HealthCheck)

	e.GET("/pokemons", GetPokemons)
	e.GET("/pokemons/:id", GetPokemonById)

	e.Logger.Fatal(e.Start(":8081"))
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Server up and running!")
}

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func GetPokemons(c echo.Context) error {
	csvContent, _ := ReadFile()
	pokemons := BuildPokemonSlice(csvContent)

	return c.JSON(http.StatusOK, pokemons)
}

func GetPokemonById(c echo.Context) error {
	id := c.Param("id")

	csvContent, _ := ReadFile()
	pokemons := BuildPokemonMap(csvContent)

	pokemon := pokemons[id]

	return c.JSON(http.StatusOK, pokemon)
}

func ReadFile() ([][]string, error) {
	csvFile, err := os.Open("pokemon_list.csv")

	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

func BuildPokemonMap(csvLines [][]string) map[string]Pokemon {
	pokemons := make(map[string]Pokemon)

	for _, line := range csvLines {
		id := line[0]
		name := line[1]

		pokemons[id] = Pokemon{id, name}
	}

	return pokemons
}

func BuildPokemonSlice(csvLines [][]string) []Pokemon {
	var pokemons []Pokemon

	for _, line := range csvLines {
		id := line[0]
		name := line[1]

		pokemon := Pokemon{id, name}

		pokemons = append(pokemons, pokemon)
	}

	return pokemons
}