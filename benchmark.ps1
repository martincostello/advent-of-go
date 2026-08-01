#!/usr/bin/env pwsh
param(
    [Parameter(Mandatory = $false)][int] $Count = 5
)

$ErrorActionPreference = "Stop"

Push-Location $PSScriptRoot

try {
    go test -benchmem -run=^$ -bench "^BenchmarkCmdRun$" ./cmd -count="$Count" -memprofile=mem.out
}
finally {
    Pop-Location
}

if ($LASTEXITCODE -ne 0) {
    throw "go test -bench failed"
}
