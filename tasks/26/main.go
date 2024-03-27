package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func checkDuplicates(s string) bool {
	s = strings.ToLower(s)
	runes := []rune(s)
	m := make(map[rune]struct{})
	for _, r := range runes {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = struct{}{}
	}
	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	fmt.Printf("%s -> %v\n", s1, checkDuplicates(s1))
	fmt.Printf("%s -> %v\n", s2, checkDuplicates(s2))
	fmt.Printf("%s -> %v\n", s3, checkDuplicates(s3))
}
