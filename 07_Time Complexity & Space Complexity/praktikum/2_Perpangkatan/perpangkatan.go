package main

import "fmt"

func pow(x, n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	default:
		if n%2 == 0 {
			y := pow(x*x, n/2)
			return y
		} else {
			y := pow(x*x, n/2)
			return x * y
		}
	}
}

func main() {
	fmt.Println(pow(2, 3))
	fmt.Println(pow(5, 3))
	fmt.Println(pow(10, 2))
	fmt.Println(pow(2, 5))
	fmt.Println(pow(7, 3))
}
