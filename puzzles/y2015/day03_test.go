// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015Day03(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  puzzles.PuzzleSolution
	}{
		{">", puzzles.PuzzleSolution{Part1: "1", Part2: "2"}},
		{"^>v<", puzzles.PuzzleSolution{Part1: "4", Part2: "3"}},
		{"^v^v^v^v^v", puzzles.PuzzleSolution{Part1: "2", Part2: "11"}},
		{"^v", puzzles.PuzzleSolution{Part1: "2", Part2: "3"}},
		{"^>v<", puzzles.PuzzleSolution{Part1: "4", Part2: "3"}},
		{"^v^v^v^v^v", puzzles.PuzzleSolution{Part1: "2", Part2: "11"}},
	}

	for _, tt := range tests {
		c := tt
		t.Run(c.input, func(t *testing.T) {
			t.Parallel()
			got, err := y2015.Day03(c.input)
			require.NoError(t, err, "Day03(%q) returned error: %v", c.input, err)
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Errorf("Day03(%q) mismatch (-want +got):\n%s", c.input, diff)
			}
		})
	}
}

func TestY2015Day03RejectsInvalidDirections(t *testing.T) {
	t.Parallel()
	input := "invalid"
	_, err := y2015.Day03(input)
	require.Error(t, err, "Day03(%q) did not return an error", input)
}
