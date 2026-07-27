// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package solver

import (
	"context"

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
func Solve(input *puzzles.PuzzleInput, ctx context.Context) (puzzles.PuzzleSolution, error) {
	if input.Year != 2015 {
		return unsolved, nil
	}

	return solve2015(input, ctx)
}

func solve2015(input *puzzles.PuzzleInput, ctx context.Context) (puzzles.PuzzleSolution, error) {
	switch input.Day {
	case 1:
		return y2015.Day01(input.Input.String()), nil
	case 2:
		return y2015.Day02(input.Input.Lines())
	case 3:
		return y2015.Day03(input.Input.String())
	case 4:
		return y2015.Day04(input.Input.String(), ctx)
	case 5:
		return y2015.Day05(input.Input.Lines()), nil
	default:
		return unsolved, nil
	}
}
