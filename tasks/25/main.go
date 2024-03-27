package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

func sleep(sec int) {
	timer := time.NewTimer(time.Duration(sec) * time.Second)
	<-timer.C
}

func main() {
	fmt.Println(1)
	sleep(5)
	fmt.Println(2)
}
