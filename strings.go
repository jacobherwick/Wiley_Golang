package main

import (
	"fmt"
	"strings"
)

func main() {
	var numWords int
	var tempWord string
	var sum int = 0
	fmt.Print("This program will find the average length of the \n" +
		"words inputted, and tell you which words had above-average\n" +
		"length. How many words would you like to input? ")

	fmt.Scanf("%d", &numWords)
	fmt.Scanf(" %c")
	slice1 := make([]string, numWords)
	fmt.Println("You may begin entering words one at a time.")
	for i := 0; i < numWords; i++ {
		fmt.Scanf("%s", &tempWord)
		fmt.Scanf(" %c")
		slice1 = append(slice1, strings.ToLower(tempWord))
		sum += len(tempWord)
	}
	var avg int = sum / numWords
	slice2 := make([]string, 0)
	slice3 := make([]string, 0)
	for _, val := range slice1 {
		if len(val) > avg {
			slice2 = append(slice2, val)
		} else if len(val) < avg {
			slice3 = append(slice3, val)
		}
	}

	fmt.Println("Original words: ")
	for j := 0; j < numWords; j++ {
		fmt.Println(slice1[j])
	}
	fmt.Println("-----	-----	-----	-----	-----")
	fmt.Printf("Average length: %d\n", avg)
	fmt.Println("-----	-----	-----	-----	-----")

	fmt.Println("Above-average length words: ")
	for _, val := range slice2 {
		fmt.Println(val)
	}
	fmt.Println("-----	-----	-----	-----	-----")
	fmt.Println("Below-average length words: ")
	for _, val := range slice3 {
		fmt.Println(val)
	}
}
