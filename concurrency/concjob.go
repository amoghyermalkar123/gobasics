package concurrency

import (
	"fmt"
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

func Job(s *Store) {
	i := 0

	for {
		time.Sleep(100 * time.Millisecond)
		i++
		val := strconv.Itoa(i)
		s.set(val)
	}
}

func Main() {
	fmt.Println("test")
	c := time.Tick(1 * time.Second)
	store := &Store{}
	go Job(store)
	for {
		select {
		case <-c:
			fmt.Println(store.get())
		}
	}
}
