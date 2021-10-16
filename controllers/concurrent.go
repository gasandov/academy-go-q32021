package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ConcurrentHandler struct {
	service concurrentService
}

type concurrentService interface {
	GetConcurrently(flag string, items, itemsWorker int64) ([][]string, error)
}

// Expects type, items and items_per_worker. Reads content file concurrently
func (cc *ConcurrentHandler) GetPokemonsConcurrently(c echo.Context) error {
	items, err := strconv.ParseInt(c.QueryParams().Get("items"), 10, 64) // how many items the endpoint should return
	if err != nil {
		return c.JSON(http.StatusBadRequest, "items contains invalid value")
	}

	itemsWorker, err := strconv.ParseInt(c.QueryParams().Get("items_per_workers"), 10, 64) // how many items the worker returns
	if err != nil {
		return c.JSON(http.StatusBadRequest, "items_per_workers contains invalid value")
	}

	if itemsWorker > items {
		return c.JSON(http.StatusBadRequest, "items_per_workers should be shorter than items")
	}

	flag := strings.ToLower(c.QueryParams().Get("type")) // if not exists return all, else return even or odd
	if flag != "odd" || flag != "even" {
		flag = "all"
	}

	pokemons, err := cc.service.GetConcurrently(flag, items, itemsWorker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	return c.JSON(http.StatusOK, pokemons)
}

func NewConcurrentController(service concurrentService) *ConcurrentHandler {
	return &ConcurrentHandler{service}
}