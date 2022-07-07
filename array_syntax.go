package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 143
	arr2 := [4]string{"hi", "how", "are", "you"}

	for idx, val := range arr2 {
		fmt.Println(idx, "	", val)
	}

	fmt.Println(arr2[0:2]) // subarray

	matrix := [3][3]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}

}
