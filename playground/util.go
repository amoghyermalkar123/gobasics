package playground

import "math/rand"

func GetGraph() []Dedge {
	allEdge := []Dedge{
		Dedge{
			source:      &Dnode{value: 0},
			destination: &Dnode{value: 1},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 0},
			destination: &Dnode{value: 7},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 1},
			destination: &Dnode{value: 2},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 0},
			destination: &Dnode{value: 7},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 2},
			destination: &Dnode{value: 8},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 2},
			destination: &Dnode{value: 3},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 2},
			destination: &Dnode{value: 5},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 7},
			destination: &Dnode{value: 8},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 7},
			destination: &Dnode{value: 6},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 8},
			destination: &Dnode{value: 6},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 6},
			destination: &Dnode{value: 5},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 5},
			destination: &Dnode{value: 3},
			weight:      rand.Intn(25),
		},
		Dedge{
			source:      &Dnode{value: 3},
			destination: &Dnode{value: 4},
			weight:      rand.Intn(25),
		},
	}
	return allEdge
}
