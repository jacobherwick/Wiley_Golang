package main

/*
Activity 4
Author: Jacob Herwick
Calculates the amount of interest earned
*/

import (
	"fmt"
)

func main() {
	var principal float64
	var interest float64
	var numDays float64
	fmt.Print("Enter values in the following format: \"[principal], [interest rate as a decimal (ex. 0.15 for 15%)], [number of days] \"\n")
	fmt.Scanf("%f, %f, %f", &principal, &interest, &numDays)
	fmt.Printf("%f", principal*interest*(numDays/365))

}
