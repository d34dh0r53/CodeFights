package arrayMaximalAdjacentDifference

import (
	"math"
)

func arrayMaximalAdjacentDifference(inputArray []int) int {
	var maximalDifference int
	for x := 0; x < len(inputArray)-1; x++ {
		current := inputArray[x]
		next := inputArray[x+1]
		tempDifference := 0

		if current >= next {
			tempDifference = int(math.Abs(float64(current - next)))
		} else {
			tempDifference = int(math.Abs(float64(next - current)))
		}

		if tempDifference > maximalDifference {
			maximalDifference = tempDifference
		}
	}

	return maximalDifference
}
