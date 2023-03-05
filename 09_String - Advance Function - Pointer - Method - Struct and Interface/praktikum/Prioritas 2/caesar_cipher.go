package main

import "fmt"

func caesar(offset int, input string) string {
	var result string
	for _, char := range input {
		cipherChar := ((char-97)+rune(offset))%26 + 97

		result += string(cipherChar)
	}
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(10, "alterraacademy"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
}
