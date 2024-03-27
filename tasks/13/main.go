package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной.
*/

func main() {
	a, b := 10, 20
	fmt.Printf("a=%v, b=%v\n", a, b)
	a, b = b, a
	fmt.Printf("a=%v, b=%v", a, b)
}
