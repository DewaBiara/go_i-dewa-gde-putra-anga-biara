package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	temp := make(map[string]bool)

	for _, value := range arrayA {
		temp[value] = true
	}
	for _, value := range arrayB {
		temp[value] = true
	}

	var result []string
	for key := range temp {
		result = append(result, key)
	}

	return result
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))
}
