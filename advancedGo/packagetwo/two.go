package packagetwo

import "fmt"

type TwoType struct {
	age int
}

func (t *TwoType) Do() {
	t.age = 22
	fmt.Printf("Age:%d", t.age)
}

func (t *TwoType) DoAnother() {
	fmt.Println("Doing another")
}
