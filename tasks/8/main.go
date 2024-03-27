package main

import "fmt"

/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	num := int64(3) // 011 -> 3
	fmt.Println(num)
	num2, _ := changeBit(num, 0, 0) //010 -> 2
	fmt.Println(num2)
}

func changeBit(a int64, i int, bit int) (int64, error) {
	if i < 0 || i > 63 {
		return 0, fmt.Errorf("bit number must be from 0 to 63")
	}
	if bit == 1 {
		return a | (1 << i), nil
	} else if bit == 0 {
		return a & (^(1 << i)), nil
	} else {
		return 0, fmt.Errorf("bit must be 0 or 1")
	}
}
