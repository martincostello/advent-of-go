// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package solver_test

import (
	"fmt"
	"testing"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func TestSolverSolveWhenUnsolved(t *testing.T) {
	var data puzzles.PuzzleData = []byte("invalid")
	tests := []struct {
		input puzzles.PuzzleInput
	}{
		{input: puzzles.PuzzleInput{Year: 2014, Day: 1, Input: &data}},
		{input: puzzles.PuzzleInput{Year: 2015, Day: 0, Input: &data}},
		{input: puzzles.PuzzleInput{Year: 2015, Day: 26, Input: &data}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%02d", tt.input.Year, tt.input.Day), func(t *testing.T) {
			_, err := solver.Solve(&tt.input, t.Context())
			if err == nil {
				t.Fatalf("Solve(%v#) did not return an error", tt.input)
			}
		})
	}
}
