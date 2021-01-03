package dsa

import (
	"fmt"
)

func funcAsValueWorkSpace() {
	value := func() string {
		return "amo"
	}
	fmt.Println(value())

	var mapStringToFunc = map[string]func() string{
		"a": func() string { return "amogh" },
		"i": func() string { return "ishan" },
	}

	fmt.Println(mapStringToFunc["i"]())
}

// Important Note:
// In Go, if your anonymous function returns something, then its necessary to mention its return type
// otherwise error is produced
