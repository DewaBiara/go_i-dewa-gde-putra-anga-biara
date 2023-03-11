package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	roman := convertToRoman(number)

	fmt.Printf("Roman numeral conversion: %s\n", roman)
}

func convertToRoman(number int) string {
	var roman string

	for number > 0 {
		switch {
		case number >= 1000:
			roman += "M"
			number -= 1000
		case number >= 900:
			roman += "CM"
			number -= 900
		case number >= 500:
			roman += "D"
			number -= 500
		case number >= 400:
			roman += "CD"
			number -= 400
		case number >= 100:
			roman += "C"
			number -= 100
		case number >= 90:
			roman += "XC"
			number -= 90
		case number >= 50:
			roman += "L"
			number -= 50
		case number >= 40:
			roman += "XL"
			number -= 40
		case number >= 10:
			roman += "X"
			number -= 10
		case number >= 9:
			roman += "IX"
			number -= 9
		case number >= 5:
			roman += "V"
			number -= 5
		case number >= 4:
			roman += "IV"
			number -= 4
		default:
			roman += "I"
			number--
		}
	}

	return roman
}
