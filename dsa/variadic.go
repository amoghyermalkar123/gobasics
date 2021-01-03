package dsa

import (
	"fmt"
)

// Variadic receives any amount of prameters and prints them
func Variadic(arg ...int) {
	for i := range arg {
		fmt.Printf("%v", arg[i])
	}
}
