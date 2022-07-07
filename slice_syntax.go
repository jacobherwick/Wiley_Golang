package main

import "fmt"

/*
similar to array. They are an extension built ontop of the array
Unlike an array, they can grow at any time. They reference
underlying data. If you change the slice, you change the array.

*/
func main() {
	size := 10
	// make a slice of size 10
	a := make([]int, size, 30) // 30 is the max capacity
	// if the max capacity is reached, it will copied to
	// a new slice with double the capacity
	fmt.Println(a)
	a = append(a, 4) // add 4 to the end of the array, increasing the size by 1

	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{-1, -2, -3}

	//define a slice of y
	b := y[0:2]

	copy(x[3:], y)
	// copy into x from y, starting at index 3
	// [1  2  3  -1  -2  -3  7]
	fmt.Println(x)
	fmt.Println(b)

}
