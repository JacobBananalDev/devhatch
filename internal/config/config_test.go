package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadValidConfig(t *testing.T) {
	input := `version: 1
language: go
golang:
  version: "1.27.4"
  modules:
    - path: ./cmd/app
      output: bin/app
  dependencies:
    - example.com/library@v1.2.3
  cgo_enabled: true
  cross_compile:
    - goos: windows
      goarch: amd64
`

	cfg, err := Load(strings.NewReader(input))
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.Version != 1 || cfg.Language != "go" {
		t.Fatalf("Load() returned version %d and language %q", cfg.Version, cfg.Language)
	}
	if cfg.Golang == nil || cfg.Golang.Version != "1.27.4" || !cfg.Golang.CGOEnabled {
		t.Fatalf("Load() returned unexpected golang config: %#v", cfg.Golang)
	}
	if len(cfg.Golang.Modules) != 1 || cfg.Golang.Modules[0].Path != "./cmd/app" {
		t.Fatalf("Load() returned unexpected modules: %#v", cfg.Golang.Modules)
	}
	if len(cfg.Golang.CrossCompile) != 1 || cfg.Golang.CrossCompile[0].GOOS != "windows" {
		t.Fatalf("Load() returned unexpected cross-compile targets: %#v", cfg.Golang.CrossCompile)
	}
}

func TestLoadFile(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd() error = %v", err)
	}

	dir := t.TempDir()
	data := []byte("version: 1\nlanguage: go\ngolang: {}\n")
	if err := os.WriteFile(filepath.Join(dir, FileName), data, 0600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Chdir() error = %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalDir); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	cfg, err := LoadFile()
	if err != nil {
		t.Fatalf("LoadFile() error = %v", err)
	}
	if cfg.Version != 1 || cfg.Language != "go" {
		t.Fatalf("LoadFile() returned unexpected config: %#v", cfg)
	}
}

func TestLoadRejectsInvalidConfigurations(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr string
	}{
		{
			name:    "empty",
			input:   "",
			wantErr: "configuration is empty",
		},
		{
			name: "unsupported schema version",
			input: `version: 2
language: go
golang: {}
`,
			wantErr: "version must be 1",
		},
		{
			name: "missing language",
			input: `version: 1
golang: {}
`,
			wantErr: "language is required",
		},
		{
			name: "missing go configuration",
			input: `version: 1
language: go
`,
			wantErr: "golang configuration is required",
		},
		{
			name: "unsupported language",
			input: `version: 1
language: python
`,
			wantErr: "unsupported language",
		},
		{
			name: "unknown field",
			input: `version: 1
language: go
golang: {}
unknown: true
`,
			wantErr: "field unknown not found",
		},
		{
			name: "missing module path",
			input: `version: 1
language: go
golang:
  modules:
    - output: bin/app
`,
			wantErr: "golang.modules[0].path is required",
		},
		{
			name: "missing cross-compile architecture",
			input: `version: 1
language: go
golang:
  cross_compile:
    - goos: windows
`,
			wantErr: "golang.cross_compile[0].goarch is required",
		},
		{
			name: "multiple documents",
			input: `version: 1
language: go
golang: {}
---
version: 1
language: go
golang: {}
`,
			wantErr: "only one YAML document",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Load(strings.NewReader(tt.input))
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Load() error = %v, want containing %q", err, tt.wantErr)
			}
		})
	}
}
