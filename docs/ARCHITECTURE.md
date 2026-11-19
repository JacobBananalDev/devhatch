# DevHatch Architecture

This document describes the design and architecture of the DevHatch project.

## Vision

DevHatch enables developers to maintain clean, isolated development environments without installing toolchains globally. By leveraging container technology and language-specific generators, it provides:

- **Environment Isolation**: Each project's dependencies run in containers
- **Reproducibility**: Consistent environments across team members and CI/CD
- **Simplicity**: Minimal configuration, sensible defaults
- **Extensibility**: Generator framework for multiple languages

## High-Level Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                      Host Machine                            │
│  ┌────────────────────────────────────────────────────────┐ │
│  │            DevHatch CLI (devhatch binary)              │ │
│  │  - Parse command-line arguments                        │ │
│  │  - Validate project configuration                      │ │
│  │  - Invoke appropriate generator                        │ │
│  └────────────────────────────────────────────────────────┘ │
│  ┌────────────────────────────────────────────────────────┐ │
│  │          Docker Container (via Docker Compose)         │ │
│  │  ┌──────────────────────────────────────────────────┐  │ │
│  │  │   Language Toolchain (Go, Node, Python, etc.)   │  │ │
│  │  │  ┌──────────────────────────────────────────┐   │  │ │
│  │  │  │  devhatch.yaml                           │   │  │ │
│  │  │  │  - Version specifications                │   │  │ │
│  │  │  │  - Dependencies                          │   │  │ │
│  │  │  │  - Custom configuration                 │   │  │ │
│  │  │  └──────────────────────────────────────────┘   │  │ │
│  │  └──────────────────────────────────────────────────┘  │ │
│  └────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## Project Components

### 1. CLI Application (`cmd/devhatch/`)

**Purpose**: Command-line interface for DevHatch operations

**Key Responsibilities**:
- Parse and validate command-line arguments
- Load and validate project configuration (`devhatch.yaml`)
- Delegate to appropriate generators
- Provide help and version information

**Current Status**: 
- Help system implemented
- Argument parsing framework in place
- Generator delegation ready
- `devhatch init go` placeholder (not implemented)

### 2. Container Environment

**Purpose**: Isolated development environment for all projects

**Components**:
- **Dockerfile**: Builds the base development image
  - Includes Go toolchain (version from `go.mod`)
  - Includes YSH for task automation
  - Includes VS Code for GUI development
  - Non-root user for security
  - Git configuration for safe directories

- **compose.yml**: Defines container runtime configuration
  - Volume mounts (workspace, X11, Wayland)
  - Environment variables (DISPLAY, XAUTHORITY, etc.)
  - User and hostname mapping
  - Resource limits (can be extended)

- **scripts/shell.sh**: Container launcher
  - Detects host platform (Linux, Windows/WSL, macOS)
  - Extracts Go version from `go.mod`
  - Builds Docker image with correct versions
  - Starts interactive shell with host identification
  - Windows/WSL detection and guidance

### 3. Task Runner (`run.ysh`, `scripts/lib/taskfile.ysh`)

**Purpose**: Automate development tasks (build, test, lint, etc.)

**Available Tasks**:
- `version`: Git describe (version detection)
- `build`: Compile CLI binary
- `fmt`: Format Go code
- `lint`: Run Go vet, Staticcheck, Revive, YSH syntax checks
- `test`: Run Go unit tests
- `vuln`: Scan for vulnerabilities (gosec, govulncheck)
- `pre-commit`: Full pre-commit checks
- `check`: Tests + lint + build (recommended before PR)
- `penv`: Run commands from container with current environment

**Implementation**: YSH (Oil Shell) for portable, maintainable shell scripts

## Configuration

### Project Configuration (devhatch.yaml)

Each project using DevHatch includes a `devhatch.yaml` file (not yet implemented):

```yaml
version: 1
language: go
golang:
  version: "1.27.4"
  modules:
    - path: ./cmd/devhatch
  dependencies:
    - github.com/some/package@latest
```

**Parser** (to be implemented):
- Validate YAML schema
- Extract version constraints
- Pass configuration to generators

## Generator Framework

Generators produce the project-specific Docker setup and configuration.

### Generator Interface (to be implemented)

Each generator implements:
- **Validation**: Check `devhatch.yaml` for language-specific fields
- **Generation**: Create `Dockerfile`, `compose.yml`, or other files
- **Verification**: Ensure generated environment is correct

### Go Generator (`devhatch init go`)

**Responsibilities**:
- Parse Go module configuration
- Identify required toolchain versions
- Generate or update Go-specific Docker configuration
- Set up development tasks for Go projects

**Current Status**: Placeholder (not implemented)

**Future Enhancements**:
- Support for Go workspaces
- Module dependency resolution
- CGO configuration
- Cross-compilation setup

### Planned Generators

