#!/usr/bin/env pwsh
param(
    [Parameter(Mandatory = $false)][int] $Count = 5
)

$ErrorActionPreference = "Stop"

$path = Join-Path $PSScriptRoot "cmd"

go test -benchmem -run=^$ -bench "^BenchmarkCmdRun$" $path -count="$Count" -memprofile=mem.out

if ($LASTEXITCODE -ne 0) {
    throw "go test -bench failed"
}
