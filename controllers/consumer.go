package controllers

import (
	"io/ioutil"
	"net/http"

	"github.com/gasandov/academy-go-q32021/services"

	"github.com/labstack/echo/v4"
)

type apiConsumer struct {}

type APIConsumerController interface {
	ConsumeAPI(c echo.Context) error
}

const apiUrl = "https://pokeapi.co/api/v2/pokemon?limit=100&offset=1"

// Consumes an external API and writes the response on a csv file
func (a *apiConsumer) ConsumeAPI(c echo.Context) error {
	data, err := http.Get(apiUrl)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "API could not be consumed")
	}

	res, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "API response could not be interpreted")
	}

	msg, err := services.WriteFile("pokemon_list.csv", res)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "There was an error while creating file")
	}

	return c.JSON(http.StatusOK, msg)
}

func NewConsumerController() APIConsumerController {
	return &apiConsumer{}
}