package main

import (
	"fmt"
)

func afterOperator(nums []string) int {

	increment := 0

	for _, v := range nums {

		if string(v[1]) == "+" {

			increment = increment + 1

		} else if string(v[1]) == "-"{

			increment = increment - 1

		}

	}

	return increment

}

func main() {

	input := []string{"--X", "++X", "++X", "++X", "++X"}

	result := afterOperator(input)

	fmt.Println(result)

}