package main

import "fmt"

type Sample struct {
	count int
}

func main() {
	count := 1
	sample := &Sample{count: count}
	for i := 0; i < 100; i++ {
		count++
	}
	fmt.Println(sample.count, count)
}
