package exer

import "fmt"

// PlusX is a closure
func PlusX(x int) (f func(int)) {
	return func(y int) {
		fmt.Println(x + y)
	}
}
