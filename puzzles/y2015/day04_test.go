// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"testing"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015GetLowestPositiveNumberHash(t *testing.T) {
	tests := []struct {
		secretKey string
		zeroes    int
		want      int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}

	for _, tt := range tests {
		t.Run(tt.secretKey, func(t *testing.T) {
			got := y2015.GetLowestPositiveNumberHash(tt.secretKey, tt.zeroes)
			if got != tt.want {
				t.Errorf("GetLowestPositiveNumberHash(%q, %d) = %d, want %d", tt.secretKey, tt.zeroes, got, tt.want)
			}
		})
	}
}
