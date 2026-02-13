package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		for i := range 10 {
			ch1 <- i*2 + 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range 10 {
			ch2 <- i * 2
		}
	}()

	go func() {
		wg.Wait()
		cancel()
		close(ch1)
		close(ch2)
	}()

	for {
		select {
		case i := <-ch1:
			fmt.Printf("%s に %d が書き込まれました\n", "ch1", i)
		case i := <-ch2:
			fmt.Printf("%s に %d が書き込まれました\n", "ch2", i)
		case <-ctx.Done():
			return
		}
	}
}
