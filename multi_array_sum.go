package main

import "fmt"

func main() {
	var sum int
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	var rowQuant = len(matrix)
	var colQuant = len(matrix[0])
	for i := 0; i < rowQuant; i++ {
		for j := 0; j < colQuant; j++ {
			sum += matrix[i][j]
		}
		fmt.Printf("Sum of row %d is %d\n", i, sum)
		sum = 0
	}
	for k := 0; k < colQuant; k++ {
		for l := 0; l < rowQuant; l++ {
			sum += matrix[l][k]
		}
		fmt.Printf("Sum of column %d is %d\n", k, sum)
		sum = 0
	}
}
