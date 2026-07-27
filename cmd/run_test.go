// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/martincostello/advent-of-go/cmd"
	"github.com/martincostello/advent-of-go/puzzles"
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

func TestCmdRun(t *testing.T) {
	tests := []struct {
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

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%02d", tt.year, tt.day), func(t *testing.T) {
			input := filepath.Join(
				inputDir,
				fmt.Sprintf("Y%d", tt.year),
				fmt.Sprintf("Day%02d", tt.day),
				"input.txt",
			)
			got, err := cmd.Run([]string{"--year", strconv.Itoa(tt.year), "--day", strconv.Itoa(tt.day), input}, t.Context())
			if err != nil {
				t.Fatalf("Run(%d, %d) returned an error: %v", tt.year, tt.day, err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Run(%d, %d) mismatch (-want +got):\n%s", tt.year, tt.day, diff)
			}
		})
	}
}

func TestCmdRunUnsolved(t *testing.T) {
	var (
		year  = 2001
		day   = 42
		input = filepath.Join(
			inputDir,
			fmt.Sprintf("Y%d", year),
			fmt.Sprintf("Day%02d", day),
			"input.txt",
		)
	)
	_, err := cmd.Run([]string{"--year", strconv.Itoa(year), "--day", strconv.Itoa(day), input}, t.Context())
	if err == nil {
		t.Fatalf("Run(%d, %d) did not return an error", year, day)
	}
}

func TestCmdRunInvalidInput(t *testing.T) {
	var (
		year  = 2015
		day   = 1
		input = filepath.Join(
			inputDir,
			"foo.txt",
		)
	)
	_, err := cmd.Run([]string{"--year", strconv.Itoa(year), "--day", strconv.Itoa(day), input}, t.Context())
	if err == nil {
		t.Fatalf("Run(%d, %d) did not return an error", year, day)
	}
}
