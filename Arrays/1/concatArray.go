package main

import (
	"fmt"
)

func concatArray(nums []int) []int {

	n := 2 * len(nums)
	nh := n / 2
	ans := make([]int, n)

	for i, _ := range nums {

		ans[i] = nums[i]
		ans[i + nh] = nums[i]

	}

	return ans
}

func main() {

	input := []int{1, 3, 4, 1, 2}

	result := concatArray(input)

	fmt.Println(result)
}