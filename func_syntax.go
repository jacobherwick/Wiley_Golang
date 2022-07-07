package main

import "fmt"

// new data type
type Operator func(int, int) int

// func Map(op Operator, a []float64) []float64 {

// }

// function with generic type
func findDataType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float")
	case nil:
		fmt.Println("Nil type")
	default:
		fmt.Println("Default")
	}

}

func multi(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func print(f Operator, x, y int) {
	fmt.Println(f(x, y))
}

//function that returns another function (which returns an int)

func fibonacci() func() int {
	first, second := 0, 1 //this line only called once
	return func() int {
		temp := first
		first, second = second, first+second
		return temp
	}
}

func sumN(numbers ...int) int { // variable list, any number of parameters
	sum := 0
	for num := range numbers {
		sum += num
	}
	return sum
}
func init() {
	fmt.Println("I am called first.")
}

func end() {
	fmt.Println("I am called at the end.")
}

// func generatePrime()int{

// }
func main() {
	// fmt.Println("Hello world")
	// a, b, c := multi(1, 2, 3)
	// fmt.Println("", a, b, c)
	// //IN-LINE FUNCTION
	// var sum int = func(a, b, c int) int {
	// 	return a + b + c
	// }(3, 4, 5)
	// fmt.Println(" ", sum)

	// print(add, 10, 12)
	// print(multi, 12, 10)

	defer end() // call at the end of main or before exit
	// from any function

	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	findDataType(10)
	findDataType(10.30)
	findDataType("string")
	findDataType(true)
	var k interface{} // empty
	findDataType(k)
}
