// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package main

import (
	"fmt"
	"time"

	"github.com/martincostello/advent-of-go/cmd"
	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func main() {
	year, day, raw := cmd.Parse()

	fmt.Printf("Solving Advent of Code for day %02d of %04d\n", day, year)
	fmt.Println()

	data := puzzles.PuzzleData(raw)

	var input = &puzzles.PuzzleInput{
		Year:  year,
		Day:   day,
		Input: &data,
	}

	started := time.Now()

	solution := solver.Solve(input)

	ended := time.Now()

	fmt.Printf("Part 1: %s\n", solution.Part1)
	fmt.Printf("Part 2: %s\n", solution.Part2)

	if len(solution.Visualization) > 0 {
		fmt.Printf("Visualization:\n%s\n", solution.Visualization)
	}

	fmt.Println()
	fmt.Printf("Solved in %s\n", ended.Sub(started))
}
