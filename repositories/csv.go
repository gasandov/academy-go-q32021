package repositories

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/gasandov/academy-go-q32021/entities"
)

type csvIO struct {}

type CSVRepo interface {
	ReadFile(name string) ([][]string, error)
	WriteFile(file *os.File, data []byte) (entities.API, error)
	CreateFile(name string) (*os.File, error)
	FileExists(name string) bool
}

// Receives a name. Read and returns file content
func (cv *csvIO) ReadFile(name string) ([][]string, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

// Receives a name. Creates and return a file
func (cv *csvIO) CreateFile(name string) (*os.File, error) {
	file, err := os.Create(name)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// Receives a file and data. Writes data on file and returns api response
func (cv *csvIO) WriteFile(file *os.File, data []byte) (entities.API, error) {
	var content entities.API

	json.Unmarshal(data, &content)

	writer := csv.NewWriter(file)

	for i := 0; i < len(content.Results); i++ {
		var row []string

		row = append(row, content.Results[i].Name)
		row = append(row, content.Results[i].Url)

		if err := writer.Write(row); err != nil {
			return entities.API{}, err
		}
	}

	defer writer.Flush()
	defer file.Close()

	return content, nil
}

// Receives name. Check if file exists
func (cv *csvIO) FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func NewCSVRepo() CSVRepo {
	return &csvIO{}
}
