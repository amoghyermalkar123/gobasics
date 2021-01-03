package advancedGo

import "fmt"

type funcMatcher struct {
	f   func(interface{}) bool
	err interface{}
}

func MatchFunc(f func(i interface{}) bool) *funcMatcher {
	return &funcMatcher{f: f}
}

type weird struct {
	name string
}

func WOWO() {
	MatchFunc(func(i interface{}) bool {
		r := i.(weird)
		fmt.Println(r)
		if 1 < 2 {
			return true
		}
		return true
	})
}
