package playground

import (
	"context"
	"fmt"
	"time"
)

func ParentHere() {
	ctx := context.Background()
	parentCopy, cancelChildrens := context.WithCancel(ctx)

	for i := 0; i < 3; i++ {
		j := i
		go workerGoroutines(parentCopy, j)
	}

	select {
	case <-time.After(time.Second):
		cancelChildrens()
	}
}

func workerGoroutines(ctx context.Context, id int) {
	fmt.Printf("\n%v here, im working\n", id)
	time.Sleep(time.Second * 3)
	fmt.Printf("\n%v work done\n", id)
}
