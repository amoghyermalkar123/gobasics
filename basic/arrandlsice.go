package basic

import "fmt"

type LinesOfText [][]byte

var text = LinesOfText{
	[]byte("Now is the time"),
	[]byte("for all good gophers"),
	[]byte("to bring some fun to the party."),
}

func Yes() {
	fmt.Println(text)
}
