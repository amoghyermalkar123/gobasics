package playground

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func (n *Node) Insert(key int) {
	// check if value is left and right are nil

	if n.left == nil && n.right == nil {
		n.value = key
	}

	if key < n.value {
		s := n.left
		s.Insert(key)
	} else if key > n.value {
		s := n.right
		s.Insert(key)
	} else {
		fmt.Println("key exists")
		return
	}

}

func BuildMap() {
	// var n Node : is same as below
	n := Node{}
	n.Insert(1)
}

type SomeShit struct {
	avar int
}

func SOLO() {
	var in1 *SomeShit
	res := here()
	in1 = res

	fmt.Println(in1.avar)
}

func here() *SomeShit {
	a := &SomeShit{
		avar: 12,
	}
	return a
}
