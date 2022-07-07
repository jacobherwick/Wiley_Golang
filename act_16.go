package main

/*
Outputs the length of an inputted string
*/
import (
	"fmt"
)

func main() {
	var str string
	var max int = 0
	fmt.Print("Please enter a string: ")
	fmt.Scanln(&str)

	for pos, _ := range str {
		if pos > max {
			max = pos + 1
		}
	}
	fmt.Printf("The length of this string is %d.\n", max)
}
