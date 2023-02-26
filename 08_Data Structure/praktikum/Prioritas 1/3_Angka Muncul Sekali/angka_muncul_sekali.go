package main

import (
	"fmt"
	"strconv"
)

func munculSekali(number string) []int {
	var result []int
	tmp := make(map[int]int)

	for _, i := range number {
		num, _ := strconv.Atoi(string(i))
		tmp[num]++
	}

	for key, value := range tmp {
		if value == 1 {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
