#!/bin/sh

set -eu

PROJECT_ROOT=$(CDPATH= cd "$(dirname "$0")/.." && pwd)

if [ ! -f "$PROJECT_ROOT/go.mod" ]; then
    printf '%s\n' "shell.sh: go.mod not found" >&2
    exit 1
fi

GO_VERSION=$(awk '$1 == "go" { print $2; exit }' "$PROJECT_ROOT/go.mod")

if [ -z "$GO_VERSION" ]; then
    printf '%s\n' "shell.sh: Go version not found in go.mod" >&2
    exit 1
fi

export GO_VERSION
cd "$PROJECT_ROOT"

exec docker compose run --build --rm -e 'PS1=[env] $ ' dev sh -i
