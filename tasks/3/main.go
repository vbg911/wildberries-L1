package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sumCh := make(chan int, len(nums))

	for _, i := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sumCh <- num * num
		}(i)
	}

	wg.Wait()
	close(sumCh)

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += <-sumCh
	}
	fmt.Print(sum)
}
