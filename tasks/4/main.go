package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(ctx context.Context, c chan int, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine ", i, "done")
			return
		default:
			fmt.Println("goroutine ", i, "read", <-c)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan int)
	n := 5 // количество воркеров
	for i := 0; i < n; i++ {
		go worker(ctx, c, i)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-interrupt:
			fmt.Print("ctrl-c pressed")
			cancel()
			close(c)
			return
		default:
			c <- rand.Int()
		}
	}
}
