package controllers

import (
	"net/http"

	"github.com/gasandov/academy-go-q32021/entities"

	"github.com/labstack/echo/v4"
)

type ConsumerHandler struct {
	service consumerService
}

type consumerService interface {
	Consume(limit, offset string) ([]byte, error)
	SaveConsumed(fileName string, content []byte) (entities.API, error)
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

	response, err := ch.service.Consume(limit, offset)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	content, err := ch.service.SaveConsumed(fileName, response)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

 	return c.JSON(http.StatusOK, content)
}

func NewConsumerController(service consumerService) *ConsumerHandler {
	return &ConsumerHandler{service}
}
