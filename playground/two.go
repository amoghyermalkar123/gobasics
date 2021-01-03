package playground

import (
	"fmt"
)

// A basic program to modify accordingly to check the basic
// channel functionality..trying sleeping channels, etc

// figure out: why dont twoM and threeM receive the same message from the channel
func OneM() {
	c := make(chan int)
	// go threeM(c)
	// go fourM(c)
	go twoM(c)
	c <- 2
	<-c
	// time.Sleep(5 * time.Second)
}

func twoM(c chan int) {
	a := <-c
	fmt.Println("two", a)
	c <- 3
}

func threeM(c chan int) {
	a := <-c
	fmt.Println("three", a)
	// c <- 4
}

func fourM(c chan int) {
	a := <-c
	fmt.Println("four", a)
	// c <- 6
}

// answer : one specific channel can only be used by 2 goroutines.
