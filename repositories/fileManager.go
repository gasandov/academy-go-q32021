package repositories

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/gasandov/academy-go-q32021/entities"
	"github.com/gasandov/academy-go-q32021/utils"
)

type FileManager struct {}

var wg sync.WaitGroup
var lock = new(sync.Mutex)

// receives filename with extension
// returns fiel
func (fm *FileManager) CreateFile(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// receives file and (data []byte)
// returns parsed data
func (fm *FileManager) WriteFile(file *os.File, data []byte) (entities.APIResponse, error) {
	var content entities.APIResponse

	json.Unmarshal(data, &content)

	writer := createCSVWriter(file)

	for i := 0; i < len(content.Results); i++ {
		var row []string

		row = append(row, content.Results[i].Name)
		row = append(row, content.Results[i].Url)

		if err := writer.Write(row); err != nil {
			return entities.APIResponse{}, err
		}
	}

	defer func() {
		writer.Flush()
		file.Close()
	}()

	return content, nil
}

// receives file name with extension
// returns file content ([][]string)
func (fm *FileManager) ReadFile(name string) ([][]string, error) {
	file, err := fm.OpenFile(name, "")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := createCSVReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// receives file name with extension
func (fm *FileManager) ReadFileConcurrently(name, flag string, items, itemsWorker int64) ([][]string, error) {
	file, err := fm.OpenFile(name, "")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var result [][]string

	content := make(chan []string, items)
	workers := items / itemsWorker

	reader := createCSVReader(file)

	for w := int64(0); w < workers; w++ {
		wg.Add(1)
		go readerWorker(reader, flag, itemsWorker, content)
	}

	go func(row chan []string) {
		wg.Wait()
		close(row)
	}(content)

	for row := range content {
		result = append(result, row)
	}

	return result, nil
}

// receives file name with extension
// return file existance (bool)
func (fm *FileManager) FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// receives file name with extension
// returns file
func (fm *FileManager) OpenFile(name, flag string) (*os.File, error) {
	switch flag {
		case "append":
			file, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return nil, err
			}
		
			return file, nil
		default:
			file, err := os.OpenFile(name, os.O_RDONLY, 0664)
			if err != nil {
				return nil, err
			}
		
			return file, nil
	}
}

// receives csv file
// returns writer
func createCSVWriter(file *os.File) *csv.Writer {
	return csv.NewWriter(file)
}

// receives csv file
// returns reader
func createCSVReader(file *os.File) *csv.Reader {
	return csv.NewReader(file)
}

// receives reader, flag ("odd", "even", "all"), itemsWorker and write channel content
// reads file line by line and writes content into content channel
func readerWorker(r *csv.Reader, flag string, itemsWorker int64, content chan<- []string) {
	defer wg.Done()

	var counter int64

	for {
		if counter == itemsWorker {
			break
		}

		lock.Lock()
		row, err := r.Read()
		lock.Unlock()

		if err == io.EOF {
			break
		}

		if len(row) == 0 {
			break
		}

		if utils.SwitchTo(flag, counter) {
			content <- row
		}
	
		counter++
	}
}

func NewFileManagerRepo() *FileManager {
	return &FileManager{}
}