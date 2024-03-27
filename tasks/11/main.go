package main

import (
	"fmt"
	"slices"
)

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	set1 := []int{2, 4, 5, 6, 8, 99}
	set2 := []int{2, 6, 99, 7}

	result := getIntersections(set2, set1)
	fmt.Println(result)
}

func getIntersections(s1 []int, s2 []int) []int {
	var result []int
	for _, i := range s1 {
		if slices.Contains(s2, i) {
			result = append(result, i)
		}
	}
	return result
}
