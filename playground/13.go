package playground

import "fmt"

func JY() {

	var yes bool
	for {
		select {
		default:
			if !yes {
				fmt.Println("dfgh")
				yes = true
			}
		}
	}
}

func AHAH() {
	for {
		if 3 < 5 {
			break
		}
	}
	fmt.Println("amogh")
}
