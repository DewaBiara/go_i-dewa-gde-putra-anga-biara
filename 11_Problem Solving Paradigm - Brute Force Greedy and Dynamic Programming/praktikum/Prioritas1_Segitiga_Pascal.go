package main

import "fmt"

func generatePascalTriangle(numRows int) [][]int {
	pascalTriangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		pascalTriangle[i] = make([]int, i+1)
		pascalTriangle[i][0], pascalTriangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			pascalTriangle[i][j] = pascalTriangle[i-1][j-1] + pascalTriangle[i-1][j]
		}
	}
	return pascalTriangle
}

func main() {
	numRows := 5
	pascalTriangle := generatePascalTriangle(numRows)
	for _, row := range pascalTriangle {
		fmt.Println(row)
	}
}
