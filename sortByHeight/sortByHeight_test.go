package sortByHeight

import "testing"

func CompareSlicesBCE(a, b []int) bool {
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

func TestSortByHeight(t *testing.T) {
	testArray := []int{-1, 150, 190, 170, -1, -1, 160, 180}
	actualResult := sortByHeight(testArray)

	var expectedResult = []int{-1, 150, 160, 170, -1, -1, 180, 190}

	if !CompareSlicesBCE(actualResult, expectedResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
