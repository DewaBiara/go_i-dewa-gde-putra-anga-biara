package main

import (
	"fmt"
	"strconv"
)

func getBinaryString(num int) string {
	if num == 0 {
		return "0"
	}

	binaryString := ""
	for num > 0 {
		rem := num % 2
		binaryString = strconv.Itoa(rem) + binaryString
		num = num / 2
	}

	return binaryString
}

func binaryRepresentation(n int) []string {
	result := make([]string, n+1)

	for i := 0; i <= n; i++ {
		result[i] = getBinaryString(i)
	}

	return result
}

func main() {
	fmt.Println(binaryRepresentation(2))
	fmt.Println(binaryRepresentation(3))
}
