package areSimilar

import (
	"sort"
)

func compareSlicesBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func areSimilar(a []int, b []int) bool {
	var c int
	if compareSlicesBCE(a, b) {
		return true
	}

	for x := 0; x < len(a); x++ {
		if a[x] != b[x] {
			c++
		}
	}

	sort.Ints(a)
	sort.Ints(b)

	return c < 3 && compareSlicesBCE(a, b)
}
