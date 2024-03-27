package main

import (
	"fmt"
	"slices"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
Символы могут быть unicode.
*/

func Reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

func main() {
	fmt.Printf("'%s' -> '%s'\n", "こんにちは世界", Reverse("こんにちは世界"))
	fmt.Printf("'%s' -> '%s'\n", "Hej världen", Reverse("Hej världen"))
}
