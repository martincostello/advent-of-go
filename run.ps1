#! /usr/bin/env pwsh
param(
    [Parameter(Mandatory = $true)][int] $Day,
    [Parameter(Mandatory = $true)][int] $Year,
    [Parameter(Mandatory = $true)][string] $InputFile
)

$ErrorActionPreference = "Stop"

go run . --day $Day --year $Year $InputFile

if ($LASTEXITCODE -ne 0) {
    throw "go run failed"
}
