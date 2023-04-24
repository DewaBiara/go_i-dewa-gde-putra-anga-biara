package main

import (
	"fmt"
	"sync"
)

func main() {
	var sum int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			sum += 1
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Total: %d\n", sum)
}
