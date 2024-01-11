

// take input as array of ints

// map letters of last item to an array

// multiplay function to cross all items of array one with array two

// loop through array from last item, if theres another item then
// 	map letters of new item and call multiplay function	
// 		result = multiply function = last array

package main 

import "fmt"

func letterCombination(digits string) int {

	// letterMapping := map[int][]string{

    //     2: []string{"a", "b", "c"},
    //     3: []string{"d", "e", "f"},
    //     4: []string{"g", "h", "i"},
    //     5: []string{"j", "k", "l"},
    //     6: []string{"m", "n", "o"},
    //     7: []string{"p", "q", "r", "s"},
    //     8: []string{"t", "u", "v"},
	// 		9: []string{"w", "x", "y", "z"},
	// 		1: []string{" "},

    // }

	// var i int = len(digits) - 1
	use := digits[2]

	// letterMapping[use]

	fmt.Printf("%c", use)

	return 0
}


func main() {

	inputArray := "59"

	result := letterCombination(inputArray)

	fmt.Println(result)

}