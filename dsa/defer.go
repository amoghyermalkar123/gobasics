package dsa

func masta() (ret int) {
	defer func() {
		ret += 3
	}()
	return 0
}
