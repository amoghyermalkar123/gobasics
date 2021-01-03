package concurrency

import (
	"fmt"
	"time"
)

func sendToTick(tick chan bool) {
	time.Sleep(1 * time.Second)
	tick <- true
}

func sendd(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func Boom() {
	tick := make(chan bool)
	done := make(chan bool)
	go sendToTick(tick)
	go sendd(done)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-done:
			fmt.Println("tock.")
		default:
			break
		}
	}
}

func BoomBam() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
