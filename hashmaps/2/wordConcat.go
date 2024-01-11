package main

import (
	"fmt"
	"reflect"
)

func findSubstring(s string, words []string) []string {

	eachWord := len(words[0])

	wordsLength := len(words)

	sliceArray := make([]string, wordsLength)

	wordsMap := make(map[string]int)
	sliceMap := make(map[string]int)

	j := 0

	//  := len(words) * eachWord

	for i:= 0; i < eachWord * wordsLength; i += eachWord {
		e := i + eachWord
		sliceArray[j] = s[i:e]
		j++
	}

	for _, value := range words {
		sliceMap[value]++
	}

	for _, valueS := range sliceArray {
		for _, valueW := range words {
			if valueS == valueW {
				wordsMap[valueS]++
				break
			}
		}
	}

	res := reflect.DeepEqual(wordsMap, sliceMap)

	fmt.Println("slice array-- ",sliceArray)
	fmt.Println("words-- ",words)
	fmt.Println("wordsmap-- ",wordsMap)
	fmt.Println("slicemap-- ",sliceMap)
	fmt.Println("equal?-- ",res)
	fmt.Println("\n")
	
	return sliceArray
}

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar","foo","the"}

	findSubstring(s, words)

	s = "goodwordgoodbestwordgoodgood"
	words = []string{"word","good","best","word"}

	findSubstring(s, words)

	s = "barfoothefoobarman"
	words = []string{"foo","bar"}

	findSubstring(s, words)
}