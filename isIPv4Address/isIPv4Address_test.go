package isIPv4Address

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsIPv4Address(t *testing.T) {

	type test struct {
		testInput      string
		expectedResult bool
	}

	var testArray []*test

	test1 := &test{"172.16.254.1", true}
	test2 := &test{"172.316.254.1", false}
	test3 := &test{".254.255.0", false}
	test4 := &test{"1.1.1.1a", false}
	test5 := &test{"172.16.254.1", true}
	test6 := &test{"172.16.254.1", true}
	test7 := &test{"172.16.254.1", true}
	test8 := &test{"172.16.254.1", true}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)
	testArray = append(testArray, test4)
	testArray = append(testArray, test5)
	testArray = append(testArray, test6)
	testArray = append(testArray, test7)
	testArray = append(testArray, test8)

	for x := 0; x < len(testArray); x++ {
		actualResult := isIPv4Address(testArray[x].testInput)
		var expectedResult = testArray[x].expectedResult
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Test: ", x+1, "Pass")
	}
}
