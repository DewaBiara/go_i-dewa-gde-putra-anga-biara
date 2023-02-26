package main

import "fmt"

func primeNumber(number int) bool {
	if number <= 3 {
		return number > 1
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(13))
	fmt.Println(primeNumber(17))
	fmt.Println(primeNumber(20))
	fmt.Println(primeNumber(35))
}
