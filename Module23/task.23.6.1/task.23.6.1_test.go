// go test -benchmem -bench . > bench_$(date +"%d-%m-%Y-%Hh%Mm").txt
package main

import (
	"math"
	"testing"
)

const (
	benchPow10 = 3 // размер массива 10 в степени benchPow10
)

var (
	// Входные данные для тестирования функций
	inputArrTest = [][]int{
		{},
		{-1},
		{0},
		{1},
		{-2, -1},
		{-1, -2},
		{-1, -1},
		{-1, 0},
		{0, -1},
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
		{2, 1},
		{1, 2},
		{-1, 1},
		{1, -1},
		{-1, -1, -1},
		{-1, -2, -3},
		{-3, -2, -1},
		{0, 0, 0},
		{1, 0, -1},
		{-1, 0, 1},
		{1, 1, 1},
		{1, 2, 3},
		{3, 2, 1},
		{-1, -2, -3, -4, -5},
		{2, 1, 0, -1, -2},
		{5, 4, 3, 2, 1},
		{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
		{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
		{4, 3, 2, 1, 0, 0, -1, -2, -3, -4},
		{-4, -3, -2, -1, 0, 0, 1, 2, 3, 4},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{23, 65, 234, 764, 1, 57, 86, 121, -84, -500},
		{-500, -84, 1, 23, 57, 65, 86, 121, 234, 764}}

	// Эталонный результат сортировки
	exemplarArrTest = [][]int{
		{},
		{-1},
		{0},
		{1},
		{-2, -1},
		{-2, -1},
		{-1, -1},
		{-1, 0},
		{-1, 0},
		{0, 0},
		{0, 1},
		{0, 1},
		{1, 1},
		{1, 2},
		{1, 2},
		{-1, 1},
		{-1, 1},
		{-1, -1, -1},
		{-3, -2, -1},
		{-3, -2, -1},
		{0, 0, 0},
		{-1, 0, 1},
		{-1, 0, 1},
		{1, 1, 1},
		{1, 2, 3},
		{1, 2, 3},
		{-5, -4, -3, -2, -1},
		{-2, -1, 0, 1, 2},
		{1, 2, 3, 4, 5},
		{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
		{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1},
		{-4, -3, -2, -1, 0, 0, 1, 2, 3, 4},
		{-4, -3, -2, -1, 0, 0, 1, 2, 3, 4},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
		{-500, -84, 1, 23, 57, 65, 86, 121, 234, 764},
		{-500, -84, 1, 23, 57, 65, 86, 121, 234, 764}}

	// Переменные для измерения производительности
	lenArrBench           int
	exemplarArrBench      = []int{} // отсортированный эталонный массив
	mediumComplexArrBench = []int{} // массив, наполовину отсортированный
	highComplexArrBench   = []int{} // массив, абсолютно несортированный
)

func init() {
	lenArrBench = int(math.Pow10(benchPow10))

	for i := 0; i < lenArrBench; i++ {
		exemplarArrBench = append(exemplarArrBench, i)
	}

	for i := lenArrBench - 1; i >= 0; i-- {
		highComplexArrBench = append(highComplexArrBench, i)
	}

	for i := 0; i < int(lenArrBench)/2; i++ {
		mediumComplexArrBench = append(mediumComplexArrBench, i)
	}
	for i := lenArrBench - 1; i >= lenArrBench/2; i-- {
		mediumComplexArrBench = append(mediumComplexArrBench, i)
	}
}

// Метод, принимающий функцию, которую тестируем
func funcTest(f func([]int) []int, t *testing.T) {
	for k, input := range inputArrTest {
		output := make([]int, len(input))
		copy(output, input)
		output = f(output)
		for i, o := range output {
			if exemplarArrTest[k][i] != o {
				t.Error("Expected  ", exemplarArrTest[k][i], ", got", o)
			}
		}
		t.Log(input, output, exemplarArrTest[k])
	}
}

// func funcTestRecursion(f func([]int, int), t *testing.T) {
// 	for k, input := range inputTestArray {
// 		output := make([]int, len(input))
// 		copy(output, input)
// 		f(output, len(input))
// 		for i, o := range output {
// 			if exemplarTestArray[k][i] != o {
// 				t.Error("Expected  ", exemplarTestArray[k][i], ", got", o)
// 			}
// 		}
// 		t.Log(input, output, exemplarTestArray[k])
// 	}
// }

func TestInsertSort0(t *testing.T) {
	funcTest(InsertSort0, t)
}

func TestInsertSort1(t *testing.T) {
	funcTest(InsertSort1, t)
}

func TestInsertSort2(t *testing.T) {
	funcTest(InsertSort2, t)
}

func TestInsertSort3(t *testing.T) {
	funcTest(InsertSort3, t)
}
func TestInsertSort4(t *testing.T) {
	funcTest(InsertSort4, t)
}

func TestBubbleSort1(t *testing.T) {
	funcTest(BubbleSort1, t)
}

func TestBubbleSort2(t *testing.T) {
	funcTest(BubbleSort2, t)
}

func TestBubbleSort3(t *testing.T) {
	funcTest(BubbleSort3, t)
}

func TestSortInts(t *testing.T) {
	funcTest(SortInts, t)
}

func TestSortSlice(t *testing.T) {
	funcTest(SortSlice, t)
}

// не работает с рекурсией
// func TestBubbleSortRecursion(t *testing.T) {
// 	funcTestRecursion(BubbleSortRecursion, t)
// }

// Метод, принимающий функцию, производительность которой измеряем
func funcBench(f func([]int) []int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		easyComplexArr := make([]int, len(exemplarArrBench))
		copy(easyComplexArr, exemplarArrBench)
		mediumComplexArr := make([]int, len(mediumComplexArrBench))
		copy(mediumComplexArr, mediumComplexArrBench)
		highComplexArr := make([]int, len(highComplexArrBench))
		copy(highComplexArr, highComplexArrBench)

		// Сортируем массивы разной сложности
		b.StartTimer()
		f(easyComplexArr)
		f(mediumComplexArr)
		f(highComplexArr)
	}
}

func BenchmarkInsertSort0(b *testing.B) {
	funcBench(InsertSort0, b)
}

func BenchmarkInsertSort1(b *testing.B) {
	funcBench(InsertSort1, b)
}

func BenchmarkInsertSort2(b *testing.B) {
	funcBench(InsertSort2, b)
}

func BenchmarkInsertSort3(b *testing.B) {
	funcBench(InsertSort3, b)
}

func BenchmarkInsertSort4(b *testing.B) {
	funcBench(InsertSort4, b)
}

func BenchmarkBubbleSort1(b *testing.B) {
	funcBench(BubbleSort1, b)
}

func BenchmarkBubbleSort2(b *testing.B) {
	funcBench(BubbleSort2, b)
}

func BenchmarkBubbleSort3(b *testing.B) {
	funcBench(BubbleSort3, b)
}

func BenchmarkSortInts(b *testing.B) {
	funcBench(SortInts, b)
}

func BenchmarkSortSlice(b *testing.B) {
	funcBench(SortSlice, b)
}
