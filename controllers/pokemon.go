package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type pokemonService interface {
	ConsumeAPI(limit, offset int64) ([]byte, error)
	StoreData(content []byte) (entities.APIResponse, error)
	GetPokemonsData() (map[string]entities.Pokemon, []entities.Pokemon, error)
	GetPokemonsDataConcurrently(flag string, items, itemsWorker int64) (map[string]entities.Pokemon, error)
}

type PokemonController struct {
	service pokemonService
}

// get all available pokemons
func (pc *PokemonController) GetPokemons(c echo.Context) error {
	pkMap, _, err := pc.service.GetPokemonsData()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, pkMap)
}

// get pokemon by id (if exists)
func (pc *PokemonController) GetPokemonById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id was not provided")
	}

	pkMap, _, err := pc.service.GetPokemonsData()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pokemon, exists := pkMap[id]
	if !exists {
		return c.JSON(http.StatusNotFound, "pokemon not found")
	}

	return c.JSON(http.StatusOK, pokemon)
}

// consumes external api
// write api response in file
// returns initial api response
func (pc *PokemonController) GetPokemonsFromAPI(c echo.Context) error {
	limit, err := strconv.ParseInt(c.QueryParams().Get("limit"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "limit contains invalid value")
	}

	offset, err := strconv.ParseInt(c.QueryParams().Get("offset"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "offset contains invalid value")
	}

	response, err := pc.service.ConsumeAPI(limit, offset)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := pc.service.StoreData(response)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, content)
}

// gets all available pokemons concurrently
func (pc *PokemonController) GetPokemonsConcurrently(c echo.Context) error {
	items, err := strconv.ParseInt(c.QueryParams().Get("items"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "items contains invalid value")
	}

	itemsWorker, err := strconv.ParseInt(c.QueryParams().Get("items_per_workers"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "items_per_workers contains invalid value")
	}

	if itemsWorker > items {
		return c.JSON(http.StatusBadRequest, "items_per_workers should be shorter than items")
	}

	flag := strings.ToLower(c.QueryParams().Get("type"))
	if flag == "" || (strings.Compare(flag, "odd") != 0 && strings.Compare(flag, "even") != 0) {
		flag = "all"
	}

	pkMap, err := pc.service.GetPokemonsDataConcurrently(flag, items, itemsWorker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, pkMap)
}

func NewPokemonController(service pokemonService) *PokemonController {
	return &PokemonController{service}
}
