package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sumCh := make(chan int)

	for _, i := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sumCh <- num * num
		}(i)
	}

	go func() {
		wg.Wait()
		close(sumCh) // Закрытие канала после завершения отправки всех значений
	}()

	var sum int
	for i := 0; i < len(nums); i++ {
		sum += <-sumCh
	}
	fmt.Print(sum)
}
