package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x, y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(p2 Point) float64 {
	return math.Sqrt((p.x-p2.x)*(p.x-p2.x) + (p.y-p2.y)*(p.y-p2.y))
}

func main() {
	p1 := NewPoint(2, 1)
	p2 := NewPoint(5, 1)
	fmt.Println(p1.Distance(p2))
}
