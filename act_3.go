package main

/*
Activity 3
Author: Jacob Herwick
Calculates compound interest
*/

import (
	"fmt"
	"math"
)

func main() {
	var deposit float64
	var interest float64
	var yearRep float64
	var yearQuant float64
	fmt.Print("Enter values in the following format: \"P, r, n, t\"\n")
	fmt.Scanf("%f, %f, %f, %f", &deposit, &interest, &yearRep, &yearQuant)
	fmt.Printf("%f", (deposit)*(math.Pow(1+(interest/yearRep), yearRep*yearQuant)))

}
