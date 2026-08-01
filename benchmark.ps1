#!/usr/bin/env pwsh
param(
    [Parameter(Mandatory = $false)][int] $Count = 5
)

$ErrorActionPreference = "Stop"

go test -benchmem -run=^$ -bench "^BenchmarkCmdRun$" "github.com/martincostello/advent-of-go/cmd" -count="$Count" -memprofile=mem.out

if ($LASTEXITCODE -ne 0) {
    throw "go test -bench failed"
}
