package main

import (
	"fmt"
	"math"
	"math/rand"
)

func bitwise(num1, num2 int) int {

	op := []string{"-", "+"}
	result := 0
	opVal := op[rand.Intn(1)]
	// fmt.Println(opVal)

	if string(opVal) == "+" {
		result = num1 + num2
	} else if string(opVal) == "-" {
		result = num1 - num2
	}

	result = int(math.Abs(float64(result)))

	// result = int(result)

	return result
}

func findArray(pref []int) []int {

	newArray := make([]int, len(pref))

	newArray[0] = pref[0]
	
	for i := 1; i < len(pref); i++ {
		
		newArray[i] = bitwise(pref[i], pref[i - 1])

	}

	fmt.Println(newArray)

	return newArray

}

func main() {
	input := []int{4, 6, 8, 0, 4, 1}
	findArray(input)
	findArray(input)
}