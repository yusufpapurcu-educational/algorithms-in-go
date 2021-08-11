package binarysearch

// In here we don't take any data actually. We just take length of array and function that will lead the way.
// isLittleThan function takes a index and compare element with wanted number.
func BinarySearch(arrayLength int, isLittleThan func(int) bool) int {
	answer, endFlag := 0, arrayLength

	for answer < endFlag {
		middleFlag := answer + (endFlag-answer)/2

		if !isLittleThan(middleFlag) {
			answer = middleFlag + 1 // Look second half
		} else {
			endFlag = middleFlag // Look first half
		}
	}

	return answer
}
