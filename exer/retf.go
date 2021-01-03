package exer

import "fmt"

// PlusTwo returns a function
func PlusTwo() (f func(int)) {
	return func(x int) {
		fmt.Println(x + 2)
	}
}

// Maper just returns a 2-multiple of given argument
func Maper(x int) int {
	return x * 2
}
