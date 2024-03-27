package main

import (
	"fmt"
	"sync"
)

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	in := make(chan int)
	Out := make(chan int)

	go func() {
		for n := range in {
			Out <- 2 * n
		}
		close(Out)
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range Out {
			fmt.Println(num)
		}
	}()

	for i := range arr {
		in <- i
	}
	close(in)
	wg.Wait()
}
