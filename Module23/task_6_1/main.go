package main

import (
	//"Module23/task61"
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

const (
	countTest = 1
	pow       = 5
)

func main() {
	var wg sync.WaitGroup

	fmt.Printf("Задание 23.6.1. Сортировка вставками. Сравнение с другими алгоритмами сортировки.\n")

	array := [3 * pow][]int{}
	for i := 0; i < pow; i = i + 1 {
		maxJ := 10 * math.Pow10(i)
		for j := 0; j < int(maxJ); j++ {
			array[i] = append(array[i], j)
		}
	}
	for i := pow; i < 2*pow; i++ {
		maxJ := 10 * math.Pow10(i-pow)
		for j := int(maxJ) - 1; j >= 0; j-- {
			array[i] = append(array[i], j)
		}
	}
	for i := 2 * pow; i < 3*pow; i++ {
		maxJ := 10 * math.Pow10(i-2*pow)
		for j := 0; j < int(maxJ)/2; j++ {
			array[i] = append(array[i], j)
		}
		for j := int(maxJ) - 1; j >= int(maxJ)/2; j-- {
			array[i] = append(array[i], j)
		}
	}

	fmt.Printf("--------------------\n")
	fmt.Println("Упорядоченные массивы:")
	for i := 0; i < pow; i++ {
		fmt.Println("[", array[i][0], array[i][1], " ... ", array[i][len(array[i])-1], "]")
	}

	fmt.Printf("--------------------\n")
	fmt.Println("Абсолютно неупорядоченные массивы:")
	for i := pow; i < 2*pow; i++ {
		fmt.Println("[", array[i][0], array[i][1], " ... ", array[i][len(array[i])-1], "]")
	}

	fmt.Printf("--------------------\n")
	fmt.Println("Упорядоченные наполовину массивы:")
	for i := 2 * pow; i < 3*pow; i++ {
		fmt.Println("[", array[i][0], array[i][1], " ... ", array[i][len(array[i])/2-1], array[i][len(array[i])/2], " ... ", array[i][len(array[i])-2], array[i][len(array[i])-1], "]")
	}

	for t := 0; t < 3*pow; t++ {
		fmt.Printf("--------------------\n")
		fmt.Println("Выполняется сортировка массива: [", array[t][0], array[t][1], " ... ", array[t][len(array[t])-1], "]")
		//fmt.Printf("--------------------\n")
		wg.Add(1)
		go func(t int) {
			defer wg.Done()

			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				InsertSort0(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка insertSort0: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()

			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				InsertSort1(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка insertSort1: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				InsertSort2(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка insertSort2: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				InsertSort3(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка insertSort3: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				InsertSort4(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка insertSort4: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				BubbleSort1(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка bubbleSort1: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				BubbleSort2(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка bubbleSort2: ", duration)
		}(t)
		wg.Wait()
		// wg.Add(1)
		// go func(t int) {
		// 	defer wg.Done()
		// 	start := time.Now()
		// 	for i := 0; i < countTest; i++ {
		// 		arrayGo := make([]int, len(array[t]))
		// 		copy(arrayGo, array[t])
		// 		task61.BubbleSortRecursion(arrayGo, len(arrayGo))
		// 	}
		// 	duration := time.Since(start)
		// 	fmt.Println("Сортировка bubbleSortRecursion: ", duration)
		// }(t)
		// wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				sort.Slice(arrayGo, func(i, j int) bool {
					return arrayGo[i] < arrayGo[j]
				})
			}
			duration := time.Since(start)
			fmt.Println("Сортировка sort.Slice: ", duration)
		}(t)
		wg.Wait()
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			start := time.Now()
			for i := 0; i < countTest; i++ {
				arrayGo := make([]int, len(array[t]))
				copy(arrayGo, array[t])
				sort.Ints(arrayGo)
			}
			duration := time.Since(start)
			fmt.Println("Сортировка sort.Ints: ", duration)
		}(t)
		wg.Wait()
	}
	fmt.Printf("--------------------\n")
	fmt.Println("Программа завершена")
}
