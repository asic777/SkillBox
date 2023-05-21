// Задание 1. Сортировка вставками
// Что нужно сделать
// Напишите функцию, сортирующую массив длины 10 вставками.

// Советы и рекомендации
// Алгоритм сортировки доступен на «Википедии».
// В качестве среды разработки используйте GoLand или VSCode.

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lenArray      uint64 = 100000
	numRangeArray int    = 100000
)

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Задание 23.6.1. Сортировка вставками.\n")
	fmt.Printf("--------------------\n")

	array := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Printf("Несортированный массив длиной %v: %v\n", len(array), array)
	fmt.Printf("Отсортированный массив длиной %v: %v\n", len(array), InsertSort2(array[:]))
	fmt.Printf("--------------------\n")
	fmt.Println()

	fmt.Println("Самообучение. Сравнительный тест разных методов сортировки на сложных массивах.")
	fmt.Printf("--------------------\n")

	if lenArray <= 0 {
		fmt.Println("Ошибка! Длина массива должна быть больше 0.")
		return
	}
	slice := GenerateArray(lenArray, numRangeArray)

	if lenArray > 10 {
		fmt.Printf("Случайный массив длиной %v: [%v %v ... %v %v]\n", lenArray, slice[0], slice[1], slice[lenArray-2], slice[lenArray-1])
	} else {
		fmt.Printf("Случайный массив длиной %v: %v\n", lenArray, slice)
	}

	fmt.Printf("--------------------\n")
	fmt.Println("Выполняется параллельная сортировка разными методами:")

	wg.Add(1)
	go func() {
		defer wg.Done()

		sortArray := make([]int, lenArray)
		copy(sortArray, slice)

		start := time.Now()
		sortArray = InsertSort2(sortArray)
		duration := time.Since(start)

		formatArrayOutput(sortArray, "InsertSort", duration)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		sortArray := make([]int, lenArray)
		copy(sortArray, slice)

		start := time.Now()
		sortArray = BubbleSort2(sortArray)
		duration := time.Since(start)

		formatArrayOutput(sortArray, "BubbleSort", duration)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		sortArray := make([]int, lenArray)
		copy(sortArray, slice)

		start := time.Now()
		sortArray = SortInts(sortArray)
		duration := time.Since(start)

		formatArrayOutput(sortArray, "sort.Ints", duration)
	}()
	wg.Wait()

	fmt.Printf("--------------------\n")
	fmt.Println("Программа завершена")
}

func formatArrayOutput(array []int, nameSort string, duration time.Duration)  {
	if len(array) > 10 {
		fmt.Printf("Сортировка %s:\t[%v %v ... %v %v] - %v\n", nameSort, array[0], array[1], array[lenArray-2], array[lenArray-1], duration)
	} else {
		fmt.Printf("Сортировка %s:\t%v - %v\n", nameSort, array, duration)
	}
}