package main

import (
	"fmt"
	"strings"
)

func palindrome(word string) bool {
	word = strings.ToLower(word)
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var word string
	fmt.Print("Apakah palindrome?\n")
	fmt.Print("Masukkan kata  : ")
	fmt.Scanf("%s", &word)
	fmt.Printf("Captured : %s\n", word)
	fmt.Print("Output : ")

	if palindrome(word) {
		fmt.Print("Palindrome")
	} else {
		fmt.Print("Bukan Palindrome")
	}
}
