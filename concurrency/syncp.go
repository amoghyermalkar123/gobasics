package concurrency

import (
	"fmt"
	"sync"
)

func Implement() {
	var counter int = 3
	var wg sync.WaitGroup
	wg.Add(counter)
	for i := 0; i < counter; i++ {
		fmt.Println(i)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("print lat cuz routine is a blocking call")
}
