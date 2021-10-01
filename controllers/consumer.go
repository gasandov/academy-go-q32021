package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gasandov/academy-go-q32021/services"

	"github.com/labstack/echo/v4"
)

type apiConsumer struct {}

type APIConsumerController interface {
	ConsumeAPI(c echo.Context) error
}

const apiUrl = "https://pokeapi.co/api/v2/pokemon"

// Consumes an external API and writes the response on a csv file
func (a *apiConsumer) ConsumeAPI(c echo.Context) error {
	limit := c.QueryParams().Get("limit")
	offset := c.QueryParams().Get("offset")

	if limit == "" {
		limit = "100"
	}

	if offset == "" {
		offset = "1"
	}

	endpoint := fmt.Sprintf("%s?limit=%s&offset=%s", apiUrl, limit, offset)

	data, err := http.Get(endpoint)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "API could not be consumed")
	}

	res, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "API response could not be interpreted")
	}

	response, err := services.WriteFile("pokemon_list.csv", res)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "There was an error while creating file")
	}

	return c.JSON(http.StatusOK, response)
}

func NewConsumerController() APIConsumerController {
	return &apiConsumer{}
}
