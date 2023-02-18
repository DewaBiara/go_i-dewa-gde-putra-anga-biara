package main

import "fmt"

func main() {
	var a, b, t float64

	fmt.Printf("Masukan Alas a : ")
	fmt.Scanf("%f\n", &a)

	fmt.Printf("Masukan Alas b : ")
	fmt.Scanf("%f\n", &b)

	fmt.Printf("Masukan Tinggi Trapesium : ")
	fmt.Scanf("%f\n", &t)

	luas := 0.5 * (a + b) * t
	fmt.Printf("Luas Trapesium    : %2.f", luas)
}
