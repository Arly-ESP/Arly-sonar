#!/bin/bash
set -e

# Print usage/help message
usage() {
  cat <<EOF
Usage: $(basename "$0") [options]

Options:
  -h, --help         Show this help message.
  -t, --tests        Run tests with coverage (default if no options provided).
  -c, --coverage     Run tests, save coverage to coverage.out, and open the HTML report.
  -b, --bench        Run benchmarks (go test ./... -bench=.).
  -m, --benchmem     Run benchmarks with memory allocation details (go test ./... -bench=. -benchmem).

Examples:
  $(basename "$0")
      Runs tests with coverage (default action).

  $(basename "$0") --coverage
      Runs tests, generates a coverage profile, and opens the coverage report.

  $(basename "$0") --bench --benchmem
      Runs benchmarks with memory allocation details.
EOF
}

RUN_TESTS=false
RUN_COVERAGE=false
RUN_BENCH=false
RUN_BENCHMEM=false

if [ $# -eq 0 ]; then
  RUN_TESTS=true
fi

while [[ "$#" -gt 0 ]]; do
  case "$1" in
    -h|--help)
      usage
      exit 0
      ;;
    -t|--tests)
      RUN_TESTS=true
      ;;
    -c|--coverage)
      RUN_COVERAGE=true
      ;;
    -b|--bench)
      RUN_BENCH=true
      ;;
    -m|--benchmem)
      RUN_BENCHMEM=true
      ;;
    *)
      echo "Unknown option: $1"
      usage
      exit 1
      ;;
  esac
  shift
done

if [ -f "go.mod" ]; then
  echo "[LOG] Already in project root."
else
  echo "[LOG] Changing directory to parent..."
  cd ..
fi

if $RUN_TESTS; then
  echo "[LOG] Running tests with coverage..."
  go test ./... -cover
fi

if $RUN_COVERAGE; then
  echo "[LOG] Running tests and generating coverage profile..."
  go test ./... -coverprofile=coverage.out
  echo "[LOG] Opening coverage report in browser..."
  go tool cover -html=coverage.out
fi

if $RUN_BENCH; then
  echo "[LOG] Running benchmarks..."
  go test ./... -bench=.
fi

if $RUN_BENCHMEM; then
  echo "[LOG] Running benchmarks with memory allocation details..."
  go test ./... -bench=. -benchmem
fi

echo "[LOG] All selected tasks completed."
