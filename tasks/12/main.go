package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(generateSet(slice))
}

func generateSet(s []string) []string {
	var result []string
	m := make(map[string]struct{})
	for _, i := range s {
		if _, ok := m[i]; !ok {
			result = append(result, i)
			m[i] = struct{}{}
		}
	}
	return result
}
