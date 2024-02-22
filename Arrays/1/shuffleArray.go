package main

import (
	"fmt"
)

func shuffleArray(nums []int, n int) []int {

	new := make([]int, len(nums))

	i := 0;
	for k := 0; k < len(nums); k = k + 2 {

		j := k + 1

		new[k] = nums[i]
		new[j] = nums[i + n]

		// fmt.Println(new)
		// fmt.Println(j)
		// fmt.Println(i)

		i++
	}

	return new

}

func main() {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8}

	n := 4

	result := shuffleArray(input, n)

	fmt.Println(result)

}