// Задание 1. Подсчёт чётных и нечётных чисел в массиве
// Что нужно сделать
// Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел.
// Для ввода и подсчёта используйте разные циклы.

// Что оценивается
// Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.

package main

import "fmt"

func main() {
	fmt.Println("Задание 15.5.1. Подсчёт чётных и нечётных чисел в массиве.")
	fmt.Println("--------------------")

	//var arr = [10]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4}
	var arr [10]int
	for i := range arr {
		fmt.Printf("Введите значение элемента массива [%d]: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("--------------------")
	numbersEven, numbersNotEven := calcEven(arr)
	fmt.Printf("В массиве четных чисел - %d, нечетных - %d", numbersEven, numbersNotEven)
}

func calcEven(array [10]int) (countMod2, countNotMod2 int) {
	for _, r := range array {
		if r%2 == 0 {
			countMod2++
		} else {
			countNotMod2++
		}
	}
	return
}
