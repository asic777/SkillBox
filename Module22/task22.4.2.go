// Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
// Что нужно сделать
// Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив.
// Сложность алгоритма должна быть минимальная.

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

	//array := []int{1, 1, 2, 2, 3, 5, 7, 8, 9, 9}
	array := generateArray(20, 100)
	array = sortArrayBubble(array)
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

	if index := indexOfArray(array, number); index != -1 {
		fmt.Printf("Первое вхождение числа в массиве с индексом: %d\n", index)
	} else {
		fmt.Println("Число не найдено в массиве.")
	}
}

func indexOfArray(array []int, numberForSearch int) (result int) {
	min := 0
	max := len(array) - 1
	for min <= max {
		result = (min + max) / 2
		if numberForSearch < array[result] {
			max = result - 1
		} else if numberForSearch > array[result] {
			min = result + 1
		} else if result == 0 || numberForSearch != array[result-1] {
			return result
		} else {
			max = result - 1
		}
	}

	return -1
}

func generateArray(lengthArray, numbersRange int) (resultArray []int) {
	resultArray = make([]int, lengthArray)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < lengthArray; i++ {
		resultArray[i] = rand.Intn(numbersRange)
	}
	return
}

func sortArrayBubble(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array
}
