package main

import (
	"fmt"
)

func twoSum(arr []int, num int) []int {

	ans := make(map[int]int)

	for index, value := range arr {
		diff := num - value 

		for i := index + 1; i < len(arr); i++ {
			if arr[i] == diff {
				ans[0] = index
				ans[1] = i
			}
		}
	}

	ar := make([]int, 2)

	ar[0] = ans[0]
	ar[1] = ans[1]

	return ar
}

func main() {

	nums := []int{2,7,11,15}
	target := 17
	// Output: [0,1]

	twoSum(nums, target)
}