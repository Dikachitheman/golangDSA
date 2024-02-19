package main

import (
	"fmt"
	"sort"
)

func sortF(a []int) []int {

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for _, v := range a {
		fmt.Println(v)
	}
	return a
}

func miniProductSum(arrFirst []int, arrSecond []int) int {

	result := 0
	end := len(arrSecond) - 1

	sortF(arrFirst)
	sortF(arrSecond)
	
	for i, _ := range arrFirst {

		result += arrFirst[i] * arrSecond[end - i]

	}

	return result
}

func main() {

	input := []int{2,1,4,5,7}
	input2 := []int{3,2,4,8,6}

	result := miniProductSum(input, input2)

	fmt.Println(result)

}