package myrpc

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func Client() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()
	// Synchronous call
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Result: %d\n", reply)
}
