package exer

import (
	"fmt"
)

// Map implements map functionality on integers
func intmap(f func(x int) int, numbers []int) (new []int) {
	for _, item := range numbers {
		ret := f(item)
		new = append(new, ret)
	}
	return new
}

func strmap(f func(s string) string, strings []string) (new []string) {
	for _, item := range strings {
		ret := f(item)
		new = append(new, ret)
	}
	return new
}

func inter(x int) int {
	return x * 2
}

func strer(s string) string {
	return s + "concat"
}

func GenericMap(genericVariable interface{}) {
	switch genericVariable.(type) {
	case []int:
		new := intmap(inter, genericVariable.([]int))
		fmt.Println(new)
	case []string:
		new := strmap(strer, genericVariable.([]string))
		fmt.Println(new)
	}
}

/*
func main() {
	s := []int{1, 2, 3, 4, 5}
	var ei interface{}
	ei = []string{"a", "b", "c"}
	exer.GenericMap(ei)
	exer.GenericMap(s)
}
*/
