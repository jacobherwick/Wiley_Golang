package main

import "fmt"

type class struct {
	Classname string
	students  []student // user defined data type
}

type student struct {
	Name   string
	RollNo int
	City   string
}

// pops the element at index x
func pop(arr []int, x int) (int, []int) {
	n := arr[x]
	//append everything after the index to everything before the index
	return n, append(arr[:x], arr[x+1:]...)[:len(arr)-1]
}

//receives a slice, sends out a slice
func reverse(arr []int) []int {
	var nArray []int
	for i := len(arr) - 1; i >= 0; i-- {
		nArray = append(nArray, arr[i])
	}
}

func main() {
	abs := student{Name: "Ross", RollNo: 30, City: "New York"}
	cbs := student{Name: "Mary", RollNo: 31, City: "London"}
	students := []student{abs, cbs, student{Name: "Jack", RollNo: 32, City: "London"}}
	students = append(students, student{Name: "Kate", RollNo: 33, City: "New York"})
	fmt.Println(students)

	class := class{"FirstA", students}
	fmt.Println(class)

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a, arr := pop(arr, 4)
	fmt.Println(a, arr)
}
