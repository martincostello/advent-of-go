// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd_test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"

	"github.com/martincostello/advent-of-go/cmd"
)

func TestCmdRun(t *testing.T) {
	tests := []struct {
		year          int
		day           int
		wantSolution1 string
		wantSolution2 string
	}{
		{2015, 1, "232", "1783"},
		{2015, 2, "1598415", "3812909"},
		{2015, 3, "2565", "2639"},
		{2015, 4, "346386", "9958218"},
		{2015, 5, "236", "51"},
	}

	for _, tt := range tests {
		_, file, _, _ := runtime.Caller(0)
		root := filepath.Dir(file)
		inputDir := filepath.Join(
			root,
			"..",
			"input",
		)
		t.Run(fmt.Sprintf("%d-%02d", tt.year, tt.day), func(t *testing.T) {
			input := filepath.Join(
				inputDir,
				fmt.Sprintf("Y%d", tt.year),
				fmt.Sprintf("Day%02d", tt.day),
				"input.txt",
			)
			got := cmd.Run([]string{"--year", strconv.Itoa(tt.year), "--day", strconv.Itoa(tt.day), input})
			if got.Part1 != tt.wantSolution1 || got.Part2 != tt.wantSolution2 {
				t.Errorf("Run(%d, %d) = (%q, %q), want (%q, %q)", tt.year, tt.day, got.Part1, got.Part2, tt.wantSolution1, tt.wantSolution2)
			}
		})
	}
}
