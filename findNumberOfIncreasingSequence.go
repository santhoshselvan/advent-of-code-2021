package adventOfCode2021


func FindNumberOfIncreasingSequence(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	numberOfIncreasingSequence := 0
	previousNumber := AtoI(lines[0])

	for i := 1; i < len(lines); i++ {
		currentNumber := AtoI(lines[i])

		if currentNumber > previousNumber {
			numberOfIncreasingSequence++
		}
		previousNumber = currentNumber
	}
	return numberOfIncreasingSequence
}

