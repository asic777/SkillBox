// Задание 1. Чётные и нечётные
// Что нужно сделать
// Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.

package main

import "fmt"

func main() {
	fmt.Printf("Задание 23.6.1. Функция, которая принимает массив и возвращает два массива: из чётных и нечётных чисел.\n")
	fmt.Printf("--------------------\n")

	arrayNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arrayNumbers, "- исходный массив")
	fmt.Printf("--------------------\n")

	oddNumbersArray, evenNumbersArray := splitArrayOddEvenNumbers(arrayNumbers)

	fmt.Println(oddNumbersArray, "- массив нечетных чисел")
	fmt.Printf("--------------------\n")
	fmt.Println(evenNumbersArray, "- массив четных чисел")
}

func splitArrayOddEvenNumbers(array []int) (oddNumbersArray, evenNumbersArray []int) {
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			evenNumbersArray = append(evenNumbersArray, i)
		} else {
			oddNumbersArray = append(oddNumbersArray, i)
		}
	}
	return
}
