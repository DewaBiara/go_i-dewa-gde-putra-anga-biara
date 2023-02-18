package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf(" fizz")
		} else if i%5 == 0 {
			fmt.Printf(" buzz")
		} else {
			fmt.Printf(" %d", i)
		}
	}
}
