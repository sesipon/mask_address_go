package services

import (
	"fmt"
	"io/ioutil"
	"strings" // strings для проверки расширения файла
)

// интерфейс для поставщика данных
type Producer interface {
	Produce() ([]string, error)
}

// реализация Producer для чтения файла
type ProducerFile struct {
	path string
}

// конструктор ProducerFile
func NewProducerFile(path string) *ProducerFile {
	return &ProducerFile{path: path}
}

// реализация метода Produce для ProducerFile
func (p *ProducerFile) Produce() ([]string, error) {
	data, err := ioutil.ReadFile(p.path)
	if err != nil {
		return nil, err // Возвращаем ошибку
	}
	return strings.Split(string(data), "\n"), nil // Разбиваем на строки и возвращаем
}

// path_to_file - функция для получения пути к файлу от пользователя
func path_to_file(path string) string {
	fmt.Print("Укажите путь к файлу: ")
	fmt.Scanln(&path)
	return path // Возвращаем путь к файлу
}