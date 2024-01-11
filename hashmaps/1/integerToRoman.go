package main

import "fmt"

// func romanToInt(s string) int {
//     // Create a map to store the values of each Roman numeral symbol
//     romanValues := map[byte]int{
//         'I': 1,
//         'V': 5,
//         'X': 10,
//         'L': 50,
//         'C': 100,
//         'D': 500,
//         'M': 1000,
//     }

//     total := 0
//     prevValue := 0

//     // Iterate through the Roman numeral string from right to left
//     for i := len(s) - 1; i >= 0; i-- {
//         currentValue := romanValues[s[i]]

//         fmt.Println("The integer representation of %s is: %d\n", total)
//         // If the value is less than the previous one, subtract it
//         if currentValue < prevValue {
//             total -= currentValue
//         } else {
//             total += currentValue
//         }

//         prevValue = currentValue
//     }

//     return total
// }
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

    integerCount := 0
    prevValue := ""

    for i := len(rom) - 1; i >= 0; i-- {

        currentvalue := rom[i]

        if romanValues[currentvalue] < romanValues[prevValue] {
            integerCount = integerCount - romanValues[currentvalue]
        } else {
            integerCount = romanValues[currentvalue]
        }

        prevValue = currentvalue

    }

        return integerCount
}

func main() {
    // Example usage
    romanNumeral := "CMXXIV"
    result := romanToInt(romanNumeral)
    fmt.Printf("The integer representation of %s is: %d\n", romanNumeral, result)
}
