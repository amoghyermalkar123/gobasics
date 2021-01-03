package basic

import "fmt"

func High() {

}

func HighD() {

	picture := make([][]uint8, 10)

	pixels := make([]uint8, 5*10)
	for i := range picture {
		picture[i], pixels = pixels[:5], pixels[5:]
	}
	fmt.Println(picture)

}
