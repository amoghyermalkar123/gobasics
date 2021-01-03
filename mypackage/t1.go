package mypackage

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

const (
	closedPool = "worker pool has reached maximum capacity"
)

type worker struct {
	// a boolean indicating that max capacity of the pool has reached
	maxCapacity bool
	// capacity of the pool containing amount of worker goroutines
	poolCapacity int
	// channel where tasks are sent to be executed concurrently
	workers chan func()
	// returns the number of goroutines running currently
	active int32
}

type workerOptions struct {
	controlKids bool
	timeout     time.Duration
	withContext context.Context
}

type computeSettings func(options *workerOptions)

func loadSettings(applySettings computeSettings) *workerOptions {
	// allocate space for new settings object
	allSettings := new(workerOptions)
	// call the function value of a function type
	applySettings(allSettings)
	return allSettings
}

// function as a value used in conjunction with function-type concept
func setOptionsForWorkers(settings workerOptions) computeSettings {
	// this is function as a value
	// gets returned from set() method which is called in load settings
	// which first makes space for an object then calls this function
	return func(s *workerOptions) {
		*s = settings
	}
}

func CreatePool() *workerOptions {
	options := workerOptions{}
	options.controlKids = true
	return loadSettings(setOptionsForWorkers(options))
}

func (w *worker) submitTasks(task func()) {
	w.workers <- task
}

func (w *worker) setup(capacity int) {
	w.poolCapacity = capacity
}

func (w *worker) getActive() int32 {
	a := atomic.LoadInt32(&w.active)
	return a
}

func (w *worker) inc() {
	if w.maxCapacity {
		return
	}
	w.poolCapacity++
}

func (w *worker) run() string {
	if !w.maxCapacity {
		go func() {
			defer func() { w.poolCapacity-- }()
			for funcs := range w.workers {
				w.inc()
				funcs()
			}
		}()
	} else {
		return closedPool
	}
	return ""
}

func sample() {
	fmt.Println("i'm a sample task")
}

func Mwain() {
	task1 := func() {
		sample()
	}
	// var w *worker creates a variable "w" that points to nil ?? why ??
	w := &worker{}
	w.workers = make(chan func())
	go w.run()
	w.submitTasks(task1)
	w.submitTasks(task1)
	for {
		// lol
		time.Sleep(time.Second)
		return
	}
}
