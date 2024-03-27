package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	testStr := "snow dog sun"
	words := strings.Split(testStr, " ")
	slices.Reverse(words)
	reversed := strings.Join(words, " ")
	fmt.Printf("%v -> %v\n", testStr, reversed)
}
