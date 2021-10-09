package usecases

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gasandov/academy-go-q32021/constants"
	"github.com/gasandov/academy-go-q32021/entities"
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

// Receives fileName and content []byte, creates file and writes the content on it
// returns api response
func (cs *ConsumerService) SaveConsumed(fileName string, content []byte) (entities.API, error) {
	file, err := os.Create(fileName)

	if err != nil {
		return entities.API{}, err
	}

	response, err := cs.repo.WriteFile(file, content)

	if err != nil {
		return entities.API{}, err
	}

	return response, nil
}

func NewConsumerService(repo csvIO) *ConsumerService {
	return &ConsumerService{repo}
}
