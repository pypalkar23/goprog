package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func sender1(datachan chan<- string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Returning from sender 1")
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			datachan <- fmt.Sprintf("From sender 1 %d", rand.Int())
		}
	}
}

func sender2(datachan chan<- string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Returning from sender 2")
	ticker := time.NewTicker(3 * time.Second)

	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			datachan <- fmt.Sprintf("From sender 2 %d", rand.Int())
		}
	}
}

func receiver(datachan1 <-chan string, datachan2 <-chan string, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("Returning from receiver")
	for {
		select {
		case <-ctx.Done():
			return
		case s1 := <-datachan1:
			fmt.Println(s1)
		case s2 := <-datachan2:
			fmt.Println(s2)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
	ctx, cancel := context.WithCancel(context.Background())
	datachan1 := make(chan string)
	datachan2 := make(chan string)

	defer close(datachan1)
	defer close(datachan2)

	wg.Add(3)
	go sender1(datachan1, ctx, &wg)
	go sender2(datachan2, ctx, &wg)
	go receiver(datachan1, datachan2, ctx, &wg)

	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-c
	cancel()
	wg.Wait()
	fmt.Println("Shutting down completely")
}
