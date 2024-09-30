package service

import (
	"fmt"
	"os"
)

// реализация Presenter для записи в файл

type PresenterFile struct {
	path string
}

// конструктор PresenterFile

func NewPresenterFile(path string) *PresenterFile {
	return &PresenterFile{path: path}
}

// реализация метода Present для PresenterFile

func (p *PresenterFile) Present(data []string) error {
	file, err := os.Create(p.path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, line := range data {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
