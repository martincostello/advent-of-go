// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

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
			got, err := y2015.GetLowestPositiveNumberHash(tt.secretKey, tt.zeroes, t.Context())
			require.NoError(t, err, "GetLowestPositiveNumberHash(%q, %d) returned error: %v", tt.secretKey, tt.zeroes, err)
			require.Equal(t, tt.want, got, "GetLowestPositiveNumberHash(%q, %d) = %d, want %d", tt.secretKey, tt.zeroes, got, tt.want)
		})
	}
}

func TestY2015Day04IfNoSolution(t *testing.T) {
	input := "invalid"
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err := y2015.Day04(input, ctx)
	require.Error(t, err, "Day04(%q) did not return an error", input)
}
