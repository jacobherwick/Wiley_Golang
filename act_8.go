package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	var num int
	fmt.Print("Please enter the value of x followed by the length of the series (n) separated by a comma (no space): ")
	fmt.Scanf("%d,%d", &x, &num)
	var ans float64 = 0.00

	for i := 0; i <= num; i++ {
		var fact int = 1

		for j := 1; j <= i; j++ {
			fact *= j
		}
		ans += ((math.Pow(float64(x), float64(i))) / float64(fact))
	}

	fmt.Printf("%.3f\n", ans)
}
