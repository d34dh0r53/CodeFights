package areEquallyStrong

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAreEquallyStrong(t *testing.T) {

	type test struct {
		yourLeft       int
		yourRight      int
		friendsLeft    int
		friendsRight   int
		expectedResult bool
	}

	var testArray []*test

	test1 := &test{10, 15, 15, 10, true}
	test2 := &test{15, 10, 15, 10, true}
	test3 := &test{15, 10, 15, 9, false}
	test4 := &test{10, 5, 5, 10, true}
	test5 := &test{10, 15, 5, 20, false}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)
	testArray = append(testArray, test4)
	testArray = append(testArray, test5)

	for x := 0; x < len(testArray); x++ {
		actualResult := areEquallyStrong(testArray[x].yourLeft,
			testArray[x].yourRight,
			testArray[x].friendsLeft,
			testArray[x].friendsRight)
		var expectedResult = testArray[x].expectedResult
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Test: ", x, "Pass")
	}
}
