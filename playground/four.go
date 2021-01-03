package playground

import "fmt"

type My struct {
	name string
}

func (m *My) K() {
	m.name = "amogh"
}

func RecreateError_InvalidMemoryAddress_NilPointerDereference() {
	// var m *My
	m := &My{}
	m.K()
	fmt.Println(m.name)
}

// ERROR EXPLANATION:
// NOTE: above lines 14 and 16 create the error :
// invalid memory address or nil pointer dereference
// why :
// you created a variable thats a reader through address on line 14
// then you called the method K() thru this variable
// but
// methods (i.e. pointer receiver functions) in go take the "address" of the type
// what you gave was a variable that reads an address and not actually an address
// SOLUTION TO THE ERROR:
// line 15 and 16 in combination are correct
// because & operator returns the address and then you use it
