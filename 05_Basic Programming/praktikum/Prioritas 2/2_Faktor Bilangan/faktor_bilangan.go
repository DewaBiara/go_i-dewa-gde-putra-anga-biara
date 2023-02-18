package main

import "fmt"

func main() {
	var numbers int

	fmt.Print("Masukkan bilangan  : ")
	fmt.Scanf("%d", &numbers)
	fmt.Println("Hasil : ")

	for i := 1; i <= numbers; i++ {
		if numbers%i == 0 {
			fmt.Println("", i)
		}
	}
}
