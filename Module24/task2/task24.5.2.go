// Задание 2. Поиск символов в нескольких строках
// Что нужно сделать
// Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune,
// а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово
// в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
// func parseTest(sentences []string, chars []rune)
// Советы и рекомендации
// Не забудьте проверить, что вы получили больше чем 0 аргументов.
// Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
// Пример входных данных
// sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
// chars := [5]rune{'H','E','L','П','М'}
// Пример вывода результата в первом элементе массива
// 'H' position 0
// 'E' position 1
// 'L' position 9

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("----------------------------------------------------\n")
	fmt.Printf("Задание 24.5.2. Поиск символов в нескольких строках.\n")
	fmt.Printf("----------------------------------------------------\n\n")

	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	fmt.Printf("Исходный массив предложений:\t[%s]\n", strings.Join(sentences, "] ["))
	fmt.Printf("Исходный массив рун:\t\t%s\n", strings.ReplaceAll(fmt.Sprintf("%c", chars), " ", "] ["))
	fmt.Printf("---------------------------------------------------------------------------------------------\n\n")

	resultArray, err := parseTest(sentences, chars)
	if err != nil {
		fmt.Println("Ошибка: ", err)
		return
	}

	fmt.Printf("Index Rune in Sentinel\t%3c\n", chars)
	for i := range sentences {
		fmt.Printf("---------------------------------------------\n")
		fmt.Printf("%s\t\t", sentences[i])
		fmt.Printf("%s\n", strings.ReplaceAll(fmt.Sprintf("%3d", resultArray[i]), "-1", " -"))
	}
	fmt.Printf("---------------------------------------------\n\n")
	fmt.Println("Программа завершена.")
}

func parseTest(sentences []string, chars []rune) ([][]int, error) {
	if len(sentences) == 0 || len(chars) == 0 {
		return nil, errors.New("the size of the incoming slice is 0")
	}

	result := make([][]int, len(sentences))
	for i := range result {
		result[i] = make([]int, len(chars))
	}

	for i, s := range sentences {
		s = strings.ToUpper(s)
		arrRunes := []rune(s)

		for j, c := range chars {
			index := -1
			for k := len(arrRunes) - 1; k >= 0; k-- {
				if arrRunes[k] == c {
					index = k
				}
				if arrRunes[k] == ' ' && index != -1 {
					k = -1
					break
				}
			}
			result[i][j] = index
		}
	}

	return result, nil
}
