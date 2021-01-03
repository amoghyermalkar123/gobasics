package playground

import (
	"context"
	"fmt"
	"log"
	"time"

	errgroup "golang.org/x/sync/errgroup"
)

func OI() {

	ctx := context.Background()
	mg, dctx := errgroup.WithContext(ctx)
	errs := make(chan error)

	mg.Go(func() (err error) {
		rr := func(ctd context.Context) error {
			fmt.Println("child 1 op started")
			time.Sleep(time.Second * 5)

			return fmt.Errorf("---")
		}(dctx)
		if rr != nil {
			log.Println("child 1 err is here")
			return rr
		}
		return nil
	})

	mg.Go(func() (err error) {
		rr := func(ctd context.Context) error {
			fmt.Println("child 2 op started")
			time.Sleep(time.Second * 2)
			return fmt.Errorf("---")
		}(dctx)
		if rr != nil {
			log.Println("child 2 err is here")
			return rr
		}
		return nil
	})

	go func() {
		err := mg.Wait()
		if err != nil {
			errs <- err
		}
	}()

	for {

		select {
		case <-errs:
			fmt.Println("first error at", time.Now())
		case <-time.After(time.Second * 10):
			return
		}
	}
}
