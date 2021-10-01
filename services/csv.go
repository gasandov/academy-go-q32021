package services

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/gasandov/academy-go-q32021/entities"
)

// Receives a file name with .csv extension and returns
// the content from it as array
func ReadFile(fileName string) ([][]string, error) {
	csvFile, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		return nil, err
	}

	return csvLines, nil
}

// Receives a file name with extension and returns
// a success string or an error
func WriteFile(fileName string, data []byte) (string, error) {
	var apiResponse entities.APIResponse

	json.Unmarshal(data, &apiResponse)

	file, err := os.Create(fileName)

	defer file.Close()

	if err != nil {
		return "", err
	}

	writer := csv.NewWriter(file)

	for i := 0; i < len(apiResponse.Results); i++ {
		var row []string

		row = append(row, apiResponse.Results[i].Name)
		row = append(row, apiResponse.Results[i].Url)

		if err := writer.Write(row); err != nil {
			return "", err
		}
	}

	defer writer.Flush()

	return "File was created successfully", nil
}
