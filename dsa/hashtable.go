package dsa

import "fmt"

// correction type : open addressing

const maxx = 100

type store [maxx]string

func (arr *store) HashFunction(word string) {
	sum := 0
	for i, _ := range word {
		charIndex := word[i] - 'a'
		sum += int(charIndex)
	}
	if arr.CheckCollision(sum) == false {
		arr[sum] = word
	} else {
		index := correctCollision(sum)
		arr[index] = word
	}
}

func InsertKey() {

}

func SearchKey() {

}

func (array *store) CheckCollision(index int) bool {
	if array[index] != "" {
		return true
	}
	return false
}

func correctCollision(i int) int {
	return i + 1
}

func StartHas() {
	var array store
	array.HashFunction("amogh")
	array.HashFunction("amogh")
	fmt.Println(array[39], array[40])
}
