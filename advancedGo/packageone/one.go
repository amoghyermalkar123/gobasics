package packageone

import "fmt"

type OneType struct {
	name string
}

func (o *OneType) Do() {
	o.name = "amogh"
	fmt.Println("my name is amogh")
}

func (o *OneType) DosomethingElse() {
	fmt.Println("Doing something else")
}
