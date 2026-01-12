#!/bin/bash
# Bash script to run Go tests and generate HTML report
# Requires: go-test-report (will attempt to install if missing)

REPORT_TOOL="go-test-report"
OUTPUT_FILE="test_report.html"

# Check if go-test-report is installed
if ! command -v $REPORT_TOOL &> /dev/null; then
    echo "Tool '$REPORT_TOOL' not found. Installing..."
    go install github.com/vakenbolt/go-test-report@latest
    
    # Add GOPATH/bin to PATH for this session if not already there
    export PATH=$PATH:$(go env GOPATH)/bin
    
    # Check again
    if ! command -v $REPORT_TOOL &> /dev/null; then
        echo "Error: Failed to install or find $REPORT_TOOL. Please install manually:"
        echo "go install github.com/vakenbolt/go-test-report@latest"
        exit 1
    fi
fi

echo "Running tests and generating HTML report..."
echo "Output will be saved to: $OUTPUT_FILE"

# Run tests with JSON output and pipe to report tool
go test -json ./pkg/test/... | $REPORT_TOOL -o $OUTPUT_FILE

EXIT_CODE=$?

# go test returns 1 on fail, but we still want the report.
# If piping fails completely, it might be different, but generally this catches the test run result.

if [ -f "$OUTPUT_FILE" ]; then
    echo "Report generated successfully: $OUTPUT_FILE"
else
    echo "Error: Failed to generate report."
    exit 1
fi

# Exit with the test exit code so CI pipelines fail if tests fail
exit $EXIT_CODE
