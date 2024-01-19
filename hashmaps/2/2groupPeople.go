package main  

import (
	"fmt"
)

func groupPeople(arr []int) {
	
	makeMap := make(map[int][]int)
	count := make(map[int]int)
	howMany := make(map[int]int)
	// i := 0

	for _, value := range arr {

		howMany[value]++
	}

	for _, value := range arr {

		makeMap[value] = make([]int, howMany[value])
		count[value] = 0
		
	}


	for index, value := range arr {

		makeMap[value][count[value]] = index
		count[value]++

		// fmt.Println(makeMap[value])
	}

	fmt.Println(makeMap)

	number := 0

	for index, value := range makeMap {

		number = (len(value) / index) + number

	}

	i := 0
	arrayArray := make([][]int, number)

	for index, value := range makeMap {

		indexA := 0
		indexB := index

		if len(value) == index {

			arrayArray[i] = value
			i++

		} else {

			for x := 0; x < len(value) / index; x++ {

				// fmt.Println("a", indexA)
				// fmt.Println("b", indexB)
				// fmt.Println("i", index)

				arrayArray[i] = value[indexA:indexB]
				indexA += index
				indexB = indexB + index
				i++

			}

		}

	}

	fmt.Println(arrayArray)

}

func main () {
	groupSizes := []int{3,3,3,3,3,1,3,2,2,3,3,3,2,2}
	groupPeople(groupSizes)

	groupSizes2 := []int{2,1,3,3,3,2}
	groupPeople(groupSizes2)

	groupSizes3 := []int{3,3,1,3}
	groupPeople(groupSizes3)
}