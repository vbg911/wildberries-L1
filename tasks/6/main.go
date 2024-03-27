package main

import (
	"context"
	"fmt"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

func workerWithCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Завершение работы")
			return
		}
	}
}

func workerWithChan(c chan struct{}) {
	for {
		select {
		case <-c:
			fmt.Println("Завершение работы")
			return
		}
	}
}

func workerWithTimer() {
	timer := time.NewTimer(2 * time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("Завершение работы")
			return
		}
	}
}

func main() {
	// завершение горутины из-за отмены контекста
	ctx, cancel := context.WithCancel(context.Background())
	go workerWithCtx(ctx)
	cancel()

	// завершение горутины из-за полученного сигнала в канале
	stopChan := make(chan struct{})
	go workerWithChan(stopChan)
	stopChan <- struct{}{}

	// завершение горутины из-за закрытия канала
	go workerWithChan(stopChan)
	close(stopChan)

	// завершение горутины из-за таймера через 2 секунды
	go workerWithTimer()

	//время всем горутинам завершить работу
	time.Sleep(3 * time.Second)
}
