package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func main() {
	var (
		rw sync.RWMutex
		wg sync.WaitGroup
	)

	data := make(map[int]int)
	n := 5
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rw.Lock()
			data[rand.IntN(100)] = rand.IntN(100)
			rw.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(data)
}
