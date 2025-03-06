package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Check if command arguments are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		fmt.Println("Example: go run main.go echo \"Hello $USER\"")
		os.Exit(1)
	}

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: Could not load .env file:", err)
	}

	// Join all arguments into a single command string
	commandArgs := os.Args[1:]
	command := strings.Join(commandArgs, " ")

	// Create the command to run in sh
	cmd := exec.Command("sh", "-c", command)

	// Pass through the environment variables
	cmd.Env = os.Environ()

	// Set up stdout and stderr to print to console
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
