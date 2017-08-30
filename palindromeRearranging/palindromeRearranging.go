package palindromeRearranging

func palindromeRearranging(inputString string) bool {
	// Start all character count frequencies at 0
	var charCount [26]int

	// Traverse the string, incrementing character frequency
	for i := 0; i < len(inputString); i++ {
		// 97 is the byte offset for the a-z characters, a == 97, b == 98, etc...
		charCount[inputString[i]-97]++
	}

	// Count the odd numbered frequencies
	var odds int
	for j := 0; j < len(charCount); j++ {
		if charCount[j]%2 == 1 {
			odds++
		}
	}

	// If there is more than one odd frequency a palindrome is impossible
	if odds < 2 {
		return true
	}

	return false
}
