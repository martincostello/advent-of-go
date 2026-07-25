package y2015

import (
	"fmt"

	"github.com/martincostello/advent-of-go/puzzles"
)

// Day01 solves the puzzle for day 1 of Advent of Code 2015.
func Day01(input *puzzles.PuzzleData) puzzles.PuzzleSolution {
	floor := 0
	instructionThatEntersBasement := -1
	hasVisitedBasement := false

	for i, ch := range input.AsString() {
		switch ch {
		case '(':
			floor++
		case ')':
			floor--
		}

		if !hasVisitedBasement && floor == -1 {
			instructionThatEntersBasement = i + 1
			hasVisitedBasement = true
		}
	}

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(floor),
		Part2: fmt.Sprint(instructionThatEntersBasement),
	}
}
