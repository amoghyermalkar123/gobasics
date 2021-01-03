package exer

// Fibonacci implements the sequence
func Fibonacci(n int) (array []int) {
	x := make([]int, n)
	x[0], x[1] = 1, 1

	for i := 2; i < n; i++ {
		x[i] = x[i-1] + x[i-2]
	}
	array = x
	return array
}
