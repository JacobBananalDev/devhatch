# DevHatch

> A lightweight, cross-platform CLI for generating reproducible, containerized development environments without installing language toolchains globally.

DevHatch simplifies setting up isolated development environments by leveraging Docker and language-specific generators. Keep your host system clean while maintaining consistent, reproducible development setups across team members and platforms.

## Features

- **Containerized Development**: All language toolchains run in Docker containers
- **Reproducible Environments**: Check in configuration, everyone gets identical setups
- **Cross-Platform Support**: Works on Linux, Windows (WSL 2), and macOS
- **Integrated IDE Support**: VS Code GUI forwarding from containerized development
- **Multiple Language Support**: Generator framework for Go, Node.js, Python, and more (coming soon)

## Getting Started

### Prerequisites

- **Docker Engine** (v20.10+)
- **Docker Compose** v2 or later
- **Git**

**Platform-specific notes:**
- **Linux**: Standard Docker installation
- **Windows**: WSL 2 with Docker Desktop and WSL integration enabled
- **macOS**: Docker Desktop

### Installation & Development Setup

Clone the repository and start the development container:

```bash
git clone https://github.com/JacobBananalDev/devhatch.git
cd devhatch
./scripts/shell.sh
```

The container starts automatically. You'll see a shell prompt with your host user, hostname, and project path:

```
user@hostname:~/Projects/devhatch [ENV] $
```

### Development Workflow

Inside the container shell, the task runner is available:

```bash
# List all available tasks
./run.ysh list-tasks

# Run the complete verification suite (tests, lint, build)
./run.ysh check

# Format Go code
./run.ysh fmt

# Run linting (Go vet, Staticcheck, Revive, YSH syntax)
./run.ysh lint

# Run tests
./run.ysh test

# Security vulnerability scan
./run.ysh vuln

# Build the binary
./run.ysh build
```

The compiled CLI binary is written to `bin/devhatch`.

### Using VS Code

Launch VS Code from inside the container with the workspace already open:

```bash
./run.ysh penv code .
```

The container automatically forwards your host's X11/Wayland display, giving VS Code access to your screen. This workflow is supported on:
- ✅ Linux with X11/Wayland
- ✅ Windows with WSLg (from Ubuntu WSL 2 terminal)
- ⚠️ macOS (container shell available, GUI forwarding not configured)

### Current Status

The CLI foundation is in development. Currently available:
- Core CLI framework with help system
- Containerized development environment
- Task runner with comprehensive build/test/lint automation

**Coming soon:**
- `devhatch init go` - Go project generator
- Additional language generators (Node.js, Python, Rust)
- Configuration validation and schema
- Project bootstrapping and templates

For detailed development information, see [CONTRIBUTING.md](docs/CONTRIBUTING.md) and [ARCHITECTURE.md](docs/ARCHITECTURE.md).

## Documentation

- **[CONTRIBUTING.md](docs/CONTRIBUTING.md)** - How to contribute to DevHatch
- **[ARCHITECTURE.md](docs/ARCHITECTURE.md)** - Technical architecture and design decisions
- **[TODO.md](TODO.md)** - Project roadmap and remaining work

## License

[Include your license here - e.g., MIT, Apache 2.0, etc.]

## Support

Found a bug or have a suggestion? Please [open an issue](https://github.com/JacobBananalDev/devhatch/issues).

Want to contribute? See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.
