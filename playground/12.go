package playground

import (
	"fmt"
	"time"

	"github.com/panjf2000/ants"
)

type exponential func(x int) int32

func operate(op exponential, num int) {
	abc := op(num)
	fmt.Println(abc)
}

func Dodo() {
	operate(
		func(x int) int32 {
			return int32(x + x)
		},
		2,
	)
}

func sample() {
	fmt.Println("im sample")
}
func Wow() {
	option := ants.Options{}
	option.ExpiryDuration = time.Second * 3
	pool, err := ants.NewPool(1, ants.WithExpiryDuration(option.ExpiryDuration))

	if err != nil {
		fmt.Println("error")
	}
	sig := func() {
		sample()
	}
	// fmt.Println(pool.Cap())
	pool.Submit(sig)
	pool.Submit(sig)
	time.Sleep(time.Second * 3)
	pool.Submit(sig)
	fmt.Println(pool.Running())
	fmt.Println(pool.Free())
}

// ===============implementing what i call the loader pattern========================
func optionLoader(opts ...func(x int)) {
	// this function takes 'n' functions and executes them
	x := 10
	for _, fn := range opts {
		fn(x)
		x++
	}
}

func option1(y int) func(x int) {
	return func(x int) {
		fmt.Println(x + y)
	}
}

func option2(y int) func(x int) {
	return func(x int) {
		fmt.Println(x * y)
	}
}

func option3(y int) func(x int) {
	return func(x int) {
		fmt.Println(x - y)
	}
}

func EntryWayForChoosingOptions() {
	optionLoader(option1(2), option2(4), option3(6))
}

// but as you can see from above this code becomes pretty unreadable
// let's use golangs function types to make reused function signatures into a type

type generic func(x int)

func loader2(generics ...generic) {
	x := 2
	for _, fn := range generics {
		fn(x)
	}
}

func op1(y int) generic {
	return func(x int) {
		fmt.Println(x + y)
	}
}

func op2(y int) generic {
	return func(x int) {
		fmt.Println(x * y)
	}
}

func op3(y int) generic {
	return func(x int) {
		fmt.Println(x - y)
	}
}

func SecondWay() {
	optionLoader(option1(2), option2(4), option3(6))
}
