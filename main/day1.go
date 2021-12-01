package main

import (
	"adventOfCode2021"
	"io/ioutil"
	"strings"
)

func main() {
	println(adventOfCode2021.FindNumberOfIncreasingSequence(parseInput()))
	println(adventOfCode2021.FindNumberOfIncreasingSlidingWindow(parseInput()))
}

func parseInput() []string {
	file, err := ioutil.ReadFile("./data/day1.txt")
	if err != nil {
		panic("Data file not found!")
	}
	lines := strings.Split(string(file), "\n")
	return lines
}
