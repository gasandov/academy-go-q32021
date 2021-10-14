package controllers

import (
	"net/http"

	"github.com/gasandov/academy-go-q32021/constants"

	"github.com/labstack/echo/v4"
)

type ConsumerHandler struct {
	cService consumerService
	pService pokemonService
}

type consumerService interface {
	Consume(limit, offset string) ([]byte, error)
}

// Expectes limit and offset. Consumes and saves content from API
func (ch *ConsumerHandler) ConsumeAPI(c echo.Context) error {
	limit := c.QueryParams().Get("limit")
	offset := c.QueryParams().Get("offset")

	if limit == "" {
		limit = "100"
	}

	if offset == "" {
		offset = "1"
	}

	response, err := ch.cService.Consume(limit, offset)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := ch.pService.Save(constants.FileName, response)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

 	return c.JSON(http.StatusOK, content)
}

func NewConsumerController(cService consumerService, pService pokemonService) *ConsumerHandler {
	return &ConsumerHandler{cService, pService}
}
