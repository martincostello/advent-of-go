#! /usr/bin/env pwsh
param()

$ErrorActionPreference = "Stop"

go test -benchmem -run=^$ -bench "^BenchmarkCmdRun$" "github.com/martincostello/advent-of-go/cmd" -memprofile=mem.out -run=' '

if ($LASTEXITCODE -ne 0) {
    throw "go test -bench failed"
}
