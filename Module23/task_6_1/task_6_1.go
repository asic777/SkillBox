// Задание 1. Сортировка вставками
// Что нужно сделать
// Напишите функцию, сортирующую массив длины 10 вставками.

// Советы и рекомендации
// Алгоритм сортировки доступен на «Википедии».
// В качестве среды разработки используйте GoLand или VSCode.

package main

import (
// "fmt"
// "math"
// "math/rand"
// "sort"
// "sync"
// "time"
)

// глупая сортировка вставками
func InsertSort0(array []int) []int {
	length := len(array)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	return array
}

// лучше на упорядоченных массивах
func InsertSort1(array []int) []int {
	length := len(array)
	for i := 1; i < length; i++ {
		for j := i; j >= 1 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

// оптимальная сортировка на любых массивах невысокой сложности
func InsertSort2(array []int) []int {
	length := len(array)
	for i := 1; i < length; i++ {
		movItem := array[i]
		j := i
		for ; j >= 1 && array[j-1] > movItem; j-- {
			array[j] = array[j-1]
		}
		array[j] = movItem
	}
	return array
}

// лучше на неупорядоченных массивах
func InsertSort3(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; ; j-- {
			if j >= 1 && array[i] < array[j-1] {
				continue
			} else {
				movItem := array[i]
				array = append(array[:i], array[i+1:]...)
				array = append(array[:j+1], array[j:]...)
				array[j] = movItem
				break
			}
		}
	}
	return array
}

// лучше на неупорядоченных массивах
func InsertSort4(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; ; j-- {
			if j >= 1 && array[i] < array[j-1] {
				continue
			} else {
				movItem := array[i]
				copy(array[j+1:i+1], array[j:i])
				array[j] = movItem
				break
			}
		}
	}
	return array
}

// лучше на упорядоченных массивах
func BubbleSort1(array []int) []int {
	isSort := false
	length := len(array)
	for !isSort {
		isSort = true
		for i := 0; i < length-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSort = false
			}
		}
	}
	return array
}

// лучше на неупорядоченных массивах
func BubbleSort2(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

// func BubbleSortRecursion(array []int, size int) []int {
// 	if size == 1 {
// 		return array
// 	}
// 	for i := 0; i < size-1; i++ {
// 		if array[i] > array[i+1] {
// 			array[i], array[i+1] = array[i+1], array[i]
// 		}
// 	}
// 	BubbleSortRecursion(array, size-1)
// 	return array
// }

// func GenerateArray(lengthArray, numbersRange int) (resultArray []int) {
// 	resultArray = make([]int, lengthArray)
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < lengthArray; i++ {
// 		resultArray[i] = rand.Intn(numbersRange)
// 	}
// 	return
// }
