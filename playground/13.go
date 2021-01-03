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
