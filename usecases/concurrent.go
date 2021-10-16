package usecases

import (
	"encoding/csv"
	"io"
	"os"
	"sync"

	"github.com/gasandov/academy-go-q32021/constants"
)

type ConcurrentService struct {
	repo csvIO
}

var wg sync.WaitGroup
var lock = new(sync.Mutex)

func (ccs *ConcurrentService) GetConcurrently(flag string, items, itemsWorker int64) ([][]string, error) {
	var result [][]string
	content := make(chan []string, items)
	workers := items / itemsWorker

	file, err := os.Open(constants.FileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	for w := int64(0); w < workers; w++ {
		wg.Add(1)
		go worker(reader, flag, itemsWorker, content)
	}

	go func(rows chan []string) {
		wg.Wait()
		close(rows)
	}(content)

	for row := range content {
		result = append(result, row)
	}

	return result, nil
}

func worker(reader *csv.Reader, flag string, itemsWorker int64, content chan<- []string) {
	defer wg.Done()
	var count int64

	for {
		if count == itemsWorker {
			break
		}

		lock.Lock()
		row, err := reader.Read()
		lock.Unlock()

		if err == io.EOF {
			break
		}

		if len(row) == 0 {
			break
		}

		if move(flag, count) {
			content <- row
			count++
		}
	}
}

func move(flag string, count int64) bool {
	switch flag {
	case "odd":
		return count % 2 != 0
	case "even":
		return count % 2 == 0
	default:
		return true
	}
}

func NewConcurrentService(repo csvIO) *ConcurrentService {
	return &ConcurrentService{repo}
}
