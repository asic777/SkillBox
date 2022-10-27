// Задание 1. Подсчёт элементов в массиве после заданного числа
// Что нужно сделать
// Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число.
// Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
// При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Задание 22.4.1. Подсчёт элементов в массиве после заданного числа.\n")
	fmt.Printf("--------------------\n")

	array := generateArray(10, 10)
	fmt.Println(array)
	fmt.Printf("--------------------\n")

	var number int
	for {
		fmt.Print("Введите число для поиска в массиве: ")
		if _, err := fmt.Scan(&number); err != nil {
			fmt.Printf("Ошибка! Введите целое чиcло в диапозоне от %d до %d.\n", math.MinInt, math.MaxInt)
		} else {
			break
		}
	}
	fmt.Printf("--------------------\n")

	if countElementsAfterIndex := arrayElementsAfterIndex(array, int(number)); countElementsAfterIndex != -1 {
		fmt.Printf("Число найдено в массиве. Элементов после числа: %v\n", countElementsAfterIndex)
	} else {
		fmt.Println("Число не найдено в массиве.")
	}
}

func generateArray(lengthArray, numbersRange int) (resultArray []int) {
	resultArray = make([]int, lengthArray)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lengthArray; i++ {
		resultArray[i] = rand.Intn(numbersRange)
	}
	return
}

func arrayElementsAfterIndex(array []int, number int) int {
	index := -1
	countElementsAfterIndex := -1

	for i := 0; i < len(array); i++ {
		if array[i] == number {
			index = i
			countElementsAfterIndex = len(array[index:]) - 1
			break // для вывода числа элементов после первого вхождения в массив
			// для вывода числа элементов после последнего вхождения в массив нужно убрать break
		}
	}
	if index == -1 {
		return countElementsAfterIndex
	}
	return countElementsAfterIndex
}
