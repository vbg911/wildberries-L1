package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func main() {
	timer := time.NewTimer(5 * time.Second)
	result := make(chan int)

	go func(c chan int) {
		for {
			c <- rand.Int()
		}
	}(result)

	for {
		select {
		case <-timer.C:
			fmt.Println("программа завершена после 5 секунд")
			return
		default:
			fmt.Println(<-result)
		}
	}
}
