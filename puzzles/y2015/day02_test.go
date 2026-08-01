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

func TestY2015Day02(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input []string
		want  puzzles.PuzzleSolution
	}{
		{[]string{"2x3x4"}, puzzles.PuzzleSolution{Part1: "58", Part2: "34"}},
		{[]string{"1x1x10"}, puzzles.PuzzleSolution{Part1: "43", Part2: "14"}},
	}

	for _, tt := range tests {
		c := tt
		t.Run(c.input[0], func(t *testing.T) {
			t.Parallel()
			got, err := y2015.Day02(c.input)
			require.NoError(t, err, "Day02(%q) returned error: %v", c.input, err)
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Errorf("Day02(%q) mismatch (-want +got):\n%s", c.input, diff)
			}
		})
	}
}

func TestY2015Day02RejectsInvalidDimensions(t *testing.T) {
	t.Parallel()
	input := []string{"2x3x4", "1x1x10", "invalid"}
	_, err := y2015.Day02(input)
	require.Error(t, err, "Day02(%q) did not return an error", input)
}
