package alternatingSums

func alternatingSums(a []int) []int {
	weightSums := []int{0, 0}
	var left, right int

	for x := 0; x < len(a); x++ {
		if x%2 == 0 {
			left += a[x]
		} else {
			right += a[x]
		}
	}
	weightSums[0] = left
	weightSums[1] = right

	return weightSums
}
