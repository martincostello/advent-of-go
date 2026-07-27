// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Parse parses the command-line flags and input for the application,
// returning the year and day of the puzzle to solve and its input data.
func Parse(args []string) (int, int, []byte, error) {
	flags := flag.NewFlagSet("cmd", flag.ContinueOnError)
	flags.SetOutput(os.Stdout)

	year := flags.Int("year", time.Now().Local().Year(), "the year of the puzzle to run")
	day := flags.Int("day", time.Now().Local().Day(), "the day of the puzzle to run")

	err := flags.Parse(args)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "parsing flags failed: %v\n", err)
		return 0, 0, nil, err
	}

	var input []byte

	if flags.NArg() < 1 {
		err = fmt.Errorf("no input file specified")
		return 0, 0, nil, err
	} else {
		path := flags.Arg(0)
		input, err = os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading file %q failed: %v\n", path, err)
			return 0, 0, nil, err
		}
	}

	return *year, *day, input, nil
}
