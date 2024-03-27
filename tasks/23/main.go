package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := 5

	fmt.Println(append(slice[:index], slice[index+1:]...))
}
