// Задание 1. Слияние отсортированных массивов
// Что нужно сделать
// Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.

package main

import "fmt"

func main() {
	firstArray := [4]int{1, 3, 88, 89}
	secondArray := [5]int{13, 14, 15, 16, 17}
	fmt.Println("Первый сортированный массив:", firstArray)
	fmt.Println("Второй сортированный массив:", secondArray)
	fmt.Println("Массив после слияния:", mergeArrays(firstArray, secondArray))
}

func mergeArrays(firstArray [4]int, secondArray [5]int) []int {
	sizeFirstArray, sizeSecondArray := len(firstArray), len(secondArray)
	sizeMergeArray := sizeFirstArray + sizeSecondArray
	mergeArray := make([]int, sizeMergeArray)

	f, s := 0, 0 //index for first and second Arrays
	for i := 0; i < sizeMergeArray; i++ {
		if f >= sizeFirstArray {
			mergeArray[i] = secondArray[s]
			s++
		} else if s >= sizeSecondArray {
			mergeArray[i] = firstArray[f]
			f++
		} else if firstArray[f] < secondArray[s] {
			mergeArray[i] = firstArray[f]
			f++
		} else {
			mergeArray[i] = secondArray[s]
			s++
		}
	}
	return mergeArray
}
