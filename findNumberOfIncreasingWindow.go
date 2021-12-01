package adventOfCode2021

func FindNumberOfIncreasingSlidingWindow(list []string) int {

	var numList []int
	for _, item := range list {
		numList = append(numList, AtoI(item))
	}
	if len(numList) < 3 {
		panic("Insufficient data!")
	}

	previousWindowSum := numList[0] + numList[1] + numList[2]
	numberOfIncreasingWindow := 0

	for i := 3; i < len(numList); i++ {
		currentSum := previousWindowSum + numList[i] - numList[i - 3]
		if currentSum > previousWindowSum {
			numberOfIncreasingWindow++
		}
		previousWindowSum = currentSum
	}
	return numberOfIncreasingWindow
}
