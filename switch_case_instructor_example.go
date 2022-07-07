package main

import "fmt"

func main() {
	var numUnits int = 220
	var cost int
	if numUnits <= 100 {
		cost = 3 * numUnits
	} else if numUnits <= 200 {
		cost = 500 + (numUnits-100)*7
	} else if numUnits <= 350 {
		cost = 500 + 700 + (numUnits-200)*10
	}
	fmt.Print(cost)
}
