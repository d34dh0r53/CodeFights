package main

import (
	"fmt"
)

func almostIncreasingSequence(sequence []int) bool {
	count := 0
	length := len(sequence)
	for iter := 0; iter < length - 1; iter++ {
		if sequence[iter] >= sequence[iter + 1] {
			count += 1
			if ((iter - 1 >= 0) && (sequence[iter - 1] >= sequence[iter + 1])) {
				if ((length - 2 > iter) && (sequence[iter] >= sequence[iter + 2])) {
					return false
				}
			}
		}
	}
	if count <= 1 {
		return true
	} else {
		return false
	}
}

func main() {
	seq := []int {1,3,2,1}
	fmt.Println(almostIncreasingSequence(seq))
}
