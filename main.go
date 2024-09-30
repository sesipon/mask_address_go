package main

import (
	"fmt"
	"os"

	"mask_adress/service"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		return
	}

	inputPath := os.Args[1]
	outputPath := "output.txt"
	if len(os.Args) > 2 {
		outputPath = os.Args[2]

		// Создание реализаций Producer и Presenter
		prod := service.NewProducerFile(inputPath)
		pres := service.NewPresenterFile(outputPath)

		// Создание сервиса
		svc := service.NewService(prod, pres)

		// Запуск сервиса
		err := svc.Run()
		if err != nil {
			fmt.Println("Ошибка выполнения сервиса:", err)
		}
	}
}
