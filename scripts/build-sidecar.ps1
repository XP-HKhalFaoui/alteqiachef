# Build the Go backend sidecar for Tauri bundling
# Must be run before `tauri build` or `tauri dev`

$ErrorActionPreference = "Stop"

# Detect the Rust target triple (e.g. x86_64-pc-windows-msvc)
$target = (rustc -vV | Select-String "host:").ToString().Trim().Split()[-1]
Write-Host "Building pos-backend for target: $target"

$root = Split-Path $PSScriptRoot -Parent
$outDir = Join-Path $root "frontend\src-tauri\binaries"

if (-not (Test-Path $outDir)) {
    New-Item -ItemType Directory -Force $outDir | Out-Null
}

$outFile = Join-Path $outDir "pos-backend-$target.exe"

Push-Location (Join-Path $root "backend")
try {
    go build -ldflags="-s -w" -o $outFile .
    Write-Host "Built: $outFile"
} finally {
    Pop-Location
}
