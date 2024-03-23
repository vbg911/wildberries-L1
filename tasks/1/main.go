package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type human struct {
	name string
	age  int
}

func (h human) sayHello() {
	fmt.Print(h.name, " говорит привет")
}

type action struct {
	actionName string
	human
}

func main() {
	a := action{
		actionName: "сказать привет",
		human: human{
			name: "Валерий",
			age:  21,
		},
	}

	a.sayHello()
}
