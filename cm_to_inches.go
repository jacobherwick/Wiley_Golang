package main

/*
Convert cm to inches and feet
Author: Jacob Herwick
Prompts the user for a measurement in cm, converts that into inches, and then converts the inch value into ft'in" format.
*/

import (
	"fmt"
	"math"
)

func main() {
	var cmVal float64
	fmt.Print("Enter a measurement in centimeters: ")
	fmt.Scanf("%f", &cmVal)

	var inchVal float64 = cmVal / 2.54

	var feetNum int = int(math.Trunc(inchVal)) / 12
	var inchNum float64 = inchVal - (12 * float64(feetNum))

	fmt.Printf("%d'%.2f\"", feetNum, inchNum)

}
