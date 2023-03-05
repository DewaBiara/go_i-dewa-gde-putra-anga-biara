package main

import (
	"fmt"
	"strings"
)

const (
	cipherTextAlphabet = "zytcgmhjrfplkoqnsixvduewba"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var result string

	lowerCase := strings.ToLower(s.name)
	for _, char := range lowerCase {
		cipherChar := cipherTextAlphabet[char-97]
		result += string(cipherChar)
	}
	return result
}

func (s *student) Decode() string {
	var result string

	lowerCase := strings.ToLower(s.nameEncode)
	for _, char := range lowerCase {
		cipherChar := strings.Index(cipherTextAlphabet, string(char)) + 97
		result += string(rune(cipherChar))
	}
	return result
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Studet's Name : ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of Student's Name ", a.name, " is ", c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Studet's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of Student's Name ", a.nameEncode, " is ", c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}
