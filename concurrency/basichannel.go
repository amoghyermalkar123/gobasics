package concurrency

import (
	"fmt"
)

func childRoutine(s string, done chan bool) {
	fmt.Println(s)
	done <- false
}

func ParentRoutine() {
	done := make(chan bool)
	go childRoutine("world", done)
	fmt.Println("hello")
	isDone := <-done

	if isDone == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

}
