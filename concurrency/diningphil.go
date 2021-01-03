package concurrency

import "fmt"

type philosopher struct {
	name     string
	status   string
	neighbor *philosopher
}

func eat() {
	// change status to eating
}

func think() {

}

func makePhil(neighbor *philosopher, name string) *philosopher {

	phil := &philosopher{name, "", neighbor}
	return phil
}

func Dine() {

	var names = []string{"Kant", "Heidegger", "Wittgenstein",
		"Locke", "Descartes", "Newton", "Hume", "Leibniz"}

	var storePhils = make([]*philosopher, len(names))
	var phil *philosopher

	for i, name := range names {
		fmt.Println(i, phil, name)
		phil = makePhil(phil, name)
		// fmt.Println(name, phil.neighbor)
		storePhils[i] = phil
	}
	storePhils[0].neighbor = phil
	fmt.Println(storePhils[0].neighbor.name)

}
