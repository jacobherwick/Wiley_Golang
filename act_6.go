package main

/*
Activity 6
Author: Jacob Herwick
Calculates compound interest
*/

import (
	"fmt"
	"math"
)

func main() {
	var userNum float64
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &userNum)
	fmt.Printf("The boolean of your number is %t\n", userNum > 0.00)
	fmt.Printf("The binary equivalent of your number is %b\n", math.Float64bits(userNum))
	fmt.Printf("The square root of your number is %.4f\n", math.Sqrt(userNum))

}
