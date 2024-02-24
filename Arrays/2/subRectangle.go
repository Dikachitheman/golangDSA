package main

import (
	"fmt"
)

type SubrectangleQueries struct {
    rectangle [][]int // store the original rectangle as a field
}

func Constructor(rectangle [][]int) SubrectangleQueries {
    return SubrectangleQueries{rectangle: rectangle} // return an instance of the class with the rectangle field initialized
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
    for i := row1; i <= row2; i++ { // loop through the rows of the subrectangle
        for j := col1; j <= col2; j++ { // loop through the columns of the subrectangle
            this.rectangle[i][j] = newValue // update the value at the current position
        }
    }
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
    return this.rectangle[row][col] // return the value at the given position
}

func main() {

	queries := []string{"SubrectangleQueries","getValue","updateSubrectangle","getValue","getValue","updateSubrectangle","getValue","getValue"}
	value := [][]int{{0,2},{0,0,3,2,5},{0,2},{3,1},{3,0,3,2,10},{3,1},{0,2}}
	matrix := [][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}}

	obj := Constructor(matrix) // create an object of the class with the given matrix
	for i := 0; i < len(queries); i++ { // loop through the queries
		switch queries[i] { // switch on the query type
		case "getValue":
			fmt.Println(obj.GetValue(value[i][0], value[i][1])) // call the GetValue method and print the result
		case "updateSubrectangle":
			obj.UpdateSubrectangle(value[i][0], value[i][1], value[i][2], value[i][3], value[i][4]) // call the UpdateSubrectangle method with the given parameters
		}
	}
}
