package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func maskAdress(text, adress string) string {
	buf := []byte(text)
	example := []byte(adress)

	// Ищем вхождения example
	for i := 0; i <= len(buf)-len(example); i++ {
		match := true
		for j := 0; j < len(example); j++ {
			if buf[i+j] != example[j] {
				match = false
				break
			}
		}

		if match {
			// Маскируем символы после example
			for k := i + len(example); k < len(buf) && buf[k] != ' '; k++ {
				buf[k] = '*'
			}
		}
	}

	fmt.Println(string(buf))
	return string(buf)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var myText string = input

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Что ищем (http просьба вводить http:\\): ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	var myAdress string = input2

	maskAdress(myText, myAdress)
}
