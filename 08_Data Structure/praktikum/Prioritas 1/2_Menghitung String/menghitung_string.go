package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	hasil := make(map[string]int)

	for _, n := range slice {
		hasil[n]++
	}

	return hasil
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}
