package exer

import "fmt"

var sol []int

func recursive(value []int, weights []int, W int, itemIndex int, picked string) int {
	if W == 0 || itemIndex == len(weights) {
		return 0
	}

	if picked == "picked" {
		if contains(sol, value[itemIndex]) == false {
			sol = append(sol, value[itemIndex])
		} else {
		}
	} else {
	}

	if weights[itemIndex] > W {
		return recursive(value, weights, W-weights[itemIndex], itemIndex+1, "notpicked")
	}
	return max(value[itemIndex]+recursive(value, weights, W-weights[itemIndex], itemIndex+1, "picked"),
		recursive(value, weights, W-weights[itemIndex], itemIndex+1, "notpicked"))
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Knapsack() {
	value := []int{1, 2, 3, 4}
	weights := []int{3, 4, 2, 6}
	maxWeight := 7
	l := recursive(value, weights, maxWeight, 0, "start")
	fmt.Println(l, sol)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else if b > a {
		return b
	} else {
		return 0
	}
}
