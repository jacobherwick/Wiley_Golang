package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 20
	var b int = 30
	fmt.Println("a = " + strconv.Itoa(a))
	fmt.Println("b = " + strconv.Itoa(b))

	if a > b {
		goto MESSAGE1 //this will jump the execution to where MESSAGE1 is defined
	} else if a == b {
		goto MESSAGE3
	} else {
		goto MESSAGE2
	}

MESSAGE1: // We define a label that we can use in a goto statement
	fmt.Println("a is greater than b")

MESSAGE2:
	fmt.Println("b is greater than a")

MESSAGE3:
	fmt.Println("be is equal to a")
}
