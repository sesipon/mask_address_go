package main

import "fmt"

func mask_adress(adress string) {
	buf := []byte(adress)

	// Ищем последовательность от 104 116 116 112 (http)
	start := -1
	for i, b := range buf {
		if b == 104 && i+3 < len(buf) && buf[i+1] == 116 && buf[i+2] == 116 && buf[i+3] == 112 {
			start = i
			break
		}
	}

	// Находим пробел 32 если есть http
	end := -1 // Объявляем end вне if
	if start > 0 {
		for i, b := range buf[start:] {
			if b == 32 {
				end = i + start
				break
			}
		}
	}
	// Замена на * и выводим результат
	if start >= 0 && end >= 0 {
		for i := start + 4; i <= end-1; i++ {
			buf[i] = '*'
		}
		fmt.Println(string(buf))
	} else {
		fmt.Println(string(buf))
	}
}

func main() {
	mask_adress("Hello, its my page:  See you")
}
