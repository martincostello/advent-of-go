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
	var (
		stdout = env.Stdout
		stderr = env.Stderr
	)

	if stdout == nil {
		stdout = io.Discard
	}

	if stderr == nil {
		stderr = io.Discard
	}

	options, err := Parse(stderr, env.Args...)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(options.FileName)

	if err != nil {
		return nil, fmt.Errorf("reading file %q failed: %w", options.FileName, err)
	}

	_, _ = fmt.Fprintf(stdout, "Solving Advent of Code for day %02d of %04d\n\n", options.Day, options.Year)

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

	duration := time.Since(started)

	err = printSolution(solution, duration, stdout)
	if err != nil {
		return nil, err
	}

	return &solution, nil
}

func printSolution(s puzzles.PuzzleSolution, d time.Duration, w io.Writer) error {

	_, err := fmt.Fprintf(w, `Part 1: %s
Part 2: %s

Solved in %s
`, s.Part1, s.Part2, d.Round(time.Millisecond))

	/*
		if len(s.Visualization) > 0 {
			_, _ = fmt.Fprintf(w, "Visualization:\n%s\n", s.Visualization)
		}
	*/

	return err
}
