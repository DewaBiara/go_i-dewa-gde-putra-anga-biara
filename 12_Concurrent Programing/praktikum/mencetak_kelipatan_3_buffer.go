package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		for i := 1; ; i++ {
			if i%3 == 0 {
				ch <- i
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Printf("Kelipatan 3 ke-%d: %d\n", i, <-ch)
	}
}
