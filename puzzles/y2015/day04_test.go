// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package y2015_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/martincostello/advent-of-go/puzzles/y2015"
)

func TestY2015GetLowestPositiveNumberHash(t *testing.T) {
	t.Parallel()
	tests := []struct {
		secretKey string
		zeroes    int
		want      int
	}{
		{"abcdef", 5, 609043},
		{"pqrstuv", 5, 1048970},
	}

	for _, tt := range tests {
		c := tt
		t.Run(c.secretKey, func(t *testing.T) {
			t.Parallel()
			got, err := y2015.GetLowestPositiveNumberHash(t.Context(), c.secretKey, c.zeroes)
			require.NoError(t, err, "GetLowestPositiveNumberHash(%q, %d) returned error: %v", c.secretKey, c.zeroes, err)
			require.Equal(t, c.want, got, "GetLowestPositiveNumberHash(%q, %d) = %d, want %d", c.secretKey, c.zeroes, got, c.want)
		})
	}
}

func TestY2015Day04IfNoSolution(t *testing.T) {
	t.Parallel()
	input := "invalid"
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err := y2015.Day04(ctx, input)
	require.Error(t, err, "Day04(%q) did not return an error", input)
}

func TestY2015GetLowestPositiveNumberHashStopsPromptlyOnCancellation(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 200*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	var err error

	go func() {
		defer close(done)
		_, err = y2015.GetLowestPositiveNumberHash(ctx, "abcdef", 8)
	}()

	select {
	case <-done:
		require.ErrorIs(t, err, context.DeadlineExceeded, "GetLowestPositiveNumberHash did not return the context's error")
	case <-time.After(5 * time.Second):
		t.Fatal("GetLowestPositiveNumberHash did not stop promptly after the context was cancelled")
	}
}
