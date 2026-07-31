#!/usr/bin/env pwsh
param()

$ErrorActionPreference = "Stop"

go test -benchmem -run=^$ -bench "^BenchmarkCmdRun$" "github.com/martincostello/advent-of-go/cmd" -count=5 -memprofile=mem.out

if ($LASTEXITCODE -ne 0) {
    throw "go test -bench failed"
}
