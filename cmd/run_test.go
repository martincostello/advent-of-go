// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/martincostello/advent-of-go/cmd"
	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

var (
	_, file, _, _ = runtime.Caller(0)
	root          = filepath.Dir(file)
	inputDir      = filepath.Join(
		root,
		"..",
		"input",
	)
)

var tests = []struct {
	year int
	day  int
	want puzzles.PuzzleSolution
}{
	{2015, 1, puzzles.PuzzleSolution{Part1: "232", Part2: "1783"}},
	{2015, 2, puzzles.PuzzleSolution{Part1: "1598415", Part2: "3812909"}},
	{2015, 3, puzzles.PuzzleSolution{Part1: "2565", Part2: "2639"}},
	{2015, 4, puzzles.PuzzleSolution{Part1: "346386", Part2: "9958218"}},
	{2015, 5, puzzles.PuzzleSolution{Part1: "236", Part2: "51"}},
}

func TestCmdRun(t *testing.T) {
	t.Parallel()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%02d", tt.year, tt.day), func(t *testing.T) {
			t.Parallel()
			input := filepath.Join(
				inputDir,
				fmt.Sprintf("Y%d", tt.year),
				fmt.Sprintf("Day%02d", tt.day),
				"input.txt",
			)
			env := newEnv("--year", strconv.Itoa(tt.year), "--day", strconv.Itoa(tt.day), input)
			got, err := cmd.Run(env, t.Context())
			require.NoError(t, err, "Run(%d, %d) returned an error: %v", tt.year, tt.day, err)
			if diff := cmp.Diff(tt.want, *got); diff != "" {
				t.Errorf("Run(%d, %d) mismatch (-want +got):\n%s", tt.year, tt.day, diff)
			}
		})
	}
}

func TestCmdRunUnsolved(t *testing.T) {
	t.Parallel()
	var (
		year  = time.Now().Year()
		day   = 25
		input = filepath.Join(
			inputDir,
			"Y2015",
			"Day01",
			"input.txt",
		)
		env = newEnv("--year", strconv.Itoa(year), "--day", strconv.Itoa(day), input)
	)
	_, err := cmd.Run(env, t.Context())
	require.Error(t, err, "Run(%d, %d) did not return an error", year, day)
}

func TestCmdRunInputFileNotFound(t *testing.T) {
	t.Parallel()
	var (
		year  = 2015
		day   = 1
		input = filepath.Join(
			inputDir,
			"foo.txt",
		)
	)
	env := newEnv("--year", strconv.Itoa(year), "--day", strconv.Itoa(day), input)
	_, err := cmd.Run(env, t.Context())
	require.Error(t, err, "Run(%d, %d) did not return an error", year, day)
}

func TestCmdRunInvalidFlag(t *testing.T) {
	t.Parallel()
	env := newEnv("--year", "2014", "--day", "1", "foo.txt")
	_, err := cmd.Run(env, t.Context())
	require.Error(t, err)
}

func BenchmarkCmdRun(b *testing.B) {
	for _, tt := range tests {
		b.Run(fmt.Sprintf("%d-%02d", tt.year, tt.day), func(b *testing.B) {
			inputFile := filepath.Join(
				inputDir,
				fmt.Sprintf("Y%d", tt.year),
				fmt.Sprintf("Day%02d", tt.day),
				"input.txt",
			)

			env := newEnv("--year", strconv.Itoa(tt.year), "--day", strconv.Itoa(tt.day), inputFile)

			var stderr strings.Builder
			options, err := cmd.Parse(&stderr, env.Args...)
			require.NoError(b, err, "Parse(%d, %d) returned an error: %v", tt.year, tt.day, err)

			bytes, err := os.ReadFile(options.FileName)
			require.NoError(b, err)

			data := puzzles.PuzzleData(bytes)

			var input = &puzzles.PuzzleInput{
				Year:  options.Year,
				Day:   options.Day,
				Input: &data,
			}

			b.ResetTimer()
			for b.Loop() {
				_, err = solver.Solve(input, b.Context())
				if err != nil {
					b.Fatalf("Solve(%d, %d) returned an error: %v", tt.year, tt.day, err)
				}
			}
		})
	}
}

func newEnv(args ...string) *cmd.Environment {
	var stdout strings.Builder
	var stderr strings.Builder
	return &cmd.Environment{
		Args:   args,
		Stdout: &stdout,
		Stderr: &stderr,
	}
}
