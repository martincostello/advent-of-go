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
	t.Parallel()
	_, file, _, _ := runtime.Caller(0)

	var (
		root  = filepath.Dir(file)
		input = filepath.Join(
			root,
			"..",
			"input",
			"Y2015",
			"Day01",
			"input.txt",
		)
		env = newEnv("--day", "1", "--year", "2015", input)
	)

	options, err := cmd.Parse(nil, env.Args...)
	require.NoError(t, err, "Parse(%v) returned an error", env.Args)
	require.Equal(t, 1, options.Day, "Parse(%v) day = %d, want 1", env.Args, options.Day)
	require.Equal(t, 2015, options.Year, "Parse(%v) year = %d, want 2015", env.Args, options.Year)
}

func TestCmdParseWhenInvalidFlag(t *testing.T) {
	t.Parallel()
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
		args := tt.args
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			t.Parallel()
			env := newEnv(args...)
			_, err := cmd.Parse(env.Stderr, env.Args...)
			require.Error(t, err, "Parse(%v) did not return an error", env.Args)
		})
	}
}

func TestCmdParseWhenNoInputFileSpecified(t *testing.T) {
	t.Parallel()
	env := newEnv("--day", "1", "--year", "2015")
	_, err := cmd.Parse(env.Stderr, env.Args...)
	require.Error(t, err, "Parse(%v) did not return an error", env.Args)
}

func TestCmdParseHelpSpecified(t *testing.T) {
	t.Parallel()
	env := newEnv("--help")
	_, err := cmd.Parse(env.Stderr, env.Args...)
	require.Error(t, err, "Parse(%v) did not return an error", env.Args)
}
