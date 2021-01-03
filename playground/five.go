package playground

import (
	"fmt"
	"strconv"
	"time"
)

func Doooooooo() {
	as := []int{1, 2, 3}

	temp := make(chan int)

	for _, item := range as {
		go func(ite int) {
			time.Sleep(time.Second * 3)
			if ite == 2 {
				panic("")
			}
			temp <- ite
		}(item)
	}

	for {
		select {
		case some := <-temp:
			fmt.Println(some)
		}
	}
}

func Shoooo() {
	sha := make(map[int]string)
	as := []int{1, 2, 3}
	for i := range as {
		sha[i] = "string is " + strconv.Itoa(i)
	}

	for key := range sha {
		fmt.Println("---", key)
	}
}

func Roooooooooo() {
	sha := make(map[int]int)
	as := []int{1, 2, 3}
	for i := range as {
		sha[i] = i + i
	}

	for key := range sha {
		go func(k int) {
			fmt.Println(k)
		}(key)
	}

	time.Sleep(time.Second)
}

func Yooo() {
	sha := make(map[int]int)
	as := []int{1, 2, 3}
	for i := range as {
		sha[i] = i + i
	}

	for i := range sha {
		go func(j int) {
			fmt.Println(j)
			if j == 0 {
				panic("0")
			}
		}(i)
	}
	time.Sleep(time.Second * 2)
}

func Hoo() {
	sha := make(map[int]int)
	as := []int{1, 2, 3}
	for i := range as {
		sha[i] = i + i
	}
	ghere := make(chan int)

	go func() {
		time.Sleep(time.Second * 13)
		ghere <- 10
	}()

	yere := <-ghere
	sha[3] = yere

	for i := range sha {
		go func(nl int) {
			fmt.Println("Ha ghya he", sha[nl])
		}(i)
	}

	for {
	}
}

func Looo() {
	sha := make(map[int]int)
	as := []int{1, 2, 3}
	for i := range as {
		sha[i] = i + i
	}
	ghere := make(chan int)

	go func() {
		ghere <- 1
		ghere <- 2
		ghere <- 3
	}()

	data := <-ghere
	go func(int) {
		fmt.Println("this is my", data)
	}(data)

	for {
	}
}
