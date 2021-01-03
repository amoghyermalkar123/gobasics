package playground

import "fmt"

func FindDuplicate(strings []string) {
	var encountered = map[string]bool{}
	var result []string
	for item := range strings {
		if !encountered[strings[item]] == true {
			encountered[strings[item]] = true
			result = append(result, strings[item])
		}
	}

	fmt.Println(result)
}

func StringRep(input string) {
	for i, c := range input {
		// internally every character in a string is converted to its rune
		fmt.Println(input[i], c) // prints rune numbers
		// hence you have to use string() method and pass it the rune number
		fmt.Println(string(input[i]), string(c)) // prints characters cuz we converted rune to string using the built in method
	}
}

func ConsecutiveCheck(input string) {

}
