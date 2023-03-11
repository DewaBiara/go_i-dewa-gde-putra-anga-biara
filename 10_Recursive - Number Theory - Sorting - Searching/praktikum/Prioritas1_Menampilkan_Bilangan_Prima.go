package main

import "fmt"

func primeX(number int) int {
	var prime int = 2
	var count int = 1
	for count < number {
		prime++
		if isPrime(prime) {
			count++
		}
	}
	return prime
}

func isPrime(number int) bool {
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
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
