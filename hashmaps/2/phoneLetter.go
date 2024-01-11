package main

import (
	"fmt"
	"strings"
)

func crossMultipy(firstArray, secondArray []string) []string {

	var result strings.Builder

	length := len(firstArray) * len(secondArray)
	resultArray := make([]string, length)

	i := 0

	for _, valueA := range firstArray {

		for _, valueB := range secondArray {

			result.WriteString(valueA)
			result.WriteString(valueB)

			s_result := result.String()
			resultArray[i] = s_result
			i++

			// fmt.Println("result.string", result.String())
			result.Reset()

			// fmt.Println("resultarray", resultArray)
		}

	}
	return resultArray
}

func letterCombination(digits string) []string {

	letterMapping := map[byte][]string{

        '2': {"a", "b", "c"},
        '3': {"d", "e", "f"},
        '4': {"g", "h", "i"},
        '5': {"j", "k", "l"},
        '6': {"m", "n", "o"},
        '7': {"p", "q", "r", "s"},
        '8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
		'1': {" "},
    }

	// firstArray := []string{" "}

	firstArray := letterMapping[digits[0]]

	for i := 1; i < len(digits); i++ {

		secondArray := letterMapping[digits[i]]

		firstArray = crossMultipy(firstArray, secondArray)
	}

	// fmt.Println("lettermapping", letterMapping[digits[1]])

	return firstArray
}

func main() {

	inputArray := "5"

	result := letterCombination(inputArray)

	fmt.Println("result", result)

}