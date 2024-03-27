package main

import (
	"fmt"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

// binarySearch выполняет бинарный поиск в отсортированном массиве arr
// возвращает индекс элемента target, если он найден, или -1, если не найден
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 13

	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Число %d найдено в позиции %d\n", target, index)
	} else {
		fmt.Printf("Число %d не найдено в массиве\n", target)
	}
}
