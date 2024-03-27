package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме
способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var a any = "Hello"
	var b any = 42
	var c any = true
	var d any = make(chan int)
	var e any = make(chan string)

	fmt.Printf("a: %s\n", reflect.TypeOf(a).Name())

	_, ok := b.(int)
	if ok {
		fmt.Printf("b: %s\n", "int")
	}

	fmt.Printf("c: %s\n", reflect.ValueOf(c).Kind().String())

	fmt.Printf("d: %T\n", d)

	switch e.(type) {
	case chan string:
		fmt.Printf("e: %s\n", "chan string")
	}
}
