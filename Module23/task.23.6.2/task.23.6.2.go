// Задание 2. Анонимные функции
// Что нужно сделать
// Напишите анонимную функцию, которая на вход получает массив типа integer, сортирует его пузырьком и переворачивает
// (либо сразу сортирует в обратном порядке, как посчитаете нужным).
// Советы и рекомендации
// В качестве среды разработки используйте GoLand или VSCode.
// Не забудьте проверить, что вы получили больше чем ноль аргументов.
// Подход не важен, можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("--------------------\n")
	fmt.Printf("Задание 23.6.2. Анонимная функция, которая сортирует массив пузырьком в обратном порядке.\n")
	fmt.Printf("--------------------\n")

	var lenArray uint16
	for {
		fmt.Print("Введите длину массива: ")
		if _, err := fmt.Scan(&lenArray); err != nil {
			fmt.Printf("Ошибка! Введите целое чиcло в диапозоне от %d до %d.\n", 1, math.MaxUint16)
		} else {
			break
		}
	}
	fmt.Printf("--------------------\n")

	generateArray := func(lenArr uint16, numRange uint16) (randomArray []int) {
		randomArray = make([]int, lenArr)
		for i := 0; i < len(randomArray); i++ {
			randomArray[i] = rand.Intn(int(numRange))
		}
		return
	}
	array := generateArray(lenArray, 10*lenArray)

	fmt.Println("Сгенерированный массив:", formatArrayOutput(array))
	fmt.Printf("--------------------\n")

	reverseBubbleSort := func(arr *[]int) {
		length := len(*arr)
		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				if (*arr)[i] < (*arr)[j] {
					(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
				}
			}
		}
	}
	reverseBubbleSort(&array)
	fmt.Println("Отсортированный массив:", formatArrayOutput(array))

	fmt.Printf("--------------------\n")
	fmt.Println("Программа завершена")

}

// массивы длинее 500, будут отображаться в сокращенном виде
func formatArrayOutput(arr []int) string {
	if len(arr) > 500 {
		return fmt.Sprintf("[%v %v ... %v %v]", arr[0], arr[1], arr[len(arr)-2], arr[len(arr)-1])
	} else {
		return fmt.Sprintf("%v", arr)
	}
}