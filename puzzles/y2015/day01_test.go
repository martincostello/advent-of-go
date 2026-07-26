// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestSolveDay01(t *testing.T) {
	tests := []struct {
		input           string
		wantFloor       string
		wantInstruction string
	}{
		{"(())", "0", "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := y2015.Day01(tt.input)
			if got.Part1 != tt.wantFloor || got.Part2 != tt.wantInstruction {
				t.Errorf("Day01(%q) = (%q, %q), want (%q, %q)", tt.input, got.Part1, got.Part2, tt.wantFloor, tt.wantInstruction)
			}
		})
	}
}
