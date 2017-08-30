package arrayMaximalAdjacentDifference

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAreEquallyStrong(t *testing.T) {

	type test struct {
		testInput      []int
		expectedResult int
	}

	var testArray []*test

	test1 := &test{[]int{2, 4, 1, 0}, 3}
	test2 := &test{[]int{1, 1, 1, 1}, 0}
	test3 := &test{[]int{-1, 4, 10, 3, -2}, 7}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)

	for x := 0; x < len(testArray); x++ {
		actualResult := arrayMaximalAdjacentDifference(testArray[x].testInput)
		var expectedResult = testArray[x].expectedResult
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Test: ", x, "Pass")
	}
}
