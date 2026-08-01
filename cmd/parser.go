// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package cmd

import (
	"flag"
	"fmt"
	"io"
	"math"
	"time"
)

type Options struct {
	Year     int
	Day      int
	FileName string
}

// Parse parses the command-line flags and input for the application,
// returning the year and day of the puzzle to solve and its input data.
func Parse(stderr io.Writer, args ...string) (*Options, error) {
	now := time.Now().Local()

	options := &Options{
		Year:     now.Year(),
		Day:      int(math.Min(float64(now.Day()), 25)),
		FileName: "",
	}

	flags := flag.NewFlagSet("aoc", flag.ContinueOnError)
	flags.SetOutput(stderr)

	flags.IntVar(&options.Year, "year", options.Year, "Year of the puzzle")
	flags.IntVar(&options.Day, "day", options.Day, "Day of the puzzle")

	flags.Usage = func() {
		flags.Output()
		fmt.Fprintf(stderr, "Usage: %s [options] <file>\n", flags.Name())
		flags.PrintDefaults()
	}

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	options.FileName = flags.Arg(0)

	if err := validate(options, now.Year()); err != nil {
		_, _ = fmt.Fprintln(stderr, err)
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

	if options.FileName == "" {
		return fmt.Errorf("no input file specified")
	}

	return nil
}
