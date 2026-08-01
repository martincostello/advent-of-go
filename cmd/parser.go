// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

type Options struct {
	Year  int
	Day   int
	Input string
}

// Parse parses the command-line flags and input for the application,
// returning the year and day of the puzzle to solve and its input data.
func Parse(args []string) (*Options, error) {
	now := time.Now().Local()

	options := &Options{
		Year:  now.Year(),
		Day:   int(math.Min(float64(now.Day()), 25)),
		Input: "",
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

	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(0)
		}
		return nil, err
	}

	options.Input = flags.Arg(0)

	if err := validate(options, now.Year()); err != nil {
		fmt.Fprintln(flags.Output(), err)
		flags.Usage()
		return nil, err
	}

	return options, nil
}

func validate(options *Options, year int) error {
	if options.Year < 2015 || options.Year > year {
		return fmt.Errorf("invalid year: %d", options.Year)
	}

	if options.Day < 1 || options.Day > 25 {
		return fmt.Errorf("invalid day: %d", options.Day)
	}

	if options.Input == "" {
		return fmt.Errorf("no input file specified")
	}

	return nil
}
