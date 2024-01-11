package main 

import (
	"fmt"
)

// func numIdenticalPairs(arr []int) int {
// 	result := 0

// 	for index, value := range arr {
// 		for i := index+1; i < len(arr); i++ {
// 			if value == arr[i] {
// 				result++
// 			}
// 		}
// 	}

// 	return result
// }

func numIdenticalPairs(arr []int) int {
	mp := make(map[int]int)

	for _, value := range arr {
		mp[value]++
	}

	ans := 0

	for _, value := range mp {
		ans += value * (value - 1) / 2
	}

	return ans
}

func main() {
    // Example usage
    nums := []int{1, 2, 3, 1, 1, 3, 2, 3}
    result := numIdenticalPairs(nums)
    fmt.Println("Number of good pairs:", result)
}