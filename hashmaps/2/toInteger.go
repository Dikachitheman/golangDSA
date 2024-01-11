package main

import (
	"fmt"
)

func romanToInt(rom string) int {

    romanValues := map[byte]int{
                'I': 1,
                'V': 5,
                'X': 10,
                'L': 50,
                'C': 100,
                'D': 500,
                'M': 1000,
            }

	result := 0

	currentValue := 0
	prevValue := 0

	for i := len(rom) - 1; i >= 0; i-- {

		currentValue = romanValues[rom[i]]

		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}

func main() {
    // Example usage
    romanNumeral := "MMDXCIII"
    result := romanToInt(romanNumeral)
    fmt.Printf("The integer representation of %s is: %d\n", romanNumeral, result)
}