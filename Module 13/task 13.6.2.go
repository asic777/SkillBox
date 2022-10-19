// Задание 2. Функция, принимающая значение по ссылке
// Что нужно сделать
// Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.

// Рекомендация
// В методе main создайте и присвойте значения двум переменным типа int, выведите значения этих переменных в консоль
// до вызова функции и после.

package main

import "fmt"

func main() {
	fmt.Println("Задание 13.6.2. Функция принимает в качестве аргументов указатели на два типа int и меняет их значения местами.")
	fmt.Println("--------------------")

	fmt.Print("Введите первое число: ")
	var firstNumber int
	fmt.Scan(&firstNumber)
	fmt.Print("Введите второе число: ")
	var secondNumber int
	fmt.Scan(&secondNumber)

	fmt.Printf("Введенные числа: %d и %d\n", firstNumber, secondNumber)
	SwapInt(&firstNumber, &secondNumber)
	fmt.Printf("После выполнения функции SwapInt: %d и %d\n", firstNumber, secondNumber)

}

func SwapInt(x1, x2 *int) {
	// *x1 += *x2
	// *x2 = *x1 - *x2
	// *x1 -= *x2
	*x1, *x2 = *x2, *x1
}
