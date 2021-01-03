package playground

import (
	"fmt"
)

func Give() {
	chana := make(chan int)

	// when you have multiple goroutines listening to the same channel only one of them receives the value from it
	// which depends on the scheduler (or someone else? idk..) which is randomly chosen
	// hence this is how data races don't happen
	// channels one-one pipelines at any given instance in time
	go func() {
		for i := range chana {
			go func(j int) {
				// time.Sleep(time.Second * 3)
				fmt.Println("A", j)
			}(i)
		}
	}()

	go func() {
		for i := range chana {
			go func(j int) {
				// time.Sleep(time.Second * 3)
				fmt.Println("B", j)
			}(i)
		}
	}()

	for i := 0; i < 10; i++ {
		chana <- i
	}
	fmt.Println("fghjk")
}
