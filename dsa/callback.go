package dsa

import (
	"fmt"
)

func callback(y int, f func(int) int) {
	ret := f(y)
	fmt.Println(ret)
}

var mapOfFuncs = map[int]func(x int) int{
	1: func(x int) int { return x * 2 },
	2: func(x int) int { return x * 4 },
}
