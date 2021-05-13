package security

/*
	Trying out encoding and decoding of all kinds

	want to learn about marshalling, encoding, serialization, etc
	a few of things amongst many i want to try out are gob, bson, protobufs etc
	trying to send them over networks etc.

	I will also try some of the video streaming protocols i've never tried before such as
	RTMP, HLS, QUIC, etc

	I am aware the package name is kind of bad but imjust going to leave it this way
*/

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func DoStuff() {
	data, err := bson.Marshal(&Person{Name: "Bob", Phone: "992878123"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q", data)
}
