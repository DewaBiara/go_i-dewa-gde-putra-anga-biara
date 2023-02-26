package main

import "fmt"

func sumDiagonal(matrix [][]int) int {
	var d1, d2 int
	var diagonalKiri []int

	var diagonalKanan []int
	for i := 0; i < len(matrix); i++ {
		diagonalKanan = append(diagonalKanan, matrix[i][i])
		d1 = d1 + matrix[i][i]
	}

	for i := 0; i < len(matrix); i++ {
		diagonalKiri = append(diagonalKiri, matrix[i][len(matrix)-1-i])
		d2 = d2 + matrix[i][len(matrix)-1-i]
	}

	fmt.Printf("Diagonal Kanan : %d\n", diagonalKanan)
	fmt.Printf("Diagonal Kiri : %d\n", diagonalKiri)

	hasil := d1 - d2

	if hasil < 0 {
		hasil = hasil * -1
	}

	return hasil
}

func main() {
	fmt.Println(sumDiagonal([][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}
