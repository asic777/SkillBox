// Задание 1. Подсчёт определителя
// Что нужно сделать
// Напишите функцию, вычисляющую определитель матрицы размером 3 × 3.

package main

import "fmt"

func main() {
	fmt.Printf("Задание 20.5.1. Вычисление определителя матрицы размером 3 × 3.\n")
	fmt.Printf("--------------------\n")
	// matrix := [][]int{
	// 	{2, 2, 1},
	// 	{6, 5, 6},
	// 	{7, 1, 9},
	// }

	matrix := [3][3]int{}
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("Введите значение элемента матрицы [%d][%d]: ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Printf("--------------------\n")
	fmt.Println("Введенная матрица 3 х 3:")
	for i := range matrix {
		fmt.Println(matrix[i])
	}
	fmt.Printf("--------------------\n")
	fmt.Println("Определитель матрицы равен:", detMatrix(matrix))
}

func detMatrix(matrix [3][3]int) (result int) {
	result = (matrix[0][0]*matrix[1][1]*matrix[2][2] -
		matrix[0][0]*matrix[1][2]*matrix[2][1] -
		matrix[0][1]*matrix[1][0]*matrix[2][2] +
		matrix[0][1]*matrix[1][2]*matrix[2][0] +
		matrix[0][2]*matrix[1][0]*matrix[2][1] -
		matrix[0][2]*matrix[1][1]*matrix[2][0])
	return
}
