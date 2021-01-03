package dsa

import "fmt"

type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

type T struct {
	age int
}

func (s S) Get() int {
	return s.i
}

func (s S) Put(v int) {
	s.i = v
}

func (t T) Get() int {
	return t.age
}

func (t T) Put(age int) {
	t.age = age
}

func gen(i I) {
	fmt.Println(i.Get())
}

func GenericFunction(genf interface{}) {
	fmt.Println(genf.(I).Get())
}

func DoThis() {
	var inter I = S{29}
	var inter2 I = T{22}
	gen(inter)
	gen(inter2)
	GenericFunction(inter)
}
