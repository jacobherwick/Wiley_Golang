package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().Unix())
	var playing bool = true
	var num int = rand.Int()
	var guess int
	fmt.Print("Guess what number I'm thinking of...\n")
	for playing{ // while playing is true
		fmt.Scanf("%d", &guess)
		switch {
		case guess == num:
			fmt.Println("Correct! Congratulations, you have beaten me.")
			playing = false
		case 
		}
		
	}
	
}