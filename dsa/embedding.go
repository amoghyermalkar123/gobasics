package dsa

import "fmt"

// Original type
type Original struct {
	field string
}

// NewOriginal is a copy of Original type
type NewOriginal Original

// Copy is embedding type
type Copy struct{ Original }

// OriginalMethodOne just prints some random shit
func (o *Original) OriginalMethodOne() {
	fmt.Println("sample")
}
