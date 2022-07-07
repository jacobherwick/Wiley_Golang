package main

/*
Outputs the reverse of an inputted integer
*/
import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Print("Please enter an integer: ")
	fmt.Scanf("%d", &num)

	var str string = strconv.Itoa(num)

	for i := len(str) - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])
	}
}
