package main

import (
	"fmt"
)

func generateMultiplesOfThree(out chan<- int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			out <- i
		}
	}
}

func printNumbers(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch := make(chan int)

	go generateMultiplesOfThree(ch)
	printNumbers(ch)
}
