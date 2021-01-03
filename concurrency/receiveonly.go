package concurrency

import (
	"fmt"
	"time"
)

func returnsChannel() chan int {
	channel := make(chan int)
	return channel
}

func back(c chan int) {
	c <- 2
}

func back2(c chan int) {
	time.Sleep(5 * time.Second)
	c <- 3
}
func sumSomething(a, b int) {
	fmt.Println(a*a, b*b)
}

func MainJob() {
	c := returnsChannel()
	go back(c)
	fmt.Println("shai98sgh")
	// go back(c)
	sumSomething(<-c, <-c)
}
