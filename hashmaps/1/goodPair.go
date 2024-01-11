package main 

import (
	"fmt"
)

// func numIdenticalPairs(nums []int) int {
// 	numCount := make(map[int]int)

// 	for _, num := range nums {
// 		numCount[num]++
// 	}

// 	goodPairs := 0

// 	for -, count := range numCount {
// 		goodPairs += count * (count - 1) / 2
// 	}

// 	return goodPairs
// }

// func numIdenticalPairs(arr []int) int {

// 	count := 0
	
// 	for index, value := range arr {
// 		for i := index + 1; i < len(arr); i++ {
// 			if value == arr[i] {
// 				count++
// 			}
// 		}
// 	}

// 	return count
// }

func numIdenticalPairs(arr []int) int {
	numCount := make(map[int]int)

	for _, value := range arr {
		numCount[value]++
	}

	goodPairs := 0

	for _, count := range numCount {
		goodPairs += count * (count - 1) / 2
	}

	retunr goodPairs
}

func main() {
    // Example usage
    nums := []int{1, 2, 3, 1, 1, 3, 2, 3}
    result := numIdenticalPairs(nums)
    fmt.Println("Number of good pairs:", result)
}