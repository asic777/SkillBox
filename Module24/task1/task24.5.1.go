// Задание 1. Чётные и нечётные
// Что нужно сделать
// Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.

package main

import "fmt"

func main() {
	fmt.Printf("-----------------------------------------------------------------------------------------------\n")
	fmt.Println("Задание 24.5.1. Функция, которая принимает массив и возвращает два: из чётных и нечётных чисел.")
	fmt.Printf("-----------------------------------------------------------------------------------------------\n\n")

	arrayNumbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	oddNumbers, evenNumbers := splitArrayOddEvenNumbers(arrayNumbers)

	fmt.Printf("Исходный массив:\t%v\n---------------------------------------------\n", arrayNumbers)
	fmt.Printf("Массив нечетных чисел:\t%v\n---------------------------------------------\n", oddNumbers)
	fmt.Printf("Массив четных чисел:\t%v\n---------------------------------------------\n\n", evenNumbers)
	fmt.Println("Программа завершена.")
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
