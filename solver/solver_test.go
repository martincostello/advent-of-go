// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package solver_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles"
	"github.com/martincostello/advent-of-go/solver"
)

func TestSolverSolveWhenUnsolved(t *testing.T) {
	input := puzzles.PuzzleInput{
		Year:  2014,
		Day:   42,
		Input: nil,
	}
	_, err := solver.Solve(&input, t.Context())
	if err == nil {
		t.Fatalf("Solve(%v#) did not return an error", input)
	}
}
