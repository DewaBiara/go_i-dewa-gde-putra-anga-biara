package main

import "fmt"

func fibo(n int) int {
	if n < 2 {
		return n
	}
	return (n - 1) + (n - 2)
}

func main() {
	fmt.Println(fibo(10))
}
