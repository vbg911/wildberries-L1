package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// интерфейс структуры адаптера
type Target interface {
	Operation()
}

// адаптируемая структура
type Adaptee struct {
}

// Метод адаптируемой структуры, который нужно вызвать где-то
func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("I am AdaptedOperation()")
}

// структура адаптера
type Adapter struct {
	*Adaptee
}

// реализация метода интерфейса, реализующего вызов адаптируемой структуры
func (adapter *Adapter) Operation() {
	adapter.AdaptedOperation()
}

// конструктор адаптера
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func main() {
	fmt.Println("Adapter demo:")
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
