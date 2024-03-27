package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает
две числовых переменных a, b, значение которых > 2^20.
*/

func main() {
	var a big.Int
	var b big.Int
	var c big.Int

	a.SetString("2000000000000000000000", 10)
	b.SetString("1000000000000000000000", 10)
	action := "+"
	switch action {
	case "+":
		c.Add(&a, &b)
	case "-":
		c.Sub(&a, &b)
	case "*":
		c.Mul(&a, &b)
	case "/":
		c.Div(&a, &b)
	}

	fmt.Println(c.Text(10))
}
