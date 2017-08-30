package alternatingSums

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

func TestAlternatingSums(t *testing.T) {
	testArray := []int{50, 60, 60, 45, 70}
	actualResult := alternatingSums(testArray)

	var expectedResult = []int{180, 105}

	if !CompareSlicesBCE(actualResult, expectedResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
