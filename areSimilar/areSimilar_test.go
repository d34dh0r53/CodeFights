package areSimilar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAreSimilar(t *testing.T) {

	type test struct {
		testArrayA     []int
		testArrayB     []int
		expectedResult bool
	}

	var testArray []*test

	test1 := &test{[]int{1, 2, 3}, []int{1, 2, 3}, true}
	test2 := &test{[]int{1, 2, 3}, []int{2, 1, 3}, true}
	test3 := &test{[]int{1, 2, 2}, []int{2, 1, 1}, false}
	test4 := &test{[]int{4, 6, 3}, []int{3, 4, 6}, false}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)
	testArray = append(testArray, test4)

	for x := 0; x < len(testArray); x++ {
		fmt.Println("Test:", x+1)
		actualResult := areSimilar(testArray[x].testArrayA, testArray[x].testArrayB)
		var expectedResult = testArray[x].expectedResult
		assert.Equal(t, expectedResult, actualResult)
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Pass")
	}
}
