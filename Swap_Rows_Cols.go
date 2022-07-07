package main

import "fmt"

/*
Swaps the top and bottom rows of the 2D array/slice
*/
func SwapRows(mat [][]int) {
	var temp []int = mat[0]
	mat[0] = mat[len(mat)-1]
	mat[len(mat)-1] = temp
}

func SwapCols(mat [][]int) {
	for _, val := range mat {
		var temp int = val[0]
		val[0] = val[len(val)-1]
		val[len(val)-1] = temp
	}
}

func main() {
	var mat [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	SwapRows(mat)
	SwapCols(mat)
	fmt.Print(mat)
}
