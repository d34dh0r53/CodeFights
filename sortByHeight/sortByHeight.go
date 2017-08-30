package sortByHeight

func sortByHeight(a []int) []int {
	var temp int

	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			for j := i + 1; j < len(a); j++ {
				if a[i] > a[j] && a[j] != -1 {
					temp = a[i]
					a[i] = a[j]
					a[j] = temp
				}
			}
		}
	}
	return a
}
