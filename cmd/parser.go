// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

// Parse parses the command-line flags and input for the application,
// returning the year and day of the puzzle to solve and its input data.
func parse(args []string) (int, int, []byte) {
	flags := flag.NewFlagSet("cmd", flag.ContinueOnError)
	flags.SetOutput(io.Discard)

	year := flags.Int("year", time.Now().Local().Year(), "the year of the puzzle to run")
	day := flags.Int("day", time.Now().Local().Day(), "the day of the puzzle to run")

	err := flags.Parse(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing flags failed: %v\n", err)
		os.Exit(1)
	}

	var input []byte

	if flags.NArg() < 1 {
		puzzle, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading stdin failed: %v\n", err)
			os.Exit(1)
		}

		input = puzzle
	} else {
		path := flags.Arg(0)
		puzzle, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading file %q failed: %v\n", path, err)
			os.Exit(1)
		}
		input = puzzle
	}

	return *year, *day, input
}
