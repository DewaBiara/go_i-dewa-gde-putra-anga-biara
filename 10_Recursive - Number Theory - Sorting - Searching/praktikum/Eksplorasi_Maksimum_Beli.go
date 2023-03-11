package main

import "fmt"

func MaximumBuyProduct(money int, productPrice []int) {
	var max int

	productPrice = mergeSort(productPrice)

	for i := 0; i < len(productPrice); i++ {
		if money >= productPrice[i] {
			money -= productPrice[i]
			max++
		}
	}

	fmt.Println(max)
}

func mergeSort(productPrice []int) []int {
	if len(productPrice) < 2 {
		return productPrice
	}

	middle := len(productPrice) / 2
	left := mergeSort(productPrice[:middle])
	right := mergeSort(productPrice[middle:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}