- **Node.js**: Node version management, package installation
- **Python**: Python version, virtualenv, dependency management
- **Rust**: Rust toolchain, Cargo configuration
- **Additional languages**: To be determined based on community feedback

## Display Forwarding Architecture

### X11/Wayland Support

For GUI applications (VS Code) in containers:

**Host Display Components**:
- `/tmp/.X11-unix/`: X11 socket for display protocol
- `/run/user/{uid}/`: Wayland runtime directory and auth files
- `DISPLAY`: Environment variable (e.g., `:0`)
- `XAUTHORITY`: Path to X11/Wayland authorization file

**Container Configuration**:
- Volumes mounted read-only for security
- Environment variables passed through
- User UID/GID mapped for file permissions
- Hostname included in prompt for context

**Flow**:
1. `scripts/shell.sh` detects host DISPLAY and XAUTHORITY
2. Exports them to `compose.yml` environment
3. Container mounts X11/Wayland resources
4. GUI applications connect to host display
5. VS Code renders on host screen

**Security Considerations**:
- Display access is **read-only** for X11 socket
- Auth files mounted **read-only**
- Container runs as **host user** (not root)
- Use only with **trusted repository content**

## Security Architecture

### Layered Security Approach

1. **Container Isolation**: Language toolchains isolated from host
2. **Non-Root User**: Container runs as host UID/GID (not root)
3. **Read-Only Mounts**: Display resources mounted read-only
4. **Git Configuration**: Safe directory setup in container
5. **Dependency Scanning**: Regular vulnerability checks

### Vulnerability Management

- **gosec**: Go-specific security scanner
- **govulncheck**: Dependency vulnerability detection
- **Regular audits**: Planned security review process

## Development Workflow

### Contributor Setup

```
1. Clone repository
   └─ Fork and clone to local machine

2. Start container
   └─ ./scripts/shell.sh
   └─ Builds image, starts shell

3. Make changes
   └─ Edit code
   └─ Run tests: ./run.ysh test
   └─ Lint: ./run.ysh lint

4. Pre-commit checks
   └─ ./run.ysh pre-commit
   └─ Verifies code quality before push

5. Create PR
   └─ Push to feature branch
   └─ Create PR on GitHub
   └─ CI runs ./run.ysh check

6. Code review
   └─ Maintainers review changes
   └─ Approve when satisfied
   └─ Merge to main
```

## Testing Strategy

### Unit Tests

- Located in `*_test.go` files alongside code
- Table-driven test pattern
- Target: >80% coverage for new code

### Integration Tests (Future)

- Test full workflow: config → generation → environment
- Test multiple languages
- Test cross-platform execution

### Manual Testing

- Container startup and shell access
- VS Code GUI forwarding
- Task runner functionality
- Platform-specific scenarios (Windows/WSL, macOS)

## Extension Points

### For Language Generators

New generators can be added by implementing:

1. **Configuration Parser**: Read language-specific config from `devhatch.yaml`
2. **Validation**: Check for required fields and constraints
3. **Code Generation**: Create Docker configuration
4. **CLI Command**: Register with main CLI app

### For Platform Support

Additional platforms supported by enhancing:

1. **scripts/shell.sh**: Platform detection and guidance
2. **compose.yml**: Platform-specific volume mounts or env vars
3. **Dockerfile**: Optional platform-specific tools

## Performance Considerations

### Image Build Optimization

- Cached layers for frequently used components
- Minimal base image (golang:1.27.4)
- Quick rebuild for small changes

### Container Runtime

- Lightweight YSH for fast task execution
- Local volume mounts (no data copying)
- Shared display forwarding (not full X server)

## Future Enhancements

See [TODO.md](../TODO.md) for detailed roadmap.

Key areas for expansion:

1. **Configuration Schema**: Full YAML validation
2. **Multi-Language Support**: Node.js, Python, Rust generators
3. **Templates and Scaffolding**: Project initialization
4. **CI/CD Integration**: GitHub Actions, GitLab CI
5. **Environment Persistence**: Save and load configurations
6. **Remote Execution**: Cloud-based development environments

## Decision Log

### Why Docker Compose?

- Simple, declarative configuration
- Cross-platform support (Linux, Windows, macOS)
- Built-in networking and volume management
- Wide adoption and excellent tooling support

### Why YSH?

- POSIX-compatible shell with modern features
- Portable across platforms
- Lightweight distribution
- Suitable for simple to moderate automation

### Why X11 Forwarding?

- Enables graphical tools (VS Code) in containers
- Secure, read-only access model
- Standard practice in development containers
- Wayland support via auth file mounting

## References

- [Docker Documentation](https://docs.docker.com/)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [Oil Shell (YSH) Documentation](https://www.oilshell.org/)
- [Go Module Documentation](https://golang.org/doc/modules/)
- [X11 Protocol](https://www.x.org/)
- [Wayland Protocol](https://wayland.freedesktop.org/)
