#! /usr/bin/env pwsh

$ErrorActionPreference = "Stop"

go test ./... -coverpkg=./... -race

if ($LASTEXITCODE -ne 0) {
    throw "go test failed"
}
