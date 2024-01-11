package main

import (
	"fmt"
)

func numJewelIsInStones(jewels string, stones string) int {
	result := make(map[byte]int)


	for i := 0; i < len(jewels); i++ {
		for j := 0; j < len(stones); j++ {
			if jewels[i] == stones[j] {
				result[jewels[i]]++
			}
		}
	} 

	add := 0
	for _, value := range result {
		add += value
	}

	return add
}

func main() {
	jewels := "aA"
	stones := "aAAbbb"

	add := numJewelIsInStones(jewels, stones)

	fmt.Println(add)

	jewels_a := "h"
	stones_a := "aAAbbddbnnbbaAn"

	add_a := numJewelIsInStones(jewels_a, stones_a)

	fmt.Println(add_a)
}