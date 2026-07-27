// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"fmt"
	"time"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func Run(args []string) (puzzles.PuzzleSolution, error) {
	year, day, raw, err := parse(args)
	if err != nil {
		return puzzles.PuzzleSolution{}, fmt.Errorf("failed to parse arguments: %w", err)
	}

	fmt.Printf("Solving Advent of Code for day %02d of %04d\n", day, year)
	fmt.Println()

	data := puzzles.PuzzleData(raw)

	var input = &puzzles.PuzzleInput{
		Year:  year,
		Day:   day,
		Input: &data,
	}

	started := time.Now()

	solution, err := solver.Solve(input)
	if err != nil {
		return puzzles.PuzzleSolution{}, fmt.Errorf("failed to solve puzzle: %w", err)
	}

	ended := time.Now()

	fmt.Printf("Part 1: %s\n", solution.Part1)
	fmt.Printf("Part 2: %s\n", solution.Part2)

	if len(solution.Visualization) > 0 {
		fmt.Printf("Visualization:\n%s\n", solution.Visualization)
	}

	fmt.Println()
	fmt.Printf("Solved in %s\n", ended.Sub(started))

	return solution, nil
}
