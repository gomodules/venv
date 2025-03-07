/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	// Check if arguments were provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command> [args...]")
		fmt.Println("Example: go run main.go ls -l")
		os.Exit(1)
	}

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

	// Get the command and its arguments
	command := os.Args[1]
	args := os.Args[2:]

	// Look up the full path of the command
	cmdPath, err := exec.LookPath(command)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: command '%s' not found\n", command)
		os.Exit(1)
	}

	// Use syscall.Exec to replace the current process
	err = syscall.Exec(cmdPath, append([]string{command}, args...), after)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}
