package Y2015

import (
	"fmt"

	"github.com/martincostello/advent-of-go/puzzles"
)

func Day01(input []byte) puzzles.PuzzleSolution {
	floor := 0
	instructionThatEntersBasement := -1
	hasVisitedBasement := false

	for i, ch := range input {
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
