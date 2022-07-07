package main

import "fmt"

func main() {
	var num int
	fmt.Print("Please enter an integer: ")
	fmt.Scanf("%d", &num)
	var flag bool = false

	for i := 2; i < num; i++ { // iterate through num, starting at 2 because that's the first prime number
		for j := 2; j <= (i / 2); j++ { // for every factor of the number
			if i%j == 0 { // if the number is cleanly divisible

				goto PLACEHOLDER // next iteration of outer loop
			}
		}
		for k := 2; k <= ((num - i) / 2); k++ {
			if (num-i)%k == 0 { // if the number is cleanly divisible
				goto PLACEHOLDER // next iteration of outer loop
			}
		}
		flag = true // if the code gets here, both i and num - i were prime
	PLACEHOLDER:
	}
	if flag {
		fmt.Println("This number is the sum of two prime numbers")
	} else {
		fmt.Println("This number is NOT the sum of two prime numbers")
	}
}
