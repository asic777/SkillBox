// Задание 2. Функция, реверсирующая массив
// Что нужно сделать
// Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке
// по сравнению с исходным.
// Напишите программу, демонстрирующую работу этого метода.
// Что оценивается
// При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции,
// реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.

package main

import "fmt"

func main() {
	fmt.Println("Задание 15.5.2. Подсчёт чётных и нечётных чисел в массиве.")
	fmt.Println("--------------------")

	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Print("Введите размер массива: ")
	var lengthArray int
	fmt.Scan(&lengthArray)

	arr := make([]int, lengthArray)

	for i := range arr {
		fmt.Printf("Введите значение элемента массива [%d]: ", i)
		fmt.Scan(&arr[i])
	}
	fmt.Printf("Начальный массив:\t%v\n", arr)
	fmt.Printf("Обратный массив:\t%v\n", reverseArray(arr))
}

func reverseArray(arr []int) []int {
	sizeArray := len(arr)
	var tempInt int
	for i := 0; i < sizeArray/2; i++ {
		tempInt = arr[i]
		arr[i] = arr[sizeArray-1-i]
		arr[sizeArray-1-i] = tempInt
	}
	return arr
}
