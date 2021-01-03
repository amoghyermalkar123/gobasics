package playground

import "fmt"

type Testser struct {
	name string
}

type He struct {
	name string
}

func SestShit() {
	var tes Testser
	h := &He{}

	h.name = "amogh"
	if 3 < 5 {
		tes.name = h.name
	}

	fmt.Println(tes)
}
