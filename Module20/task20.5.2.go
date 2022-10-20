// Задание 2. Умножение матриц
// Что нужно сделать
// Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("Задание 20.5.2. Умножение первой матрицы 3 х 5 на вторую 5 х 4.\n")
	fmt.Printf("--------------------\n")
	matrixA := [][]int{
		{2, 2, 1, 8, 3},
		{6, 5, 6, 2, 9},
		{7, 1, 9, 2, 4},
	}
	matrixB := [][]int{
		{2, 2, 1, 7},
		{1, 3, 6, 5},
		{7, 9, 3, 1},
		{1, 3, 5, 4},
		{4, 2, 8, 9},
	}
	fmt.Println("Первая матрица 3 х 5:")
	for i := 0; i < len(matrixA); i++ {
		fmt.Println(matrixA[i])
	}
	fmt.Printf("--------------------\n")
	fmt.Println("Вторая матрица 5 х 4:")
	for i := 0; i < len(matrixB); i++ {
		fmt.Println(matrixB[i])
	}
	fmt.Printf("--------------------\n")

	if matrixC, err := productMatrix(matrixA, matrixB); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат умножения первой матрицы на вторую:")
		for i := 0; i < len(matrixC); i++ {
			fmt.Println(matrixC[i])
		}
	}
}

func productMatrix(matrix1 [][]int, matrix2 [][]int) ([][]int, error) {
	if len(matrix1[0]) != len(matrix2) {
		return nil, errors.New("Матрица 1 не может быть умножена на матрицу 2.")
	}
	resultMatrix := make([][]int, len(matrix1))
	for i := 0; i < len(matrix1); i++ {
		resultMatrix[i] = make([]int, len(matrix2[0]))
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				resultMatrix[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return resultMatrix, nil
}
