package main

import (
	"fmt"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{5, 82, 49, 15, 228, 116, 120, 36, 10, 20, 13, 39, 11, 10, 0}
	sortedArr := qsort(arr)
	fmt.Println(sortedArr)
}

func qsort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	pivotIndex := len(a) / 2
	pivot := a[pivotIndex]

	left := make([]int, 0)
	right := make([]int, 0)
	for i := range a {
		if i == pivotIndex {
			continue
		}
		if a[i] <= pivot {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}

	sortedLeft := qsort(left)
	sortedRight := qsort(right)

	sortedArr := append(sortedLeft, pivot)
	sortedArr = append(sortedArr, sortedRight...)

	return sortedArr
}
