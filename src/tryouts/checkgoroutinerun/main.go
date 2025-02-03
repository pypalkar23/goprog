package main

import (
	"fmt"
	"time"
)

func main() {
	go parallelroutine()

	time.Sleep(10 * time.Second)
}

func parallelroutine() {
	i := 0
	for range time.Tick(2 * time.Second) {
		fmt.Println(i)
		i++
	}
}
