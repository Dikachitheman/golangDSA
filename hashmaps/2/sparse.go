package main 

import (
	"fmt"
)

func DotProduct(vecA, vecB map[int]int) int {

	result := 0

	for indexA, valueA := range vecA {
		if valueB, exists := vecB[indexA]; exists {
			result = result + (valueB * valueA)
		}
	}
}

func SparseToMap(arr []int) map[int]int {

	result := make(map[int]int)
	
	for index, value := range arr {
		if value != 0 {
			result[index] = value
		}
	}

	return result
}

func main() {
	arrayE1 := []int{3,0,0,0,0,7,2,0,0,0}
	arrayE2 := []int{0,0,0,0,0,2,5,6,0,0}

	e1 := SparseToMap(arrayE1)
	e2 := SparseToMap(arrayE2)

	vc := DotProduct(e1, e2)

	fmt.Println(vc)
}