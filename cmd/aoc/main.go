// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/martincostello/advent-of-go/cmd"
)

func main() {
	_, err := cmd.Run(&cmd.Environment{
		Args:   os.Args[1:],
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}, context.Background())
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
