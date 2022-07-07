package main

/*
Activity 5
Author: Jacob Herwick
Creates 5 values and displays 6 comparisons of those values. Three of
the comparisons return true, and three return false.
*/

import (
	"fmt"
)

func main() {
	a, b, c, d, e := 0, -10, 10, 5, 10

	fmt.Printf("a = %d\nb = %d\nc = %d\nd = %d\ne = %d\n", a, b, c, d, e)
	fmt.Printf("a < b = %t\n", a < b)
	fmt.Printf("a == b = %t\n", a == b)
	fmt.Printf("e != c = %t\n", e != c)
	fmt.Printf("d < c = %t\n", d < c)
	fmt.Printf("b != e = %t\n", b != e)
	fmt.Printf("a < d = %t\n", a < d)

}
