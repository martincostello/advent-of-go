// Copyright (c) Martin Costello, 2026. All rights reserved.
// Licensed under the Apache 2.0 license. See the LICENSE file in the project root for full license information.

package main

import (
	"os"

	"github.com/martincostello/advent-of-go/cmd"
)

func main() {
	_ = cmd.Run(os.Args[1:])
}
