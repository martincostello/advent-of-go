// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015Day03(t *testing.T) {
	tests := []struct {
		input     string
		wantPart1 string
		wantPart2 string
	}{
		{">", "1", "2"},
		{"^>v<", "4", "3"},
		{"^v^v^v^v^v", "2", "11"},
		{"^v", "2", "3"},
		{"^>v<", "4", "3"},
		{"^v^v^v^v^v", "2", "11"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := y2015.Day03(tt.input)
			if got.Part1 != tt.wantPart1 || got.Part2 != tt.wantPart2 {
				t.Errorf("Day03(%q) = (%q, %q), want (%q, %q)", tt.input, got.Part1, got.Part2, tt.wantPart1, tt.wantPart2)
			}
		})
	}
}
