// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd_test

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/martincostello/advent-of-go/cmd"
)

func TestCmdParseWhenValidFlags(t *testing.T) {
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Dir(file)
	input := filepath.Join(
		root,
		"..",
		"input",
		"Y2015",
		"Day01",
		"input.txt",
	)

	args := []string{"--day", "1", "--year", "2015", input}
	options, err := cmd.Parse(args)
	require.NoError(t, err, "Parse(%v) returned an error", args)
	require.Equal(t, 1, options.Day, "Parse(%v) day = %d, want 1", args, options.Day)
	require.Equal(t, 2015, options.Year, "Parse(%v) year = %d, want 2015", args, options.Year)
}

func TestCmdParseWhenInvalidFlag(t *testing.T) {
	tests := []struct {
		args []string
	}{
		{args: []string{"--invalid", "value"}},
		{args: []string{"--day", "foo"}},
		{args: []string{"--day", "0"}},
		{args: []string{"--day", "26"}},
		{args: []string{"--year", "foo"}},
		{args: []string{"--year", "2014"}},
		{args: []string{"--year", strconv.Itoa(time.Now().Year() + 1)}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			_, err := cmd.Parse(tt.args)
			require.Error(t, err, "Parse(%v) did not return an error", tt.args)
		})
	}
}

func TestCmdParseWhenNoInputFileSpecified(t *testing.T) {
	args := []string{"--day", "1", "--year", "2015"}
	_, err := cmd.Parse(args)
	require.Error(t, err, "Parse(%v) did not return an error", args)
}
