package services

import (
	"encoding/csv"
	"os"
)

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