package Y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles/Y2015"
)

func TestSolveDay01(t *testing.T) {
	tests := []struct {
		name            string
		wantFloor       string
		wantInstruction string
	}{
		{"(())", "0", "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Y2015.Day01([]byte(tt.name))
			if got.Part1 != tt.wantFloor || got.Part2 != tt.wantInstruction {
				t.Errorf("Day01(%q) = (%q, %q), want (%q, %q)", tt.name, got.Part1, got.Part2, tt.wantFloor, tt.wantInstruction)
			}
		})
	}
}
