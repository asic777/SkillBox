// Задание 1. Определение количества слов, начинающихся с большой буквы
// Что нужно сделать
// Напишите программу, которая выведет количество слов, начинающихся с большой буквы в строке: Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software.

// Рекомендация
// Пример работы программы:
// Определение количества слов, начинающихся с большой буквы в строке:
// Go is an Open source programming Language that makes it Easy
// to build simple, reliable, and efficient Software
// Строка содержит 5 слов с большой буквы.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const text = "a10 10 20b 20 30c30 30 dd"

func main() {
	fmt.Println("Строка для определения чисел в десятичном формате:")
	fmt.Println(text)

	tempSliceText := strings.Split(text, " ")
	resultText := ""
	for i := 0; i < len(tempSliceText); i++ {
		if _, err := strconv.ParseInt(tempSliceText[i], 10, 0); err == nil {
			resultText += tempSliceText[i] + " "
		}
	}

	resultText = strings.Trim(resultText, " ")
	fmt.Printf("Десятичные числа в строке: %v", resultText)
}
