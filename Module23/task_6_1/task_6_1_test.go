package main

import (
	"math"
	"testing"
)

const (
	benchPow10 = 3
)

var (
	inputTestArray = [][]int{
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

	exemplarTestArray = [][]int{
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

	lengthArray        int
	exemplarArray      = []int{}
	mediumComplexArray = []int{}
	highComplexArray   = []int{}
)

func init() {
	lengthArray = int(math.Pow10(benchPow10))

	for i := 0; i < lengthArray; i++ {
		exemplarArray = append(exemplarArray, i)
	}

	for i := lengthArray - 1; i >= 0; i-- {
		highComplexArray = append(highComplexArray, i)
	}

	for i := 0; i < int(lengthArray)/2; i++ {
		mediumComplexArray = append(mediumComplexArray, i)
	}
	for i := lengthArray - 1; i >= lengthArray/2; i-- {
		mediumComplexArray = append(mediumComplexArray, i)
	}

}

func funcTest(f func([]int) []int, t *testing.T) {
	for k, input := range inputTestArray {
		output := make([]int, len(input))
		copy(output, input)
		output = f(output)
		for i, o := range output {
			if exemplarTestArray[k][i] != o {
				t.Error("Expected  ", exemplarTestArray[k][i], ", got", o)
			}
		}
		t.Log(input, output, exemplarTestArray[k])
	}
}

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

func BenchmarkInsertSort2(b *testing.B) {
	input := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	exemplar := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := make([]int, len(input))
	copy(output, input)
	output = InsertSort2(output)
	for i, o := range output {
		if exemplar[i] != o {
			b.Error("Expected  ", exemplar[i], ", got", o)
		}
	}
}
