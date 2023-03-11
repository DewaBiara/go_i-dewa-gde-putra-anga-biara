package main

import "fmt"

func Frog(harga []int) int {
	n := len(harga)
	// Inisialisasi array dp dengan nilai tak terhingga
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1e9
	}
	dp[0] = 0

	// Hitung nilai dp
	for i := 0; i < n-1; i++ {
		for j := i + 1; j <= i+2 && j < n; j++ {
			biaya := abs(harga[i] - harga[j])
			if dp[j] > dp[i]+biaya {
				dp[j] = dp[i] + biaya
			}
		}
	}

	// Kembalikan nilai dp untuk batu terakhir
	return dp[n-1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         //30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) //40
}
