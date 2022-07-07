package main

import "fmt"

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var arr2 [10]int
	for pos, val := range arr {
		arr2[len(arr2)-pos-1] = val
	}
	fmt.Println(arr)
	fmt.Println(arr2)
}
