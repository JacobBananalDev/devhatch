# Contributing to DevHatch

Thank you for your interest in contributing to DevHatch! This guide will help you get started.

## Code of Conduct

Please be respectful and constructive in all interactions with the community.

## Getting Started

### Prerequisites

- Git
- Docker Engine (v20.10+)
- Docker Compose v2+
- For Windows: WSL 2 with Ubuntu

### Setting Up Your Development Environment

1. **Fork and Clone**
   ```bash
   git clone https://github.com/YOUR_USERNAME/devhatch.git
   cd devhatch
   ```

2. **Start the Development Container**
   ```bash
   ./scripts/shell.sh
   ```

   This builds the Docker image and starts an interactive shell with:
   - Go 1.27.4 (from `go.mod`)
   - YSH task runner
   - Nano for terminal editing and commit messages
   - VS Code (for GUI development)
   - All project dependencies

3. **Verify Your Setup**
   ```bash
   ./run.ysh check
   ```

   This runs tests, linting, and builds the project. All should pass.

4. **Enable the Git Pre-Commit Hook**
   ```bash
   git config core.hooksPath .githooks
   git config core.editor nano
   ```

   These are one-time, local Git settings. The hook runs formatting checks, lint,
   vulnerability scans, and unit tests before each commit. Commit inside the rebuilt
   development container so Nano, YSH, and Go are available.

## Development Workflow

### Common Tasks

Inside the container shell, use the task runner to manage your work:

```bash
./run.ysh list-tasks          # Show all available tasks
./run.ysh check               # Run tests, lint, and build (recommended before pushing)
./run.ysh unit-test           # Run Go tests
./run.ysh fmt                 # Format Go code
./run.ysh lint                # Run Go vet, Staticcheck, Revive, and YSH syntax checks
./run.ysh vuln                # Scan for security vulnerabilities
./run.ysh build               # Build the binary to bin/devhatch
```

The configured Git hook runs the pre-commit checks automatically.

### Opening in VS Code

The shell launcher installs the **Dev Containers** extension in host VS Code when the `code` command is available. Alternatively, install it manually. Open the repository in host VS Code and run **Dev Containers: Reopen in Container** from the Command Palette. Your host VS Code profile and sign-in remain available, while workspace tools run in the container.

Alternatively, `./run.ysh penv code .` in the container shell launches a separate container VS Code profile. It does not share the host sign-in and requires X11 or WSLg.

## Making Changes

### File Organization

```
devhatch/
├── cmd/
│   └── devhatch/          # CLI entry point
│       └── main.go
├── docs/                  # Project documentation
│   ├── CONTRIBUTING.md    # This file
│   ├── ARCHITECTURE.md    # Technical design
│   └── TODO.md           # Roadmap and remaining work
├── scripts/
│   ├── shell.sh          # Container launcher
│   └── lib/
│       └── taskfile.ysh  # Task definitions
├── Dockerfile            # Container image definition
├── compose.yml           # Docker Compose configuration
├── go.mod               # Go module declaration
├── run.ysh              # Task runner entry point
└── README.md            # Project overview
```

### Coding Standards

- **Go**: Follow [effective Go](https://golang.org/doc/effective_go) conventions
  - Use `gofmt` for formatting (run `./run.ysh fmt`)
  - Run `go vet ./...` before committing
  - Aim for >80% test coverage for new code

- **Shell Scripts**: Use POSIX sh when possible
  - Run `sh -n scripts/shell.sh` to check syntax
  - Avoid bash-specific features

- **YSH**: Follow the [YSH language guide](https://www.oilshell.org/release/latest/doc/index.html)
  - Run YSH syntax checks with `./run.ysh lint`

### Commits

- Write clear, descriptive commit messages
- Reference issues when applicable (e.g., "Fixes #123")
- Keep commits focused and atomic
- Include the co-author trailer for attribution:

  ```
  Commit message here

  Co-authored-by: Assistant <223556219+Copilot@users.noreply.github.com>
  ```

### Pull Requests

1. **Create a Feature Branch**
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. **Make Your Changes**
   - Edit code
   - Write/update tests
   - Update documentation as needed

3. **Verify Everything Works**
   ```bash
   ./run.ysh check    # Run all checks
   git diff --check   # Ensure no trailing whitespace
   ```

4. **Push and Create PR**
   ```bash
   git push origin feature/your-feature-name
   ```
   Then create a pull request on GitHub.

5. **PR Guidelines**
   - Provide a clear description of changes
   - Reference related issues
   - Ensure CI checks pass
   - Request review from maintainers

## Testing

### Running Tests

```bash
# Run all tests
./run.ysh unit-test

# Run tests in a specific package
./run.ysh unit-test ./cmd/devhatch

# Run with verbose output
./run.ysh unit-test -v ./...

# Run with coverage
./run.ysh unit-test -cover ./...
```

### Writing Tests

- Place test files in the same package as the code they test
- Use the `_test.go` suffix (e.g., `main_test.go`)
- Write table-driven tests for multiple scenarios
- Aim for meaningful test coverage

Example:
```go
func TestExample(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {"valid input", "test", "expected", false},
        {"invalid input", "", "", true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := functionUnderTest(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("got error %v, want %v", err != nil, tt.wantErr)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Linting and Formatting

The project uses multiple linters to maintain code quality:

- **gofmt**: Go standard formatter
- **go vet**: Go compiler's linter
- **Staticcheck**: Static analysis tool
- **Revive**: Go linter with customizable rules
- **YSH syntax checker**: For shell scripts

Run all linters:
```bash
./run.ysh lint
```

## Security

- Run the vulnerability scanner before submitting PRs:
  ```bash
  ./run.ysh vuln
  ```

- Keep dependencies up to date
- Review dependencies before updating
- Report security issues privately to maintainers

## Documentation

- Update documentation for user-facing changes
- Add comments for non-obvious code
- Keep README.md current with new features
- Update ARCHITECTURE.md if design decisions change
- Update TODO.md when completing roadmap items

## Project Structure

For more details on the technical architecture, see [ARCHITECTURE.md](ARCHITECTURE.md).

## Common Issues

### Container Won't Start

```bash
# Clean up and rebuild
docker compose down
./scripts/shell.sh
```

### Display Forwarding Issues (VS Code)

Ensure your X11/Wayland session is configured:
```bash
echo $DISPLAY      # Should show something like :0
echo $XAUTHORITY   # Should show a path to auth file
```

If you're on Windows, make sure you're running from an Ubuntu WSL 2 terminal, not PowerShell or Git Bash.

### Permission Issues

The container runs as your host user (UID/GID), so file permissions should match your host system. If you get permission errors:

```bash
# Exit the container and restart it
exit
./scripts/shell.sh
```

## Getting Help

- **Questions**: Open a GitHub Discussion
- **Bugs**: Create an Issue with reproduction steps
- **Ideas**: Start a Discussion to gather feedback
- **Documentation**: Suggest improvements in an Issue

## Reviewing PRs

If you're reviewing another contributor's PR:

- Test the changes locally
- Verify tests pass
- Check code quality and style
- Leave constructive feedback
- Approve when satisfied

## Release Process

(To be documented as the project matures)

## Recognition

Contributors are recognized in:
- Pull request comments
- Release notes
- Project acknowledgments (when implemented)

Thank you for contributing to DevHatch! 🎉
