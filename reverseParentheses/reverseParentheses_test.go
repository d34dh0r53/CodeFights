package reverseParentheses

import (
	"testing"
)

func TestReverseParentheses(t *testing.T) {
	/*	actualResult := reverseParentheses("a(bc)de")
		var expectedResult = "acbde"

		if actualResult != expectedResult {
			t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
		}
	*/

	actualResult2 := reverseParentheses("a(bcdefghijkl(mno)p)q")
	var expectedResult2 = "apmnolkjihgfedcbq"

	if actualResult2 != expectedResult2 {
		t.Fatalf("Expected %s but got %s", expectedResult2, actualResult2)
	}
}
