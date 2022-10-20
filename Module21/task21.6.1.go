// Задание 1. Расчёт по формуле
// Что нужно сделать
// Напишите функцию, производящую следующие вычисления.
// S = 2 × x + y ^ 2 − 3/z, где x — int16, y — uint8, a z — float32.
// Тип S должен быть во float32.

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Задание 20.6.1. Расчёт по формуле с приведением типов.\n")
	fmt.Printf("--------------------\n")

	var numberX int16
	for {
		fmt.Print("Введите x: ")
		if _, err := fmt.Scan(&numberX); err != nil {
			fmt.Printf("Ошибка! Введите целое чило в диапозоне от %d до %d.\n", math.MinInt16, math.MaxInt16)
		} else {
			break
		}
	}
	var numberY uint8
	for {
		fmt.Print("Введите y: ")
		if _, err := fmt.Scan(&numberY); err != nil {
			fmt.Printf("Ошибка! Введите целое чило в диапозоне от %d до %d.\n", 0, math.MaxUint8)
		} else {
			break
		}
	}
	var numberZ float32
	for {
		fmt.Print("Введите z: ")
		if _, err := fmt.Scan(&numberZ); err != nil {
			fmt.Printf("Ошибка! Введите целое чило в диапозоне от %v до %v.\n", -math.MaxFloat32, math.MaxFloat32)
		} else {
			break
		}
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("Расчет по формуле S = 2*x + y^2 - 3/z = %f", calcFunc(numberX, numberY, numberZ))
}

func calcFunc(x int16, y uint8, z float32) float32 {
	return 2*float32(x) + float32(math.Pow(float64(y), 2)) - (3 / float32(z))
}
