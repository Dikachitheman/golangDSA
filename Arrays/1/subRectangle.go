package main

import (
	"fmt"
)

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	for i:=row1; i <= row2; i++ {
		for j:=col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}

func main() {

	queries := []string{"getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"}
	value := [][]int{{0,2},{0,0,2,2,5},{0,2},{3,1},{3,1,3,1,4},{3,1},{0,2}}
	rectangle := [][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}}

	obj := Constructor(rectangle)

	fmt.Println(obj.rectangle)

	for i, v := range queries {

		if v == "updateSubrectangle" {
			obj.UpdateSubrectangle(value[i][0], value[i][1], value[i][2], value[i][3], value[i][4])
			fmt.Println(obj.rectangle)
		} else if v == "getValue" {
			getvalue := obj.GetValue(value[i][0], value[i][1])
			fmt.Println(getvalue)
		}

	}

	
}
