package utils

import (
	"os/exec"
)

// RunGitCommand runs a Git command with the provided arguments and returns the output.
func gitRunner(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
