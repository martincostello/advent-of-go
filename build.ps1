#! /usr/bin/env pwsh

$ErrorActionPreference = "Stop"

go test -v ./...

if ($LASTEXITCODE -ne 0) {
    throw "go test failed"
}
