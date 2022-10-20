// Задание 2. Анонимные функции
// Что нужно сделать
// Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает
// и вызывает её при выходе (через defer).

// Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное,
// чтобы все три выполняли разное действие.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Задание 21.6.2. Анонимные функции.\n")
	fmt.Printf("--------------------\n")

	var firstNumber int
	for {
		fmt.Print("Введите первое число: ")
		if _, err := fmt.Scan(&firstNumber); err != nil {
			fmt.Printf("Ошибка! Введите целое чило в диапозоне от %d до %d.\n", math.MinInt, math.MaxInt)
		} else {
			break
		}
	}
	fmt.Printf("--------------------\n")
	var secondNumber int
	for {
		fmt.Print("Введите второе число: ")
		if _, err := fmt.Scan(&secondNumber); err != nil {
			fmt.Printf("Ошибка! Введите целое чило в диапозоне от %d до %d.\n", math.MinInt, math.MaxInt)
		} else {
			break
		}
	}
	fmt.Printf("--------------------\n")
	funcTest(func(a, b int) int { return a + b }, firstNumber, secondNumber, "+")
	funcTest(func(a, b int) int { return a - b }, firstNumber, secondNumber, "-")
	funcTest(func(a, b int) int { return a * b }, firstNumber, secondNumber, "*")
}

func funcTest(A func(int, int) int, firstNumber int, secondNumber int, oper string) {
	defer fmt.Println(A(firstNumber, secondNumber))

	fmt.Printf("Результат операции %d %s %d = ", firstNumber, oper, secondNumber)
}
