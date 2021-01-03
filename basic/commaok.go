package basic

import (
	"fmt"
	"log"
	"time"
)

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func Offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

func say(a string) {
	fmt.Println(a)
}
func SHirit() {
	for i := 0; i < 3; i++ {
		go func(j int) {
			defer say("done")
			time.Sleep(time.Second)
			fmt.Printf("im %v", j)
		}(i)
	}

	for {
	}
}

func HHHH() {

	for i := 0; i < 12; i++ {
		if i == 9 {
			continue
		}
		fmt.Println(i)
	}
}
