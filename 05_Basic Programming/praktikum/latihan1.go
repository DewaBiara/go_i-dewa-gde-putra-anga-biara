package main

import "fmt"

func main() {
	var hasil int8

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			hasil += 1
		}
	}

	fmt.Printf("Hasil : %d", hasil)
}
