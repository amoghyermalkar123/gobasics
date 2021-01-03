package exer

import (
	"fmt"
	"sync"
)

func Tes() {
	var MAXROUTINES int = 3
	var wg sync.WaitGroup

	wg.Add(MAXROUTINES)

	for i := 0; i < 3; i++ {
		// fmt.Println(i)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("continue operation of main routine, children have done their work")
}
