package arrayChange

import (
	"math"
)

func arrayChange(inputArray []int) int {
	last := len(inputArray) - 1
	var changes int
	for x := 1; x <= last; x++ {
		current := inputArray[x]
		previous := inputArray[x-1]
		pDelta := previous - current
		if pDelta == 0 {
			/* We're equal to the previous element, so increment by one */
			changes++
			inputArray[x]++
		} else if current > previous {
			/* We're already greater than the previous element so nothing to do */
			continue
		} else {
			/* Increment by the delta + 1 */
			changes += int(math.Abs(float64(pDelta))) + 1
			inputArray[x] = current + int(math.Abs(float64(pDelta))) + 1
		}
	}
	return changes
}
