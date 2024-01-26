package main

import (
	"fmt"
)

func findMatrix(nums []int) [4][4]int {

	makeMap := make(map[int]int)
	largestNum := 0

	for _, value := range nums {

		makeMap[value]++

		if makeMap[value] > largestNum {

			largestNum = makeMap[value]
			
		}

	}

	arrayArray := [4][4]int{}

	i := 0

	for index, value := range makeMap {

		// for i := 0; i < largestNum; i++ {

			for j := 0; j < value; j++ {
	
				arrayArray[j][i] = index
				
				// if j == value - 1 {
				// 	break
				// }

				fmt.Println(arrayArray)
	
			}

			i = i + 1
	
		// }

	}

	fmt.Println(makeMap)

	return arrayArray

}

func main() {

	input := []int{1, 3, 4, 1, 2, 3, 1, 1}

	result := findMatrix(input)

	fmt.Println(result)
}