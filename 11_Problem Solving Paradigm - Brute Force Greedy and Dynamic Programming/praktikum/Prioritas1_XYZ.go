package main

import "fmt"

func main() {
	var A, B, C int
	fmt.Print("Masukkan A, B, dan C: ")
	fmt.Scan(&A, &B, &C)

	for z := 1; z <= 1000; z++ {
		for x := 1; x <= 1000; x++ {
			y := A - x - z
			if x*x+y*y+z*z == C && x*y*z == B {
				fmt.Printf("x = %d, y = %d, z = %d\n", x, y, z)
				return
			}
		}
	}
	fmt.Println("Tidak ada solusi")
}
