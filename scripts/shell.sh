#!/bin/sh

set -eu

case "$(uname -s)" in
    MINGW* | MSYS* | CYGWIN*)
        printf '%s\n' "shell.sh: Windows shell detected outside WSL." >&2
        printf '%s\n' "Open an Ubuntu WSL 2 terminal and run this script there." >&2
        printf '%s\n' "If WSL is not installed, run 'wsl --install -d Ubuntu' from PowerShell as Administrator." >&2
        printf '%s\n' "Install Docker Desktop and enable its WSL integration for Ubuntu." >&2
        exit 1
        ;;
esac

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

HOST_USER=$(id -un)
HOST_NAME=$(hostname)
DEV_UID=$(id -u)
DEV_GID=$(id -g)
HOST_HOSTNAME=$HOST_NAME
XAUTHORITY_HOST=${XAUTHORITY:-/dev/null}
PROMPT_PATH=$PROJECT_ROOT

case "$PROJECT_ROOT" in
    "$HOME")
        PROMPT_PATH='~'
        ;;
    "$HOME"/*)
        PROMPT_PATH="~${PROJECT_ROOT#"$HOME"}"
        ;;
esac

PROMPT="${HOST_USER}@${HOST_NAME}:${PROMPT_PATH} [ENV] $ "

export DEV_UID DEV_GID HOST_HOSTNAME XAUTHORITY_HOST

exec docker compose run --build --rm -e "PS1=$PROMPT" dev sh -i
