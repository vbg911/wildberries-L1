package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	for _, i := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Printf("квадрат числа %d = %d \n", num, num*num)
		}(i)
	}
	wg.Wait()
}
