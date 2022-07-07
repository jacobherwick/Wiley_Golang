package main

import "fmt"

func increment(arr []int) {
	for pos, _ := range arr {
		var b *int = &arr[pos]
		*b += 1
	}
}

func changeByVal(x int) {
	x = x + 10
	// if you print the address of this x, it will be different
	// than the address of what was passed in. It's a copy
}

func changeByAddress(x *int) {
	*x = *x + 10
}

// func increment2(arr *[]int) {
// 	for pos, _ := range *arr {
// 		&arr[pos] += 1
// 	}
// }

type contactInfo struct {
	Email   string
	ZipCode int
}

type Person1 struct {
	FirstName string
	LastName  string
	contactInfo
}

func (p Person1) print() {
	fmt.Println(p)
	fmt.Printf("%+v \n", p)
}

func (p Person1) updateFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person1) updateFirstNameV2(firstName string) {
	(*p).FirstName = firstName
}

func main() {

	jill := Person1{
		FirstName: "Jill",
		LastName:  "Jackson",
		contactInfo: contactInfo{
			Email:   "Jill.Jackson@Sample.com",
			ZipCode: 12345,
		},
	}

	//reference to jill
	jillPointer := &jill

	// This won't work, because the function does not pass by reference
	//jill.updateFirstName("Jill2")

	jillPointer.updateFirstNameV2("Jill2")

	jill.print()

	var a int = 20
	var b *int

	b = &a // b stores the memory address of variable a

	a = 10

	fmt.Println(b) // print address of b

	fmt.Println(*b)

	*b = 30 // also changes a

	//d := 10

	//f := 0

	//var c *int = &d

	//var e *int = &f

	arr := []int{0, 10, 20}

	increment(arr) // should be [31, 11, 1]

	// for _, val := range arr {
	// 	fmt.Printf("%d, ", *val)
	// }

	fmt.Println(arr)

	fmt.Println(a)

	//create a go lang function which takes an int array as
	// parameter and change the value of the array variable to
	// increase by 1

	//pass int array into it, every element should increase by 1

	x := 10
	px := &x
	ppx := &px // pointer, ppx, to another pointer, px

	*px += 2
	**ppx += 10
	fmt.Println(**ppx)
	fmt.Println(*px)
	fmt.Println(x)

}
