package arrayChange

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArrayChange(t *testing.T) {

	type test struct {
		testArray      []int
		expectedResult int
	}

	var testArray []*test

	test1 := &test{[]int{1, 1, 1}, 3}
	test2 := &test{[]int{-1000, 0, -2, 0}, 5}
	test3 := &test{[]int{2, 1, 10, 1}, 12}
	test4 := &test{[]int{2, 3, 3, 5, 5, 5, 4, 12, 12, 10, 15}, 13}
	test5 := &test{[]int{-787, -773, -93, 867, -28, 118, 372, 255,
		355, 598, -179, -752, 794, 961, -84, 296, -714, 14, 666, -265,
		-905, 942, -691, -379, -698, -650, 637, 523, 709, -674, 574,
		-239, 805, -434, 597, -677, 664, 384, 726, -389, -387, 772,
		-603, 685, 249, 446, -631, 454, 983, 867, -158, 932, -440,
		891, -12, 400, -916, 503, 185, -802, -255, 207, -952, -506,
		-689, 425, 747, -907, -30, 102, 553, 981, -664, 75, -957,
		-42, 99, -750, -277, 686, -884, -972, 470, 32, 439, 163,
		887, 895, -555, -654, 793, 333, 143, 73, 181, -512, -915,
		-68, 542, 799}, 89510}

	testArray = append(testArray, test1)
	testArray = append(testArray, test2)
	testArray = append(testArray, test3)
	testArray = append(testArray, test4)
	testArray = append(testArray, test5)

	for x := 0; x < len(testArray); x++ {
		fmt.Println("Test:", x+1, "Input Array: ", testArray[x].testArray)
		actualResult := arrayChange(testArray[x].testArray)
		var expectedResult = testArray[x].expectedResult
		require.Equal(t, expectedResult, actualResult)
		fmt.Println("Pass")
	}
}
