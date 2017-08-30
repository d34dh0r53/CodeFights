package areEquallyStrong

func arrSort(inputArray [2]int) [2]int {
	if inputArray[0] <= inputArray[1] {
		return inputArray
	}
	return [2]int{inputArray[1], inputArray[0]}
}

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	// Setup the arrays
	var yourStrength [2]int
	var friendsStrength [2]int

	yourStrength[0] = yourLeft
	yourStrength[1] = yourRight

	friendsStrength[0] = friendsLeft
	friendsStrength[1] = friendsRight

	// Sort the arrays, we care about arms being equal, doesn't matter which side
	yourStrength = arrSort(yourStrength)
	friendsStrength = arrSort(friendsStrength)

	// Retruns true if the strengths are the same, false if they are different
	return yourStrength == friendsStrength
}
