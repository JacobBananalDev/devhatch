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

if ! command -v code >/dev/null 2>&1; then
    printf '%s\n' "shell.sh: VS Code's 'code' command was not found; install Dev Containers manually if needed." >&2
elif ! INSTALLED_EXTENSIONS=$(code --list-extensions); then
    printf '%s\n' "shell.sh: Could not list VS Code extensions." >&2
    exit 1
elif ! printf '%s\n' "$INSTALLED_EXTENSIONS" | grep -Fxiq 'ms-vscode-remote.remote-containers'; then
    printf '%s\n' "Installing the VS Code Dev Containers extension..."
    if ! code --install-extension ms-vscode-remote.remote-containers; then
        printf '%s\n' "shell.sh: Could not install the VS Code Dev Containers extension." >&2
        exit 1
    fi
fi

HOST_USER=$(id -un)
HOST_NAME=$(hostname)
DEV_UID=$(id -u)
DEV_GID=$(id -g)
HOST_HOSTNAME=$HOST_NAME
XAUTHORITY_HOST=${XAUTHORITY:-/dev/null}
PROMPT_PATH=$PROJECT_ROOT
WAYLAND_SOCKET=

if [ -n "${WAYLAND_DISPLAY:-}" ] && [ -n "${XDG_RUNTIME_DIR:-}" ]; then
    case "$WAYLAND_DISPLAY" in
        /*)
            WAYLAND_SOCKET=$WAYLAND_DISPLAY
            ;;
        *)
            WAYLAND_SOCKET=$XDG_RUNTIME_DIR/$WAYLAND_DISPLAY
            ;;
    esac
fi

case "$PROJECT_ROOT" in
    "$HOME")
        PROMPT_PATH='~'
        ;;
    "$HOME"/*)
        PROMPT_PATH="~${PROJECT_ROOT#"$HOME"}"
        ;;
esac

if [ -t 1 ] && [ "${TERM:-dumb}" != dumb ]; then
    PROMPT=$(printf '\033[01;32m%s@%s\033[00m:\033[01;34m%s\033[00m [ENV] $ ' "$HOST_USER" "$HOST_NAME" "$PROMPT_PATH")
else
    PROMPT="${HOST_USER}@${HOST_NAME}:${PROMPT_PATH} [ENV] $ "
fi

export DEV_UID DEV_GID HOST_HOSTNAME WAYLAND_SOCKET XAUTHORITY_HOST

exec docker compose run --build --rm -e "PS1=$PROMPT" dev sh -i
