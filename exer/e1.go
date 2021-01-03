package exer

// BubbleSort implements bubble sort algorithm
func BubbleSort(numbers []int) (array []int) {

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	array = numbers
	return array
}
