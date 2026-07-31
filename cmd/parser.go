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
	flags := flag.NewFlagSet("cmd", flag.ContinueOnError)
	flags.SetOutput(os.Stdout)

	now := time.Now().Local()

	options := &Options{
		Year:  now.Year(),
		Day:   int(math.Max(1, math.Min(float64(now.Day()), 25))),
		Input: nil,
	}

	flags.IntVar(&options.Year, "year", options.Year, "The year of the puzzle to run")
	flags.IntVar(&options.Day, "day", options.Day, "The day of the puzzle to run")

	err := flags.Parse(args)
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "parsing flags failed: %v\n", err)
		return nil, err
	}

	if flags.NArg() < 1 {
		err = fmt.Errorf("no input file specified")
		return nil, err
	} else {
		path := flags.Arg(0)
		options.Input, err = os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading file %q failed: %v\n", path, err)
			return nil, err
		}
	}

	return options, nil
}
