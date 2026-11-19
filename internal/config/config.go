// Package config parses and validates DevHatch project configuration files.
package config

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// FileName is the name of a DevHatch project configuration file.
const FileName = "config.yml"

// Config is the root of a DevHatch project configuration file.
type Config struct {
	Version  int       `yaml:"version"`
	Language string    `yaml:"language"`
	Golang   *GoConfig `yaml:"golang"`
}

// GoConfig contains settings used by the Go project generator.
type GoConfig struct {
	Version      string         `yaml:"version"`
	Modules      []Module       `yaml:"modules"`
	Dependencies []string       `yaml:"dependencies"`
	CGOEnabled   bool           `yaml:"cgo_enabled"`
	CrossCompile []CrossCompile `yaml:"cross_compile"`
}

// Module describes a Go package to build.
type Module struct {
	Path   string `yaml:"path"`
	Output string `yaml:"output"`
}

// CrossCompile describes a Go target platform.
type CrossCompile struct {
	GOOS   string `yaml:"goos"`
	GOARCH string `yaml:"goarch"`
}

// Load decodes and validates a DevHatch YAML configuration.
func Load(r io.Reader) (Config, error) {
	decoder := yaml.NewDecoder(r)
	decoder.KnownFields(true)

	var cfg Config
	if err := decoder.Decode(&cfg); err != nil {
		if errors.Is(err, io.EOF) {
			return Config{}, errors.New("configuration is empty")
		}
		return Config{}, fmt.Errorf("decode configuration: %w", err)
	}

	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		if err == nil {
			return Config{}, errors.New("configuration must contain only one YAML document")
		}
		return Config{}, fmt.Errorf("decode configuration: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

// LoadFile reads and validates config.yml from the current directory.
func LoadFile() (Config, error) {
	data, err := os.ReadFile(FileName)
	if err != nil {
		return Config{}, fmt.Errorf("open configuration %q: %w", FileName, err)
	}

	cfg, err := Load(bytes.NewReader(data))
	if err != nil {
		return Config{}, fmt.Errorf("load configuration %q: %w", FileName, err)
	}
	return cfg, nil
}

// Validate checks the supported configuration schema and its required fields.
func (cfg Config) Validate() error {
	if cfg.Version != 1 {
		return fmt.Errorf("version must be 1, got %d", cfg.Version)
	}
	if strings.TrimSpace(cfg.Language) == "" {
		return errors.New("language is required")
	}
	if cfg.Language != "go" {
		return fmt.Errorf("unsupported language %q: only \"go\" is currently supported", cfg.Language)
	}
	if cfg.Golang == nil {
		return errors.New("golang configuration is required when language is \"go\"")
	}

	for i, module := range cfg.Golang.Modules {
		if strings.TrimSpace(module.Path) == "" {
			return fmt.Errorf("golang.modules[%d].path is required", i)
		}
	}
	for i, dependency := range cfg.Golang.Dependencies {
		if strings.TrimSpace(dependency) == "" {
			return fmt.Errorf("golang.dependencies[%d] must not be empty", i)
		}
	}
	for i, target := range cfg.Golang.CrossCompile {
		if strings.TrimSpace(target.GOOS) == "" {
			return fmt.Errorf("golang.cross_compile[%d].goos is required", i)
		}
		if strings.TrimSpace(target.GOARCH) == "" {
			return fmt.Errorf("golang.cross_compile[%d].goarch is required", i)
		}
	}
	return nil
}
