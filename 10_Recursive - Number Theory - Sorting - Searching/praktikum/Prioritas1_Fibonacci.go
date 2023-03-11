package main

import "fmt"

func fibocci(number int) int {
	if number == 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	return fibocci(number-1) + fibocci(number-2)
}

func main() {
	fmt.Println(fibocci(0))  // 0
	fmt.Println(fibocci(2))  // 1
	fmt.Println(fibocci(9))  // 34
	fmt.Println(fibocci(10)) // 55
	fmt.Println(fibocci(12)) // 144
}
