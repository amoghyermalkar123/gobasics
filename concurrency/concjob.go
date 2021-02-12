package concurrency

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Store struct {
	value string
}

func (s *Store) set(str string) {
	s.value = str
}

func (s *Store) get() string {
	return s.value
}

type ready struct{}

func Job(s *Store, done chan<- ready) {
	time.Sleep(100 * time.Millisecond)
	val := strconv.Itoa(12)
	s.set(val)
	// when our job is done it sends "ready" message to our main routine
	done <- ready{}
}

func Main() {
	fmt.Println("test")
	c := make(chan ready)
	store := &Store{}
	// starting an independent job
	go Job(store, c)
	// we can continue with further ops without waiting here for our job
	fmt.Println("Ops that can be continued without worring about Job()")
	log.Println("doing main job....")
	time.Sleep(time.Second)
	// now independent ops were done, for further ops we wait for our job to return
	<-c
	fmt.Println("job has finished")
	// now that our job returned, we can do further ops where the output of job is required
	fmt.Println("continuing further ops where job's output is required\t", store.get())
	fmt.Println("DONE :)")
}
