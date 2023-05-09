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
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Задание 23.6.2. Поиск символов в нескольких строках.\n")
	fmt.Printf("--------------------\n")

	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	// for i := 0; i < 4; i++ {
	// 	fmt.Println(len(sentences[i]))
	// }

	fmt.Println(parseTest(sentences, chars))

}
func parseTest(sentences []string, chars []rune) (map[string]int, error) {
	result := make(map[string]int)

	for _, c := range chars {
		for _, s := range sentences {
			s = strings.ToUpper(s)
			index := strings.LastIndexAny(s, c)
			if index != -1 {
				result[string(c)] = index
			}
		}

	}
	return result, nil
}
