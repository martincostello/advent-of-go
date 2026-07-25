package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func Parse() (int, int, []byte) {

	year := flag.Int("year", time.Now().Local().Year(), "the year of the puzzle to run")
	day := flag.Int("day", time.Now().Local().Day(), "the day of the puzzle to run")

	flag.Parse()

	var input []byte

	if flag.NArg() < 1 {
		puzzle, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading stdin failed: %v\n", err)
			os.Exit(1)
		}

		input = puzzle
	} else {
		path := flag.Arg(0)
		puzzle, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading file %q failed: %v\n", path, err)
			os.Exit(1)
		}
		input = puzzle
	}

	return *year, *day, input
}
