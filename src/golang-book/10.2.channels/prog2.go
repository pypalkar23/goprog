package main

import (
	"fmt"
	"time"
)

//in this function one can only send data to the channel.attempt to receive from
//it will result in compilation error
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

//in this function one can only receive data from the channel.attempt to receive from
//it will result in compilation error
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var inp string
	fmt.Scanln(&inp)
}
