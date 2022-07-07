package main

import "fmt"

func main() {
	i := []interface{}{1, 2, 3}
	j := []interface{}{"ABC", "CDE", "EFG"}

	k := append(i, j)
	fmt.Println(k)
}
