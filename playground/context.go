package playground

import (
	"fmt"
	"time"
)

type done = struct{}

func NtoOneNotificationPattern(ucanstart <-chan done) {

	<-ucanstart

	fmt.Println("get start")
}

func workersforNtoOne(id int, imdone chan<- done) {
	time.Sleep(time.Second * 2)
	fmt.Printf("\n%v monitor here, im ready\n", id)
	imdone <- done{}
}

func Starta() {
	imready := make(chan done)

	go workersforNtoOne(1, imready)
	go NtoOneNotificationPattern(imready)
	for {

	}
}

func OneToNPattern() {
	iWillCancelChildren := make(chan done)
	givemeOutput := make(chan int)
	for i := 0; i < 3; i++ {
		j := i
		go workersforOneToN(j, iWillCancelChildren, givemeOutput)
	}

	for op := range givemeOutput {
		fmt.Printf("\none of my child gave me this %v, amazing!\n", op)
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("fuck my kids")
			return
		}
	}
}

func workersforOneToN(id int, myParentCancelledMme <-chan done, hereisyourOutput chan<- int) {
	for {
		select {
		case <-myParentCancelledMme:
			fmt.Println("im not going to complete, i'll only waste resources")
			return
		default:
			fmt.Printf("im starting %v\t", id)
			// time.Sleep(time.Second * 2)
			fmt.Printf("\n%v here, i'm done working\n", id)
			hereisyourOutput <- id * 2
			return
		}
	}
}
