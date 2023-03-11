package main

import (
	"fmt"
)

func fibonacciTopDown(n int, temp []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if temp[n] == 0 {
		temp[n] = fibonacciTopDown(n-1, temp) + fibonacciTopDown(n-2, temp)
	}
	return temp[n]
}

func Fibonacci(n int) int {
	temp := make([]int, n+1)
	return fibonacciTopDown(n, temp)
}

func main() {
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(3))

	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(6))
	fmt.Println(Fibonacci(7))

	fmt.Println(Fibonacci(9))
	fmt.Println(Fibonacci(10))
}
