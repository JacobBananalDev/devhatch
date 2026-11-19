# Development Handoff

## Current branch

`feature/go-mvp-foundation`

## Project direction

The current focus is the contributor workflow for developing DevHatch itself, not the files that `devhatch init go` will eventually generate.

The intended contributor flow is shell-first:

1. Clone the repository and enter it.
2. Run `./scripts/shell.sh` to enter a Docker environment using the Go version in the checked-out branch's `go.mod`.
3. From that environment, run `./run.ysh penv code .` to open VS Code.

The VS Code handoff is not implemented yet. YSH is not installed in the container, and `run.ysh` does not exist.

## Implemented so far

- `go.mod` declares module `github.com/JacobBananalDev/devhatch` and Go `1.27.4`.
- The root `Dockerfile` accepts `GO_VERSION` as a build argument.
- `compose.yml` passes `GO_VERSION` to the Docker build.
- `scripts/shell.sh` reads the Go version from `go.mod`, exports it, and starts an interactive Compose shell with an `[env]` prompt.
- `cmd/devhatch/main.go` provides help and validates `init go`; generation intentionally reports that it is not implemented yet.

## Decisions

- Use `scripts/shell.sh` as the contributor shell launcher.
- Use Docker directly in scripts; do not invoke `sudo` from the scripts.
- Refuse to overwrite generated files when the generator is implemented.
- Keep the shell-first workflow; a portable way for the container to open host VS Code still needs investigation.
- Do not install Go or YSH globally.

## Validation status

- Docker image build with `GO_VERSION=1.27.4` works.
- CLI help and invalid/unimplemented command cases were exercised in Docker.
- `gofmt -d` is clean and `go test ./...` passes; no Go tests exist yet.
- `sh -n scripts/shell.sh` passes. The launcher itself has not yet been run end-to-end.
- On Ubuntu, Docker group membership was added. `docker info` worked in a fresh login shell; `newgrp` is unavailable.

## Next task

From the repository root in a shell with direct Docker access, run `./scripts/shell.sh`. Inside the container, verify the `[env]` prompt and that `go version` matches the `go` directive in `go.mod`, then exit. After that, continue with YSH installation and the VS Code handoff.
