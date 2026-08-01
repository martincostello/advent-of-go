// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func Run(args []string, ctx context.Context) (*puzzles.PuzzleSolution, error) {
	options, err := Parse(args)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Solving Advent of Code for day %02d of %04d\n", options.Day, options.Year)
	fmt.Println()

	data := puzzles.PuzzleData(options.Input)

	var input = &puzzles.PuzzleInput{
		Year:  options.Year,
		Day:   options.Day,
		Input: &data,
	}

	started := time.Now()

	solution, err := solver.Solve(input, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to solve puzzle: %w", err)
	}

	ended := time.Now()

	fmt.Printf("Part 1: %s\n", solution.Part1)
	fmt.Printf("Part 2: %s\n", solution.Part2)

	if len(solution.Visualization) > 0 {
		fmt.Printf("Visualization:\n%s\n", solution.Visualization)
	}

	fmt.Println()
	fmt.Printf("Solved in %s\n", ended.Sub(started))

	return &solution, nil
}
