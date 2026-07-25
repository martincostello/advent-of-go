package y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles"
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
			input := puzzles.PuzzleData([]byte(tt.input))
			got := y2015.Day01(&input)
			if got.Part1 != tt.wantFloor || got.Part2 != tt.wantInstruction {
				t.Errorf("Day01(%q) = (%q, %q), want (%q, %q)", tt.input, got.Part1, got.Part2, tt.wantFloor, tt.wantInstruction)
			}
		})
	}
}
