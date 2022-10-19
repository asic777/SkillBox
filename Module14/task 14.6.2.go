// Задание 2. Функция, возвращающая несколько значений
// Что нужно сделать
// Напишите программу, которая с помощью функции генерирует две случайные точки в двумерном пространстве (две координаты),
// а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	randomRangeX = 10
	randomRangeY = 10
)

func main() {
	fmt.Println("Задание 13.6.1. Вычисление суммы чётных чисел заданного диапазона.")
	fmt.Println("--------------------")

	// fmt.Print("Введите первое число: ")
	// var firstNumber int
	// fmt.Scan(&firstNumber)
	// fmt.Print("Введите второе число: ")
	// var secondNumber int
	// fmt.Scan(&secondNumber)

	rand.Seed(time.Now().UnixNano())

	x, y := GeneratingTwoCoordinates()
	fmt.Printf("Первая случайная координата, расчитанная по формуле: x = %v, y = %v\n", x, y)
	x, y = GeneratingTwoCoordinates()
	fmt.Printf("Вторая случайная координата, расчитанная по формуле: x = %v, y = %v\n", x, y)
}

func GeneratingTwoCoordinates() (int, int) {
	return calcCoordinates(rand.Intn(randomRangeX), rand.Intn(randomRangeY))
}

func calcCoordinates(x, y int) (int, int) {
	return 2*x + 10, -3*y - 5
}
