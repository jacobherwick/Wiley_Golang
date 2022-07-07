package main

/*
Activity 7
Author: Jacob Herwick
Prompts the user for 5 numbers then calculates the sum, product, and average of all 5 numbers.
*/

import (
	"fmt"
)

func main() {
	var num1, num2, num3, num4, num5 float64
	fmt.Print("Enter 5 numbers: Number 1: ")
	fmt.Scanf("%f", &num1)
	fmt.Print("\nNumber 2: ")
	fmt.Scanf("%f", &num2)
	fmt.Print("\nNumber 3: ")
	fmt.Scanf("%f", &num3)
	fmt.Print("\nNumber 4: ")
	fmt.Scanf("%f", &num4)
	fmt.Print("\nNumber 5: ")
	fmt.Scanf("%f", &num5)

	var sum float64 = num1 + num2 + num3 + num4 + num5
	var prdct float64 = num1 * num2 * num3 * num4 * num5
	var avg float64 = sum / 5.00
	fmt.Printf("You entered the values: %f, %f, %f, %f, %f\n", num1, num2, num3, num4, num5)
	fmt.Printf("Sum of the numbers: %f\n", sum)
	fmt.Printf("Product of the numbers: %f\n", prdct)
	fmt.Printf("Average of the numbers: %f\n", avg)

}
