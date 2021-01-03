package advancedGo

import (
	"fmt"
	"time"
)

/*
observe-send-store pattern (below story is subject to change, plz read it and change it according to the current state of the code)

this is the story of a type named event who had a method. This method was started in the background inside the womb of a different function.
this function was birthed by a method of a different type. the responsibility of this type was to observe events. it was friends with
the event type. this function was called when messages were received on its channels. it would incoke this method of type event.
who sent messages to him though. A different function. noone knows what its type was.
what did our event types method do tho? isn't it obvious, give data to our event's bag.

this is what i call observe-send-store pattern
one type is responsible for some observation
one for storage
inbetween comms happen by any other patterns
*/

type event struct {
	state interface{}
}

type eventObserver struct {
	listeners      map[int64]map[int64]func(*event)
	nextListenerID int64
	typer          func() interface{} // fill this with some sort of state filler
}

type eagleView struct {
	pipe chan func() interface{}
}

// this method sets the change in event state that was observed by
func (eo *eventObserver) change() event {
	e := event{}

	if eo.typer != nil {
		e.state = eo.typer()
	}
	return e
}

type done struct{}

func (ev *eagleView) monitor(handler func() event, out chan<- done) {
	for {
		select {
		case otherHandler := <-ev.pipe:
			e := handler()
			fmt.Println(e.state)
			fmt.Println(otherHandler())
			out <- done{}
		}
	}
}

func (ev *eagleView) changeGiver(oth func() interface{}) {
	ev.pipe = make(chan func() interface{})
	<-time.After(time.Second)
	ev.pipe <- oth
}

func (e *eventObserver) changeInitializer(fn func() interface{}) {
	e.typer = fn
}

func E() {
	done := make(chan done)
	ev := &eagleView{}
	e := &eventObserver{}
	ff := func() interface{} {
		return "amogh is amazing"
	}
	e.changeInitializer(ff)
	go ev.monitor(e.change, done)
	ev.changeGiver(ff)
	for {
		select {
		case <-done:
			return
		}
	}
}
