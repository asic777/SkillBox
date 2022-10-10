// Задание 2. Минимальный тип данных
// Что нужно сделать
// Достаточно часто при передаче по Сети или сохранении больших объёмов данных приходится выбирать тип с минимальным размером памяти, чтобы экономить трафик или место на диске. Напишите программу, в которую пользователь вводит два числа (int16), а программа выводит, в какой минимальный тип данных можно сохранить результат умножения этих чисел.

// Советы и рекомендации
// Обратите внимание, что положительный результат можно сохранить в меньшем типе за счёт использования uint8, uint16. Чтобы не возникло проблем с переполнением в процессе умножения, числа считывайте в int16, а перед умножением переведите их в int32.

// Проверить на примерах:
// 1 1 результат uint8
// 1 −1 результат int8
// 640 100 результат uint16
// −640 100 результат int32
// 3000 3000 результат uint32
// −3000 3000 результат int32

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		firstNumber   int
		secondNumber  int
		productNumber int32
	)

	for {
		fmt.Print("Введите первое число: ")
		fmt.Scan(&firstNumber)
		if firstNumber > math.MaxInt16 {
			fmt.Println("Введите чило в диапозоне от -32768 до 32767")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите второе число: ")
		fmt.Scan(&secondNumber)
		if secondNumber > math.MaxInt16 {
			fmt.Println("Введите чило в диапозоне от -32768 до 32767")
		} else {
			break
		}
	}

	productNumber = int32(firstNumber) * int32(secondNumber)

	if productNumber >= 0 {
		if productNumber <= math.MaxUint8 {
			fmt.Println("Минимальный тип данных: uint8")
		} else if productNumber <= math.MaxUint16 {
			fmt.Println("Минимальный тип данных: uint16")
		} else {
			fmt.Println("Минимальный тип данных: uint32")
		}
	} else {
		if productNumber >= math.MaxInt8 {
			fmt.Println("Минимальный тип данных: int8")
		} else if productNumber >= math.MaxInt16 {
			fmt.Println("Минимальный тип данных: int16")
		} else {
			fmt.Println("Минимальный тип данных: int32")
		}
	}

}
