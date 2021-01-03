package basic

import "fmt"

func CheckIfString(value interface{}) {

	if str, ok := value.(string); ok {
		fmt.Println(str)
	} else {
		fmt.Println("not a string that u provided. Please provide a string")
	}

}
