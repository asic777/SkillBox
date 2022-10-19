// Задание 4. Область видимости переменных
// Что нужно сделать
// Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
// Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
// Затем вызовите по очереди три функции, передавая результат из одной в другую.

package main

import "fmt"

var (
	a = 1
	b = 2
	c = 3
)

func main() {
	fmt.Println("Задание 14.6.4. Последовательный вызов функций, использующих глобальные переменные.")
	fmt.Println("--------------------")

	fmt.Print("Введите число для преобразования: ")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Printf("Результат после первой функции: %v\n", funcAB(firstNumber))
	fmt.Printf("Результат после второй функции: %v\n", funcBC(firstNumber))
	fmt.Printf("Результат после третьей функции: %v\n", funcAC(firstNumber))

	fmt.Printf("Результат последовательного вызова трех функций: %v\n", funcAB(funcBC(funcAC(firstNumber))))
}

func funcAB(number int) int {
	return number + a + b
}

func funcBC(number int) int {
	return number + b + c
}

func funcAC(number int) int {
	return number + a + c
}
