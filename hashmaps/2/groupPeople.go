package main 

import (
	"fmt"
)

func groupPeople(arr []int) {

	makeMap := make(map[int][]int)
	count := make(map[int]int)
	// i := 0

	for _, value := range arr {

		makeMap[value] = make([]int, value)
		count[value] = 0
		
	}

	for index, value := range arr {

		makeMap[value][count[value]] = index
		count[value]++

		fmt.Println(makeMap[value])
	}

	fmt.Println(makeMap)

	// return
}

func main() {
	
	// groupSizes := []int{3,3,3,3,3,1,3}
	// output := groupPeople(groupSizes)

	// groupSizes2 := []int{2,1,3,3,3,2}
	// output2 := groupPeople(groupSizes2)

	groupSizes := []int{3,3,1,3}
	groupPeople(groupSizes)
}