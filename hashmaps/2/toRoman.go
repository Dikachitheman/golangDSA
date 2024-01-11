package main

import (
	"fmt"
	"strings"
)

func intToRoman(number int) string {
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result strings.Builder
	index := 0
	j := 0

	for number > 0 {
		for i := index; i < len(values); i++ {
			if number >= values[i] {
				result.WriteString(symbols[i])
				number -= values[i]
				break
			}

			index = i

			j++
		}

		fmt.Println(j)
		j = 0
	}

	return result.String()
}

func main() {
	// Example usage
	number := 27
	romanNumeral := intToRoman(number)
	fmt.Printf("The Roman numeral for %d is %s\n", number, romanNumeral)
}