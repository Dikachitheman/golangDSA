package main

import (
	"fmt"
)

// func twoSum(nums []int, target int) []int {

// 	first := 0
// 	second := 0
    
// 	for index, value := range nums {
// 		diff := target - value

// 		for i := index + 1; i < len(nums); i++ {
// 			if diff == nums[i] {
// 				first = index
// 				second = i

// 				break
// 			}

		
// 		}
// 	}

// 	arr := make([]int, 2)

// 	arr[0] = first
// 	arr[1] = second

// 	fmt.Println(arr)

// 	return arr
// }

func twoSum(nums []int, target int) []int {
	
	mp := make(map[int]int)

	for index, value := range nums {
		diff := target - value

		for i := index + 1; i < len(nums); i++ {
			if diff == nums[i] {
				mp[1] = index
				mp[2] = i
			}
		}
	}

	arr := make([]int, 2)

	arr[0] = mp[1]
	arr[1] = mp[2]

	fmt.Println(arr)
	
	return arr
}

func main() {

	nums := []int{2,7,11,15}
	target := 17
	// Output: [0,1]

	twoSum(nums, target)
}