package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter float: ")
	response, error := input.ReadString('.')
	if error != nil {
		fmt.Println("uh oh")
	}
	response = strings.TrimSuffix(response, ".")

	fmt.Printf("Float without decimals: "+"%s", response)

}
