package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	rwm   sync.RWMutex
	count int
}

func (c *Counter) Increment() {
	c.rwm.Lock()
	defer c.rwm.Unlock()
	c.count++
}

func (c *Counter) GetValue() int {
	c.rwm.Lock()
	defer c.rwm.Unlock()
	return c.count
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	N := 1000
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Counter value: %d\n", counter.GetValue())
}
