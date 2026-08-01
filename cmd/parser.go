// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

type Options struct {
	Year  int
	Day   int
	Input []byte
}

// Parse parses the command-line flags and input for the application,
// returning the year and day of the puzzle to solve and its input data.
func Parse(args []string) (*Options, error) {
	now := time.Now().Local()

	options := &Options{
		Year:  now.Year(),
		Day:   int(math.Max(1, math.Min(float64(now.Day()), 25))),
		Input: nil,
	}

	flags := flag.NewFlagSet("advent-of-go", flag.ContinueOnError)
	flags.SetOutput(os.Stdout)

	flags.IntVar(&options.Year, "year", options.Year, "Year of the puzzle")
	flags.IntVar(&options.Day, "day", options.Day, "Day of the puzzle")

	flags.Usage = func() {
		flags.Output()
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <file>\n", flags.Name())
		flags.PrintDefaults()
	}

	err := flags.Parse(args)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		return nil, err
	}

	path := flags.Arg(0)

	if path == "" {
		return nil, fmt.Errorf("no input file specified")
	}

	options.Input, err = os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("reading file %q failed: %v\n", path, err)
	}

	return options, nil
}
