package dsa

import (
	"fmt"
)

type student struct {
	name string
	age  int
}

func arrayWorkspace() {
	var balance [10]string
	var food = []int{1, 2, 3, 4}
	fmt.Println(balance)
	fmt.Println(food)
}

func sliceWorkspace() {
	// var balance []int
	var parent = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var child = parent[0:5]
	fmt.Println(child)
}

func structSpace() {
	var amogh = student{"amogh", 22}
	ishan := student{"ishan", 18}
	fmt.Println(amogh, ishan)
}
