package dsa

import "fmt"

// NameAge is a basic custom made struct
type NameAge struct {
	name string
	age  int
}

// Do is a basic receiever function
func (n *NameAge) Do(name string, age int) {
	n.name = name
	n.age = age
	fmt.Println(n)
}


