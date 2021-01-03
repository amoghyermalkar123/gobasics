package playground

import "fmt"

type tst struct {
	name string
}

var db tst

func Work() {
	g := &tst{
		name: "amogh",
	}
	db = *g
}

func P() {
	// t := tst{}
	Work()
	fmt.Print("m", db.name)
}
