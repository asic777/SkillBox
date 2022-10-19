// Задание 3. Именованные возвращаемые значения
// Что нужно сделать
// Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его.
// Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.

package main

import "fmt"

func main() {
	fmt.Println("Задание 14.6.3. Преобразование числа с помощью функций с именованными возвращаемыми значениями.")
	fmt.Println("--------------------")

	fmt.Print("Введите число для преобразования: ")
	var firstNumber int
	fmt.Scan(&firstNumber)

	fmt.Printf("Результат после сложения: %v\n", sum100(firstNumber))
	fmt.Printf("Результат после умножения: %v\n", product100(firstNumber))
}

func sum100(number int) (result int) {
	result = number + 100
	return
}

func product100(number int) (result int) {
	result = number * 100
	return
}
