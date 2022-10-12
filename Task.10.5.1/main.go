// Задание 1. Разложение ex в ряд Тейлора
// Что нужно сделать
// Напишите программу, вычисляющую ex посредством разложения в ряд Тейлора с заданной пользователем точностью. Пользователь вводит значение x и с точностью до какого знака после запятой необходимо вычислить значение функции.

// Советы и рекомендации
// Получить значение точности (эпсилон) можно, разделив 1 на 10, возведённую в степень, равную количеству знаков после запятой.

// Для x = 0 и n = 1 вывод должен быть 1
// Для x = 1 и n = 3 вывод должен быть 2,7182539682539684
// Для x = 1 и n = 5 вывод должен быть 2,7182815255731922

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вычисление функции f(x) = exp(x)")

	var x float64
	fmt.Print("Введите значение x: ")
	fmt.Scan(&x)

	var n int
	for {
		fmt.Print("Задайте значение точности n (разрядов после запятой): ")
		fmt.Scan(&n)
		if n <= 0 {
			fmt.Print("Введите положительное целое чиcло\n")
		} else {
			break
		}
	}
	precision := 1 / math.Pow(10, float64(n))
	fmt.Printf("exp(%v) с точностью %v = %v", x, n, exp(x, precision))
}

func fact(n uint) uint {
	if n <= 1 {
		return 1
	}
	return n * fact(n-1)
}

func exp(x float64, epsilon float64) float64 {
	result := 0.0
	currentMemberFunction := 0.0
	var n uint = 0
	for {
		currentMemberFunction = math.Pow(x, float64(n)) / float64(fact(n))
		result += currentMemberFunction
		if currentMemberFunction < epsilon {
			return result
		}
		n++
	}
}
