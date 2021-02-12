package playground

import (
	"context"
	"fmt"
	"sort"
	"time"
)

func M() {
	pctx := context.Background()
	ctx, cancel := context.WithCancel(pctx)
	go w(ctx)
	cancel()
	for {
	}
}

func w(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return nil
		default:
			time.Sleep(time.Second * 5)
			fmt.Println("working")
		}
	}
}

func UI() {

	slice_2 := []string{"FGHJ", "FGH", "gfg", "KJHGF"}

	var f1, f2, f3 string
	f1 = "GEEKs"
	f2 = "C"
	f3 = "gfg"

	sort.Strings(slice_2)

	fmt.Println("Slice 2: ", slice_2)

	res1 := sort.SearchStrings(slice_2, f1)
	res2 := sort.SearchStrings(slice_2, f2)
	res3 := sort.SearchStrings(slice_2, f3)

	fmt.Println(res1, res2, res3)

}
