package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	fmt.Println("Killed via Ctrl+C")
}
