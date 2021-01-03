package concurrency

import (
	"fmt"
	"sync"
)

func ImplementMutualExclusion() {

	var wg sync.WaitGroup
	var globalArray = []int{1, 2, 3, 4, 56} //  shared memory
	// var once sync.Once
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(j int) {
			globalArray[0] = 2 * j
			fmt.Println("my updates:", globalArray)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("wait over")
}

func ImplementRaceCondition() {

	var wg sync.WaitGroup
	var globalArray = []int{1, 2, 3, 4, 56}
	// var once sync.Once
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(j int) {
			globalArray[0] = 2 * j //  shared memory without lock => race condition
			fmt.Println("my updates:", globalArray)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("wait over")
}
