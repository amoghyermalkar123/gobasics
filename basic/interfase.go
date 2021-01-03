package basic

import (
	"fmt"
)

type someType struct {
	name string
}

type someOtherType struct {
	age int
}

func (st *someType) CommonBehavior() {
	st.name = "amogh"
	fmt.Println(st.name)
}

func (sot someOtherType) CommonBehavior() {
	sot.age = 22
	fmt.Println(sot.age)
}

type Behavior interface {
	CommonBehavior()
}

func motherFunction(b Behavior) {
	b.CommonBehavior()
}

func DoIntShit() {
	one := &someType{}
	two := &someOtherType{}
	// two := someOtherType{}
	motherFunction(one)
	motherFunction(two)
}
