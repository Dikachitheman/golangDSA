package main

import (
	"fmt"
)

func buildArray(nums []int) []int {

	// ans := [6]int{}

	ans := make([]int, len(nums))

	for i, v := range nums {

		ans[i] = nums[v] //nums[nums[i]]
		fmt.Println(ans[i])

	}

	return ans

}

func main() {

	input := []int{0, 2, 1, 5, 3, 4}

	result := buildArray(input)

	fmt.Println(result)

}