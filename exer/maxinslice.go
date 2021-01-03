package exer

// MaxInlice finds max in array
func MaxInlice(arr []int) int {
	var max int = 0
	for _, item := range arr {
		if max < item {
			max = item
		} else {
			continue
		}
	}
	return max
}
