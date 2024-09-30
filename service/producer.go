package service

import (
	"bufio"
	"fmt"
	"os"
)

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
	file, err := os.Open(p.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var messages []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}
