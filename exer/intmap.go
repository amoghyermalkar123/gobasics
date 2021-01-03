package exer

// Map implements map functionality on integers
func Map(f func(x int) int, numbers []int) (new []int) {
	for _, item := range numbers {
		ret := f(item)
		new = append(new, ret)
	}
	return new
}
