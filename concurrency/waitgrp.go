package concurrency

import (
	"fmt"
	"sync"
)

func W() {
	var max int = 5
	var wgA, wgB sync.WaitGroup

	wgA.Add(max)
	wgB.Add(1)
	for i := 0; i < max; i++ {
		go func() {
			wgB.Wait()
			fmt.Println("goroutine id: ", i)
			wgA.Done()
		}()
	}

	wgB.Done()
	wgA.Wait()
	fmt.Println("wait is over ..furhter ops begin")
}

func W2() {
	var max int = 3
	var wg sync.WaitGroup

	wg.Add(max)
	for i := 0; i < max; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops continue")

}
