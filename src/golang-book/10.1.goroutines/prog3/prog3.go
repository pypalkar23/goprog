package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		f(i)
	}
	var inp string
	fmt.Scanln(&inp)
}

func f(n int) {
	for i := 0; i <= 10; i++ {
		fmt.Println(n, "->", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt) // Putting go routine to sleep for random interval
	}
}
