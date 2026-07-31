#!/usr/bin/env pwsh
param(
    [Parameter(Mandatory = $true)][int] $Day,
    [Parameter(Mandatory = $true)][int] $Year,
    [Parameter(Mandatory = $false)][string] $InputFile = $null
)

$ErrorActionPreference = "Stop"

if ([string]::IsNullOrEmpty($InputFile)) {
    $InputFile = Join-Path $PSScriptRoot "input" "Y${Year}" ("Day{0:D2}" -f $Day) "input.txt"
}

go run . --day $Day --year $Year $InputFile

if ($LASTEXITCODE -ne 0) {
    throw "go run failed"
}
