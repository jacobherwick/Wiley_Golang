package main

import "fmt"

func main() {
	var flag bool = true
	for x := 3; x < 100; x++ {
		for b := 2; b <= (x / 2); b++ {
			if x%b == 0 {
				flag = false
			}
		}
		if flag {
			fmt.Println(x)
		}
		flag = true
	}

}
