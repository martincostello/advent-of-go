// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package solver

import (
	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

var unsolved = puzzles.PuzzleSolution{
	Part1:         "Not implemented",
	Part2:         "Not implemented",
	Visualization: "",
}

// Solve returns the solution for the given puzzle input, or unsolved if no
// solution has been implemented yet for the specified year and day.
func Solve(input *puzzles.PuzzleInput) puzzles.PuzzleSolution {
	if input.Year != 2015 {
		return unsolved
	}

	switch input.Day {
	case 1:
		return y2015.Day01(input.Input)
	case 2:
		return y2015.Day02(input.Input)
	default:
		return unsolved
	}
}
