package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	// Define the Roman numeral symbols and their corresponding values
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result strings.Builder

	for num > 0 {
		for index, value := range values {
			if num >= value {
				result.WriteString(symbols[index])
				num -= value
				break
			}
		}
	}

	return result.String()
}

func main() {
	// Example usage
	number := 27
	romanNumeral := intToRoman(number)
	fmt.Printf("The Roman numeral for %d is %s\n", number, romanNumeral)
}