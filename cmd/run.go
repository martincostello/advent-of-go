// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

type Environment struct {
	Args   []string
	Stdout io.Writer
	Stderr io.Writer
}

func Run(env *Environment, ctx context.Context) (*puzzles.PuzzleSolution, error) {
	options, err := Parse(env.Stderr, env.Args...)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(options.FileName)

	if err != nil {
		return nil, fmt.Errorf("reading file %q failed: %w", options.FileName, err)
	}

	_, _ = fmt.Fprintf(env.Stdout, "Solving Advent of Code for day %02d of %04d\n\n", options.Day, options.Year)

	data := puzzles.PuzzleData(bytes)

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

	_, _ = fmt.Fprintf(env.Stdout, "Part 1: %s\n", solution.Part1)
	_, _ = fmt.Fprintf(env.Stdout, "Part 2: %s\n", solution.Part2)

	if len(solution.Visualization) > 0 {
		_, _ = fmt.Fprintf(env.Stdout, "Visualization:\n%s\n", solution.Visualization)
	}

	_, _ = fmt.Fprintln(env.Stdout)
	_, _ = fmt.Fprintf(env.Stdout, "Solved in %s\n", ended.Sub(started))

	return &solution, nil
}
