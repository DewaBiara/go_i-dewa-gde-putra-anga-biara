package main

import "fmt"

func main() {
	var x int

	fmt.Printf(" Masukan sebuah bilangan : ")
	fmt.Scanf("%d\n", &x)

	if x%2 == 0 {
		fmt.Printf("Bilangan %d adalah genap\n", x)
	} else {
		fmt.Printf("Bilangan %d adalah ganjil\n", x)
	}
}
