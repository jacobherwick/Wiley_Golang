package main

import (
	"fmt"
	"sort"
)

type node struct {
	Value interface{}
}

// key can be any type
func main() {
	//key   value
	var myGreeting map[string]string
	fmt.Println(myGreeting)

	// this will not work, the map is not initialized
	// myGreeting["A"] = "ABC"
	fmt.Println(myGreeting == nil) // true

	myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Arno"] = "Bonjour."
	myGreeting["Mr. Owl"] = "Good evening."
	myGreeting["Henry"] = "Hello."
	myGreeting["Billy"] = "Bonjour."

	//this will be in alphabetical order
	fmt.Println(myGreeting)

	// this will not be in alphabetical order
	for i, j := range myGreeting {
		fmt.Println(i, j) // key, value
	}

	delete(myGreeting, "Billy") // removes the entry with the "Billy" key

	temp := map[string]int{"GHI": 5, "DEF": 6, "ABC": 7, "ZXX": 8}

	// let's sort it

	// slice of strings (the keys)
	keys := make([]string, 0, len(temp))

	for i := range temp {
		keys = append(keys, i) // add the keys to the slice, keys.
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, j := range keys {
		fmt.Println(j, temp[j])
	}

	//create a map where values must be a structure type

	structMap := make(map[int]node)

	structMap[0] = node{Value: "I am a node."}
	structMap[10] = node{Value: 15}
	// in no particular order
	for i, j := range structMap {
		fmt.Println(i, " : ", j.Value)
	}

}
