# PowerShell Script to run Go tests and generate HTML report
# Requires: go-test-report (will attempt to install if missing)

$reportTool = "go-test-report"
$outputFile = "test_report.html"

# Check if go-test-report is installed
if (-not (Get-Command $reportTool -ErrorAction SilentlyContinue)) {
    Write-Host "Tool '$reportTool' not found. Installing..."
    go install github.com/vakenbolt/go-test-report@latest
    
    # Check again
    if (-not (Get-Command $reportTool -ErrorAction SilentlyContinue)) {
        # Try to find it in likely GOPATH/bin locations
        $goPath = go env GOPATH
        $toolPath = Join-Path $goPath "bin" $reportTool
        if (Test-Path "$toolPath.exe") {
            $reportTool = "$toolPath.exe"
        }
        else {
            Write-Error "Failed to install or find $reportTool. Please install manually: go install github.com/vakenbolt/go-test-report@latest"
            exit 1
        }
    }
}

Write-Host "Running tests and generating HTML report..."
Write-Host "Output will be saved to: $outputFile"

# Use cmd /c to handle piping to avoid PowerShell ConstrainedLanguage mode and encoding issues
cmd /c "go test -json ./pkg/test/... | $reportTool -o $outputFile"

if ($LASTEXITCODE -eq 0) {
    Write-Host "Report generated successfully: $outputFile"
    Write-Host "You can open it with: Invoke-Item $outputFile"
}
else {
    Write-Host "Tests finished with exit code $LASTEXITCODE. Check report for details."
}
