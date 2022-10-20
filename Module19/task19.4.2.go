// Задание 2. Сортировка пузырьком
// Что нужно сделать
// Отсортируйте массив длиной шесть пузырьком.

package main

import "fmt"

func main() {
	fmt.Printf("Задание 19.4.2. Сортировка массивов пузырьком.\n")
	fmt.Printf("--------------------\n")
	firstArray := [6]int{6, 5, 4, 3, 2, 1}
	fmt.Println("Массив до сортировки:", firstArray)
	fmt.Println("Массив после сортировки пузырьком:", sortArray(firstArray))

}

func sortArray(array [6]int) [6]int {
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j-1] > array[j] {
				array[j-1], array[j] = array[j], array[j-1]
			}
		}
	}
	return array
}
