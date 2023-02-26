package main

import "fmt"

func PairSum(arr []int, target int) []int {
	l := 0
	h := len(arr) - 1

	for l <= h {
		if arr[l]+arr[h] == target {
			return []int{l, h}
		} else if arr[l]+arr[h] < target {
			l++
		} else if arr[l]+arr[h] > target {
			h--
		}
	}

	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))
}
