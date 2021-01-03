package exer

import "fmt"

func Childone(done chan bool) {
	fmt.Println("ping")
	done <- true
}

func Childtwo(done chan bool) {
	fmt.Println("pong")
	done <- true
}
