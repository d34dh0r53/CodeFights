package palindromeRearranging

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPalindromeRearranging(t *testing.T) {

	type test struct {
		testString     string
		expectedResult bool
	}

	var testArray []*test

	test1 := &test{"aabb", true}
	test2 := &test{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc", false}
	test3 := &test{"abbcabb", true}
	test4 := &test{"zyyzzzzz", true}
	test5 := &test{"z", true}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)
	testArray = append(testArray, test4)
	testArray = append(testArray, test5)

	for x := 0; x < len(testArray); x++ {
		fmt.Println("Test:", x+1, "String: ", testArray[x].testString)
		actualResult := palindromeRearranging(testArray[x].testString)
		var expectedResult = testArray[x].expectedResult
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Pass")
	}
}
