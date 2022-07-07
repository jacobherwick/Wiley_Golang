package main

import "fmt"

type Node struct {
	Value int
}

type Employee struct {
	Name   string
	Empid  int
	Salary float32
}

func (e Employee) getSalary() float32 {
	return e.Salary
}

// convert struct to string (and I think it's a method?)
func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

type Stack struct {
	nodes []*Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Node) {
	// make sure you're at the end
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}

	s.count-- // point to next position
	return s.nodes[s.count]

}

type Queue struct {
	nodes []*Node
	count int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) QueueIn(n *Node) {
	q.nodes = append(q.nodes[:q.count], n) // add node to end
	q.count++
}

func (q *Queue) QueueOut() *Node {
	if q.count == 0 {
		return nil
	}
	//q.pos++ // ignore this index in the future
	var n *Node = &Node{q.nodes[0].Value}
	q.nodes = q.nodes[1:] // remove first index
	q.count--
	return n
}

type person struct {
	First string //these names are not publicly visible
	Last  string
	Age   int
}

//inheritance
type doubleZero struct {
	person
	First         string //this name is publicly visible
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double O Seven",
		LicenseToKill: true,
	}
	fmt.Println(p1.First, p1.person.First)

	e := Employee{Name: "Johan", Empid: 12, Salary: 50000.50}
	e1 := Employee{Name: "Peter", Empid: 12, Salary: 40000.50}
	e2 := Employee{Name: "Rozer", Empid: 12, Salary: 70000.50}
	fmt.Println(e.getSalary())
	fmt.Println(e1.getSalary())
	fmt.Println(e2.getSalary())

	s := NewStack()
	s.Push(&Node{10})
	s.Push(&Node{12})
	s.Push(&Node{14})
	s.Push(&Node{16})

	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop(), s.Pop())

	q := NewQueue()
	q.QueueIn(&Node{10})
	q.QueueIn(&Node{12})
	q.QueueIn(&Node{14})
	q.QueueIn(&Node{16})

	fmt.Println(q.QueueOut(), q.QueueOut(), q.QueueOut(), q.QueueOut(), q.QueueOut())
	fmt.Println(len(q.nodes))

}
