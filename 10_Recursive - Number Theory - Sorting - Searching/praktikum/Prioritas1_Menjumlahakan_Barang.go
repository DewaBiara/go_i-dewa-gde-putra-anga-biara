package main

import "fmt"

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	var result []pair

	for i := range items {
		if !isExist(result, items[i]) {
			result = append(result, pair{items[i], 1})
		} else {
			for j := range result {
				if result[j].name == items[i] {
					result[j].count++
				}
			}
		}
	}

	return mergeSort(result)
}

func isExist(result []pair, item string) bool {
	for i := range result {
		if result[i].name == item {
			return true
		}
	}
	return false
}

func mergeSort(items []pair) []pair {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := mergeSort(items[:middle])
	right := mergeSort(items[middle:])

	return merge(left, right)
}

func merge(left, right []pair) []pair {
	var result []pair

	for len(left) > 0 && len(right) > 0 {
		if left[0].count <= right[0].count {
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
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
