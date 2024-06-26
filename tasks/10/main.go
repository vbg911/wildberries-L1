package main

import (
	"fmt"
)

/*
Дана последовательность температурных колебаний:
-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножествах не важна.

Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	res := make(map[int][]float32)
	for _, t := range temperatures {
		key := (int(t / 10)) * 10
		res[key] = append(res[key], t)
	}

	fmt.Println(res)
}
