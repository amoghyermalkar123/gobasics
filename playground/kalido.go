package playground

import (
	"fmt"
	"reflect"
)

type ATM interface {
	nameshit(string)
	ageshit(int)
}

type One struct {
	name string
	age  int
}

type Two struct {
	age int
}

func (o *One) nameshit(name string) {
	o.name = name
	fmt.Println(o.name)
}

func (t *One) ageshit(age int) {
	t.age = age
	fmt.Println(t.age)
}

func HowNotToUseInterface() {
	// the following code
	// var a ATM
	// a.nameshit()
	// a.ageshit()
	// produces a invalid memory addr/ nil pointer dereference error
	/*
		that's because thw variable a of type ATM is nil so it doesnt hold references to anything
		how to use interfaces then?
	*/
}

func HowToUseInterface() {
	// basically this line makes the interface variable point to a type
	var a ATM = &One{}
	// we cast out the type we want from the interface for whose method we want to interact with
	b := a.(*One)
	// printing the cast type
	fmt.Println(reflect.TypeOf(b))
	// accessing the behavior a.k.a the method we wanted from the interface
	b.nameshit("amogh")
}

func BetterWayToUseInterface(takeInAnyTypeThatImplementsInterface ATM, age int, name string) {
	takeInAnyTypeThatImplementsInterface.ageshit(age)
	takeInAnyTypeThatImplementsInterface.nameshit(name)
}

func UseTheBetterWay() {
	var a ATM = &One{}
	b := a.(*One)
	BetterWayToUseInterface(b, 4, "amogh")
}
