package main

import "fmt"

func makeRow(numEntries int, slice []int, rows [][]int) []int {
	var entry int
	for i := 0; i < numEntries; i++ {
		fmt.Scanf("%d", &entry)
		fmt.Scanf(" %c")
		slice = append(slice, entry)
	}
	return slice
}

func main() {
	var adding bool = true
	var input string
	var numEntries int
	//var entry int

	rows := make([][]int, 0) // slice of slices with a starting capacity of 3

	for adding {

		fmt.Print("Would you like to add a row to the matrix? Enter \"Yes\" if so: ")
		fmt.Scanf("%s", &input)
		fmt.Scanf(" %c")
		if input == "Yes" {
			fmt.Print("How many entries are in this row? ")
			fmt.Scanf("%d", &numEntries)
			fmt.Scanf(" %c")
			slice := make([]int, 0)
			fmt.Println("Please enter the entries one at a time:")
			slice = makeRow(numEntries, slice, rows)
			rows = append(rows, slice)

			// for i := 0; i < numEntries; i++ {
			// 	fmt.Scanf("%d", &entry)
			// 	fmt.Scanf(" %c")
			// 	slice = append(slice, entry)
			// }
			// rows = append(rows, slice)
		} else {
			adding = false
		}
	}
	for _, val := range rows {
		fmt.Println(val)
	}
	//fmt.Print(rows)

	// var mat [3][3]int

	// for pos, val := range mat {
	// 	for pos1, val1 := range val {
	// 		fmt.Printf("Please enter a number to add to the matrix at position (%d, %d)", pos, pos1)
	// 		fmt.Scanf("%d", &val1)
	// 		fmt.Scanf(" %c")
	// 		mat[pos][pos1] = val1
	// 		fmt.Println()
	// 	}
	// }
	// fmt.Print(mat)
}
