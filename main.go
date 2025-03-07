package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Check if command arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		fmt.Println("Example: go run main.go echo \"Hello $USER\"")
		os.Exit(1)
	}

	// Join all arguments into a single command string
	commandArgs := os.Args[1:]
	command := strings.Join(commandArgs, " ")

	// Create the command to run in sh
	cmd := exec.Command("sh", "-c", command)

	// Pass through the environment variables
	before := os.Environ()
	after := before[:0]
	for _, env := range before {
		if k, v, found := strings.Cut(env, "="); found && strings.HasPrefix(v, "vs://") {
			if data, err := os.ReadFile(strings.TrimPrefix(v, "vs://")); err == nil {
				after = append(after, fmt.Sprintf("%s=%s", k, string(data)))
			}
		} else {
			after = append(after, env)
		}
	}
	cmd.Env = after

	// Set up stdout and stderr to print to console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
