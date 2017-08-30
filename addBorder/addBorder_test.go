package addBorder

import "testing"

func CompareSlicesBCE(a, b []string) bool {
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

func TestAddBorder(t *testing.T) {
	testArray := []string{"abc", "ded"}
	actualResult := addBorder(testArray)

	var expectedResult = []string{"*****", "*abc*", "*ded*", "*****"}

	if !CompareSlicesBCE(actualResult, expectedResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}
