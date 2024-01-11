// take string and take array of word

// len = take length of each word 

// num = number of words in array

// loop through string, len at a time
// each cycle
// 	take len * num of letters	
// 		check if each len of letters exists in the array of words
// 			check if it exists in checked array, if not add.
// 			do it num times

// 		check if each len of letters exists in the copy of array of words
// 		called copy
// 			if it does take index of copy	
// 				copy new = copy old - index\
// 					from 0 to len of copy old	
// 						if i = index, 
// 							i = index + 1
// 						else i = index
	
		
// 		if all the words exist 
// 			take index of first word and move to next 

package main

import (
	"fmt"
)

func copyArray(index int, arr []string) []string {

	other := make([]string, len(arr))

	for i := 0; i < len(other) ; i++ {

		if i+1 == len(arr){
			// fmt.Println(i)
			// fmt.Println(len(arr))
            break
        } else if i >= index{
            other[i] = arr[i + 1]
        } else {
            other[i] = arr[i]
        }
	}

	return other
}

func sliceWord(index int, length int, s string) string {

	other := make([]string, length)

	for i := 0; i < len(other) ; i++, index++{
		other[i] = s[index]
	}


}

func findSubstring(s string, words []string) []int {

	length := len(words[0])
	// num := len(words)
	snum := len(s)

	// copyWords := words

	// other := copyArray( 1, copyWords)

	// fmt.Println(length)

	for i := 0; i < snum ; i+length {
		sliced = sliceWord(i, length, )
	}

	return []int{0}
}

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar","foo","the"}

	findSubstring(s, words)

	s = "wordgoodgoodgoodbestword"
	words = []string{"word","good","best","word"}

	findSubstring(s, words)

	s = "barfoothefoobarman"
	words = []string{"foo","bar"}

	findSubstring(s, words)
}