package main

import "fmt"

// a new type, Operator, which is a function that
// takes a float64 as a parameter and a return value
type Operator func(x float64) float64

func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x) // perform operation on all entries
	}
	return res
}

// func A(t interface{}) {
// 	fmt.Println("Calling A", t)
// }

// func B(t interface{}) {
// 	fmt.Println("Calling B", t)
// }

// func C(t interface{}) {
// 	fmt.Println("Calling C", t)
// }

func A() {
	fmt.Println("Calling A")
}

func B(t int) {
	fmt.Println("Calling B", t)
}

func C(A string, f float32) {
	fmt.Println("Calling C", A, " ", f)
}

func main() {
	// array of functions
	// var fns []func(interface{})
	// fns = append(fns, A) // make it a slice and add function A
	// fns = append(fns, B)
	// fns = append(fns, C)

	// fns[0](12)
	// fns[1]("Hello")
	// fns[2](122.55)

	// for _, f := range fns {
	// 	f() // call the function
	// }

	f := []interface{}{A, B, C}
	f[0].(func())()
	f[1].(func(int))(15)
	f[2].(func(string, float32))("Hello", 15.15)

}
