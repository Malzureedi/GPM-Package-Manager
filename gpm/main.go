package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// RunCommand executes shell commands
func RunCommand(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:\n  gpm -i <package-repo-name>\n  gpm -f <repo-name>\n  gpm -r <package-dir>")
		return
	}

	command := os.Args[1]
	argument := os.Args[2]

	switch command {
	case "-i": // Install package
		fmt.Println("Installing:", argument)
		if err := RunCommand("git", "clone", argument); err != nil {
			fmt.Println("Error cloning repository:", err)
			return
		}

		// Extract repo name
		repoName := strings.TrimSuffix(argument[strings.LastIndex(argument, "/")+1:], ".git")
		repoPath := "./" + repoName

		// Check if gpmbuild.json exists
		buildFile := repoPath + "/gpmbuild.json"
		if _, err := os.Stat(buildFile); os.IsNotExist(err) {
			fmt.Println("No gpmbuild.json found in", repoName+". You will need to manually compile.")
		} else {
			// Automatically call autobuild executable
			fmt.Println("Running autobuild for", repoName)
			if err := RunCommand("./autobuild/autobuild", repoPath); err != nil {
				fmt.Println("Error running autobuild:", err)
			}
		}

	case "-f": // Find repo
		fmt.Println("Searching:", argument)
		RunCommand("git", "ls-remote", argument)

	case "-r": // Remove package
		fmt.Println("Uninstalling:", argument)
		if err := os.RemoveAll(argument); err != nil {
			fmt.Println("Error removing package:", err)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
