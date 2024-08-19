package services

import (
	"io/ioutil"
	"strings"
)

// интерфейс для обработчика результата
type Presenter interface {
	Present(data []string) error
}

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
	output := strings.Join(data, "\n")
	return ioutil.WriteFile(p.path, []byte(output), 0644) //0644 флаг для записи файла (запись для владельца и чтение для всех
}