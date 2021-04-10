package playground

import (
	"fmt"
	"os"
)

type MessageHandleFunc func(string)

func GVVVH() {
	go func() {
		myHandler := func(handler MessageHandleFunc) {

		}

		myFun := func(msg string) {
			fmt.Println(msg)
		}
		myHandler(myFun)
	}()

	go func() {
		myHandler := func(handler MessageHandleFunc) {

		}

		myFun := func(msg string) {
			fmt.Println(msg)
		}
		myHandler(myFun)
	}()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Automate() {
	f, err := os.Create("migrate.sql")
	check(err)

	defer f.Close()

	_, er := f.WriteString("INSERT INTO curated_companies_source(name, location, companyType)" + "\n")
	check(er)

	_, er2 := f.WriteString("VALUES" + "\n")
	check(er2)

	for _, in := range curatedCompaniesTable {
		query := fmt.Sprintf("('%s', '%s', '%s'),", in.name, in.location, in.companyType)
		_, err := f.WriteString(query + "\n")
		check(err)
	}

	// fmt.Printf("wrote %d bytes\n", n3)
}

type lol struct {
	aaaa []int64
}

func HAHAHA() {
	var abc []int64

	abc = []int64{1, 2, 3}
	d := &lol{
		aaaa: abc,
	}
	fmt.Println(d.aaaa)
}

type shii struct {
	i int64
}

func Search(i int, locations []shii) bool {
	low := 0
	high := len(locations) - 1

	for low <= high {
		median := (low + high) / 2

		if int(locations[median].i) < i {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(locations) || int(locations[low].i) != i {
		return false
	}

	return true
}

func HHAMogh() {
	s := []shii{
		{
			1,
		},
		{
			2,
		},
		{
			3,
		},
		{
			4,
		},
	}

	fmt.Println(Search(4, s))
}

func OAOW() {
	maker := make(map[int][]int)

	ss := []int{1, 2, 3, 4}
	for i := range ss {
		maker[i] = append(maker[i], ss[i], ss[i]+1)
	}
	fmt.Println(maker)
}
