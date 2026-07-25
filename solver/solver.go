package solver

import (
	"github.com/martincostello/advent-of-go/puzzles"
	Y2015 "github.com/martincostello/advent-of-go/puzzles/Y2015"
)

var unsolved = puzzles.PuzzleSolution{
	Part1:         "Not implemented",
	Part2:         "Not implemented",
	Visualization: "",
}

func Solve(input *puzzles.PuzzleInput) puzzles.PuzzleSolution {
	if input.Year == 2015 {
		switch input.Day {
		case 1:
			return Y2015.Day01(input.Input)
		default:
			return unsolved
		}
	} else {
		return unsolved
	}
}
