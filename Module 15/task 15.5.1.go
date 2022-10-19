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
	for i, _ := range arr {
		fmt.Printf("Введите значение элементе массиве [%d]: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("--------------------")
	fmt.Printf("В массиве четных чисел - %d, нечетных - %d", calcMod2(arr), calcNotMod2(arr))
}

func calcMod2(array [10]int) (countMod2 int) {
	for _, r := range array {
		if r%2 == 0 {
			countMod2++
		}
	}
	return
}

func calcNotMod2(array [10]int) (countNotMod2 int) {
	for _, r := range array {
		if r%2 != 0 {
			countNotMod2++
		}
	}
	return
}
