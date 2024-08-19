package main

import (
	"fmt"
	"mask_adress/services"
	"os"

)

func main() {
	// Получение пути к входному файлу из аргументов запуска
	inputPath := ""
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	} else {
		fmt.Println("Не указан путь к входному файлу")
		return
	}

	// Создание реализаций Producer и Presenter
	prod := services.NewProducerFile(inputPath)
	pres := services.NewPresenterFile("output.txt")

	// Создание сервиса
	service := services.NewService(prod, pres, "C:\\projects\\GO\\output.txt")

	// Запуск сервиса
	err := service.Run(inputPath)
	if err != nil {
		fmt.Println("Ошибка выполнения сервиса:", err)
	}
	// Ввод поискового значения
	//reader = bufio.NewReader(os.Stdin)
	//fmt.Print("Что ищем (http просьба вводить http:\\): ")
	//input2, _ := reader.ReadString('\n')
	//input2 = strings.TrimSpace(input2)
	//var myAdress string = input2
}
