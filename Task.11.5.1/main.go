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
	"strings"
)

const text = "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"

func main() {
	fmt.Printf("Строка для подсчета слов с большой буквы: \n%v\n", text)

	tempText := strings.Title(text)
	tempSliceText := strings.Split(tempText, " ")
	countWordTitle := 0
	for i := 0; i < len(tempSliceText); i++ {
		if strings.Contains(text, tempSliceText[i]) == true {
			countWordTitle++
		}
	}

	fmt.Printf("Строка содержит %v слов с большой буквы.", countWordTitle)
}
