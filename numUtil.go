package adventOfCode2021

import "strconv"

func AtoI(a string) int {
	atoi, err := strconv.Atoi(a)
	if err != nil {
		panic("Invalid number!")
	}

	return atoi
}
