// Задание 1. Функция, принимающая аргументы
// Что нужно сделать
// Напишите функцию, которая принимает в качестве аргументов два числа типа int, вычисляет сумму чётных чисел заданного диапазона и выводит результат в консоль.

// Рекомендация
// Если первое число, переданное в качестве аргумента, будет больше второго, просто поменяйте их местами.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Задание 13.6.1. Вычисление суммы чётных чисел заданного диапазона.")
	fmt.Println("--------------------")

	fmt.Print("Введите первое число: ")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Print("Введите второе число: ")
	var secondNumber int
	fmt.Scan(&secondNumber)

	sumEvenNumberFromRange(firstNumber, secondNumber)
}
func sumEvenNumberFromRange(x1, x2 int) {
	if x1 > x2 {
		x1, x2 = x2, x1
	}

	result := 0
	for i := x1; i <= x2; i++ {
		if i%2 == 0 {
			result += i
		}
	}

	fmt.Printf("Сумма четных чисел в диапозомене между числами %d и %d = %d", x1, x2, result)
}
