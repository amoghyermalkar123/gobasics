package jsonpr

import (
	"encoding/json"
	"fmt"
)

// basics of json with golang

type sample struct {
	Name string `json:"Name"`
}

func One() {
	s := &sample{
		Name: "amogh",
	}
	marshalled, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("err:\t%T", err)
	}
	fmt.Println(string(marshalled))
	s = &sample{}
	_ = json.Unmarshal(marshalled, s)
	fmt.Printf("%+v", s)
}
