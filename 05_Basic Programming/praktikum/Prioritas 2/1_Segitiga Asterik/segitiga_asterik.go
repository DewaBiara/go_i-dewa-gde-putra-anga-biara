package main

import "fmt"

func segitiga_asterik(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	var x int
	fmt.Print("Input  : ")
	fmt.Scanf("%d", &x)
	fmt.Println("Output : ")

	segitiga_asterik(x)
}
