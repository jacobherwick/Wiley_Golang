package main

/*
Prints the digits 1-7 and then removes a number from the end
each time the outter loop repeats. Then it repeats this with
capital letters A-D.
*/
import "fmt"

func main() {
	for i := 1; i <= 7; i++ {
		for j := 1; j <= 8-i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
	for i := 65; i <= 68; i++ {
		for j := 65; j <= i; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}

	name := "Robert Pike"
	// for pos, v := range name {
	// 	fmt.Println("Pos ", pos, "	", v, "	", string(v))
	// }
	var spacePos int
	for x := 0; x < len(name); x++ {
		if name[x] == ' ' {
			spacePos = x
		}
	}
	for x := spacePos - 1; x >= 0; x-- {
		fmt.Printf("%c", name[x])
	}
	fmt.Print(" ")
	for x := len(name) - 1; x > spacePos; x-- {
		fmt.Printf("%c", name[x])
	}

}
