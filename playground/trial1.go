package playground

import (
	"encoding/json"
	"fmt"
)

type event struct {
	Status interface{}
}

type pollVote struct {
	PollID int64
}

type Poll struct {
	Votes []*pollVote
}

func one() {
	e := &event{
		Status: &pollVote{
			PollID: 1,
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	eventOut := &event{Status: &pollVote{}}

	err = json.Unmarshal(b, eventOut)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %T\n", eventOut, eventOut.Status)
}

func two() {
	e := &event{
		Status: []*pollVote{
			{
				PollID: 1,
			},
			{
				PollID: 2,
			},
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	eventOut := &event{Status: []*pollVote{}}

	err = json.Unmarshal(b, eventOut)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %T\n", eventOut, eventOut.Status)
}

func three() {
	e := &event{
		Status: &[]*pollVote{
			{
				PollID: 1,
			},
			{
				PollID: 2,
			},
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	eventOut := &event{Status: new([]*pollVote)}

	err = json.Unmarshal(b, eventOut)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %T\n", eventOut, eventOut.Status)
}

func four() {
	e := &event{
		Status: &Poll{
			Votes: []*pollVote{
				{
					PollID: 1,
				},
				{
					PollID: 2,
				},
			},
		},
	}

	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	eventOut := &event{Status: &Poll{}}

	err = json.Unmarshal(b, eventOut)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v %T\n", eventOut, eventOut.Status)
}

func U() {
	one()
	two()
	three()
	four()
}
