// Package main implements the DevHatch command-line interface.
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		fmt.Println("Usage: devhatch init go")
		fmt.Println("Initialize a Go development environment")
		return
	}

	if args[0] == "init" {
		if len(args) != 2 {
			fmt.Fprintln(os.Stderr, "Usage: devhatch init go")
			os.Exit(1)
		}

		if args[1] != "go" {
			fmt.Fprintln(os.Stderr, "Only the Go environment is supported")
			os.Exit(1)
		}

		fmt.Fprintln(os.Stderr, "Go environment generation is not implemented yet")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Unknown command: %q\n", args[0])
	os.Exit(1)
}
