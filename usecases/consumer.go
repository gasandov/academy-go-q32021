package usecases

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gasandov/academy-go-q32021/constants"
)

type ConsumerService struct {
	repo csvIO
}

// Receives limit and offset, calls api endpoint with query params
// and returns response (body []btye)
func (cs *ConsumerService) Consume(limit, offset string) ([]byte, error) {
	endpoint := fmt.Sprintf("%s?limit=%s&offset=%s", constants.APIUrl, limit, offset)

	data, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	res, err := ioutil.ReadAll(data.Body)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewConsumerService(repo csvIO) *ConsumerService {
	return &ConsumerService{repo}
}
