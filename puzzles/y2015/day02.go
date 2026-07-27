// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015

import (
	"fmt"
	"math"

	"github.com/martincostello/advent-of-go/puzzles"
)

// Day02 solves the puzzle for day 2 of Advent of Code 2015.
func Day02(input []string) (puzzles.PuzzleSolution, error) {
	var (
		totalArea   int
		totalLength int
	)

	for _, dimensions := range input {
		var l, w, h float64

		_, err := fmt.Sscanf(dimensions, "%fx%fx%f", &l, &w, &h)
		if err != nil {
			return puzzles.PuzzleSolution{}, fmt.Errorf("failed to parse dimensions %q: %v", dimensions, err)
		}

		area := 2*l*w + 2*h*w + 2*h*l
		extra := math.Min(math.Min(l*w, h*w), h*l)

		volume := h * l * w
		perimeter := math.Min(math.Min(2*(l+w), 2*(h+w)), 2*(h+l))

		totalArea += int(area + extra)
		totalLength += int(perimeter + volume)
	}

	return puzzles.PuzzleSolution{
		Part1: fmt.Sprint(totalArea),
		Part2: fmt.Sprint(totalLength),
	}, nil
}
