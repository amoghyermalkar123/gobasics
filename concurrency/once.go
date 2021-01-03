package concurrency

import (
	"fmt"
	"sync"
)

func IOnce() {
	toBeExecutedOnce := func() {
		fmt.Println("im executed once")
	}

	var wg sync.WaitGroup
	var o sync.Once

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println("executed multiple times")
			o.Do(toBeExecutedOnce)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("all goroutines have finished their execution")

}
