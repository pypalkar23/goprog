package main

import "fmt"

func main() {
	//goroutine call
	go f(0)
	//need to an arrangement like this otherwise the goroutine wont get enough time to execute
	//main would execute first and whole program stops
	var exitChar string
	fmt.Scanln(&exitChar)
}

func f(n int) {
	for i := 0; i <= 10; i++ {
		fmt.Println(n, "->", i)
	}
}
