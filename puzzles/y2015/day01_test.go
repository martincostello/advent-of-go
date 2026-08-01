// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015Day01(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  puzzles.PuzzleSolution
	}{
		{"(())", puzzles.PuzzleSolution{Part1: "0", Part2: "-1"}},
	}

	for _, tt := range tests {
		c := tt
		t.Run(c.input, func(t *testing.T) {
			t.Parallel()
			got := y2015.Day01(c.input)
			if diff := cmp.Diff(c.want, got); diff != "" {
				t.Errorf("Day01(%q) mismatch (-want +got):\n%s", c.input, diff)
			}
		})
	}
}
