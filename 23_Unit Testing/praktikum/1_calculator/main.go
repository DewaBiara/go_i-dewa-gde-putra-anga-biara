package main

import (
	"fmt"

	"calculator/lib/calculate"
)

func main() {
	fmt.Println("Calculator")

	var a, b float64
	var c int
	fmt.Println("Enter 2 numbers (Space separated): ")
	fmt.Scanln(&a, &b)
	fmt.Println("Enter operation: (1 for addition, 2 for subtraction, 3 for division, 4 for multiplication)")
	fmt.Scanln(&c)
	switch c {
	case 1:
		fmt.Println("Result is", calculate.Addition(a, b))
	case 2:
		fmt.Println("Result is", calculate.Subtraction(a, b))
	case 3:
		fmt.Println("Result is", calculate.Division(a, b))
	case 4:
		fmt.Println("Result is", calculate.Multiplication(a, b))
	default:
		fmt.Println("Invalid input")
	}
}
