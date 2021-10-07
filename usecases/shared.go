package usecases

import (
	"os"

	"github.com/gasandov/academy-go-q32021/entities"
)

type csvIO interface {
	ReadFile(name string) ([][]string, error)
	WriteFile(file *os.File, data []byte) (entities.API, error)
	CreateFile(name string) (*os.File, error)
	FileExists(name string) bool
}
