# DevHatch Project Roadmap

This document outlines the remaining work to complete DevHatch. Items are organized by milestone and priority.

## Current Status

**Foundation Complete** ✓
- CLI framework and argument parsing
- Containerized development environment
- Task runner (build, test, lint, vulnerability scanning)
- VS Code GUI forwarding support (Linux, Windows/WSL)
- Comprehensive contributor documentation

**Next Phase**: Language generators

## Milestone 1: Go Generator (MVP)

**Goal**: Enable `devhatch init go` to generate reproducible Go development environments

### Core Implementation
- [x] Parse `config.yml` configuration file
  - [x] YAML schema definition
  - [x] Schema validation
  - [x] Error handling and reporting
- [ ] Go generator implementation
  - [ ] Read `go.mod` to extract version requirements
  - [ ] Generate/update `Dockerfile` with correct Go version
  - [ ] Generate/update `compose.yml` with Go-specific config
  - [ ] Handle existing files (validate before overwriting)
  - [ ] Add Go-specific build tasks
- [ ] CLI integration
  - [ ] Implement `devhatch init go` command
  - [ ] Interactive setup (optional)
  - [ ] Validation and error messages

### Configuration Schema
```yaml
version: 1
language: go
golang:
  version: "1.27.4"  # From go.mod or manual override
  modules:
    - path: ./cmd/app
      output: bin/app
  cgo_enabled: false  # Optional
  cross_compile:      # Optional
    - goos: windows
      goarch: amd64
```

### Testing
- [x] Unit tests for configuration parsing
- [ ] Unit tests for generator logic
- [ ] Integration tests with sample Go projects
- [ ] Cross-platform testing (Linux, Windows/WSL, macOS)

### Documentation
- [ ] Update README with `devhatch init go` example
- [x] Document `config.yml` schema
- [ ] Add troubleshooting guide for Go projects

## Milestone 2: Additional Language Support

### Node.js Generator
- [ ] Configuration schema for Node.js versions
- [ ] Parse `package.json` for version detection
- [ ] Generate Node-specific Docker config
- [ ] Support npm, yarn, pnpm
- [ ] Integration tests
- [ ] Documentation

### Python Generator
- [ ] Configuration schema for Python versions
- [ ] Parse `pyproject.toml` or `requirements.txt`
- [ ] Generate Python Docker config
- [ ] Virtual environment setup
- [ ] Integration tests
- [ ] Documentation

### Rust Generator
- [ ] Configuration schema for Rust versions
- [ ] Parse `Cargo.toml` for version requirements
- [ ] Generate Rust Docker config
- [ ] Support different Rust editions
- [ ] Integration tests
- [ ] Documentation

## Milestone 3: User Experience Improvements

### Project Initialization
- [ ] Interactive `devhatch init` wizard
- [ ] Detect language from repository contents
- [ ] Suggest appropriate configuration
- [ ] Validate project structure
- [ ] Generate starter configuration

### Configuration Management
- [ ] `devhatch validate` command
- [ ] `devhatch update` to refresh versions
- [ ] `devhatch status` to show environment info
- [ ] Configuration migration for version updates

### Enhanced Error Handling
- [ ] Better error messages with suggestions
- [ ] Validation error details
- [ ] Platform-specific troubleshooting hints
- [ ] Debug mode for detailed logging

## Milestone 4: CI/CD Integration

### GitHub Actions
- [ ] DevHatch development action
- [ ] CI configuration examples
- [ ] Integration tests in workflows
- [ ] Artifact caching strategies

### GitLab CI
- [ ] `.gitlab-ci.yml` examples
- [ ] Integration documentation

### Local CI
- [ ] Pre-push hooks
- [x] Pre-commit integration
- [ ] Local validation commands

## Milestone 5: Advanced Features

### Environment Persistence
- [ ] Save/restore development state
- [ ] Project-specific settings
- [ ] Team configuration sharing
- [ ] Version lock files

### Templates and Scaffolding
- [ ] Language-specific project templates
- [ ] Starter applications
- [ ] Example configurations
- [ ] Generated documentation

### Remote Execution
- [ ] Development in cloud containers
- [ ] Remote debugging support
- [ ] Collaborative development features

### macOS GUI Support
- [ ] Implement X11/Wayland forwarding on macOS
- [ ] Test with XQuartz or native solutions
- [ ] Document macOS-specific setup

## Infrastructure & Quality

### Testing Infrastructure
- [ ] Multi-platform CI/CD pipeline
  - [ ] Linux testing
  - [ ] Windows/WSL testing
  - [ ] macOS testing
- [ ] Test coverage targets (>80% for new code)
- [ ] Integration test suite
- [ ] E2E testing for real projects

### Documentation
- [ ] API documentation
- [ ] CLI documentation (auto-generated)
- [ ] Architecture decisions (ADR format)
- [ ] Migration guides between versions
- [ ] FAQ for common issues
- [ ] Video tutorials (optional)

### Performance
- [ ] Image size optimization
- [ ] Build time optimization
- [ ] Container startup time
- [ ] Benchmarking suite

### Security
- [ ] Dependency vulnerability scanning (automated)
- [ ] Security audit process
- [ ] Signed releases
- [ ] Security advisory handling
- [ ] Container image scanning

## Community & Release

### Release Process
- [ ] Semantic versioning adoption
- [ ] Release notes generation
- [ ] Changelog maintenance
- [ ] GitHub releases with artifacts
- [ ] Package manager support (Homebrew, etc.)

### Community
- [ ] Issue templates
- [ ] Discussion categories
- [ ] Contribution guidelines (in progress)
- [ ] Code review process
- [ ] Contributor recognition
- [ ] Code of conduct

## Deferred/Future Considerations

These items are valuable but intentionally deferred:

- **IDE Integrations**: JetBrains plugins, VS Code extension
- **GUI Application**: Web-based interface for configuration
- **Package Registry**: Custom package/template hosting
- **Analytics**: Optional usage telemetry
- **Commercial Features**: Paid plans for enterprise features
- **Desktop Application**: Electron/Tauri-based GUI

## Success Criteria

### MVP Completion (Go Generator)
- ✓ `devhatch init go` creates working Go environment
- ✓ All tests pass on Linux, Windows/WSL
- ✓ Documentation covers all features
- ✓ No security vulnerabilities
- ✓ <30 second container startup time

### Full Release (Milestone 3+)
- Multiple language support (Go, Node.js, Python, Rust)
- Interactive setup wizard
- <5 minute first-time user experience
- Comprehensive documentation
- Active community with contributions

## Contributing to This Roadmap

We welcome feedback and contributions! To help:

1. **Report Issues**: Let us know what's missing or broken
2. **Suggest Features**: Open a GitHub Discussion
3. **Contribute Code**: See [CONTRIBUTING.md](docs/CONTRIBUTING.md)
4. **Improve Docs**: Documentation improvements are always welcome

## Questions?

- **GitHub Issues**: Report bugs or request features
- **GitHub Discussions**: Ask questions and discuss ideas
- **Contributing**: See [CONTRIBUTING.md](docs/CONTRIBUTING.md) for development guidelines

---

**Last Updated**: November 18, 2024
**Next Review**: After Go generator completion
