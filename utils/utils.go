package utils

import (
	"fmt"
	"os/exec"
	"time"
)

// RunGitCommand executes a Git command with the provided arguments and returns
// the combined standard output and standard error as a string.
func RunGitCommand(args ...string) (string, error) {
	// Create a new command to run Git with the specified arguments
	cmd := exec.Command("git", args...)
	// Execute the Git command and capture its combined output
	output, err := cmd.CombinedOutput()
	// Convert the output to a string and return it along with any errors
	return string(output), err
}

func RunGitBash(args ...string) (string, error) {
	// Create a new command to run bash with the specified arguments
	cmd := exec.Command("bash", args...)
	// Execute the Git bash command and capture its combined output
	output, err := cmd.CombinedOutput()
	// Convert the output to a string and return it along with any errors
	return string(output), err
}

func Spinner(done chan bool) {
	// Define a set of spinner frames or characters
	// You can use any spiner characters you want from the SPINNER.md file
	frames := []string{"⠈⠁", "⠈⠑", "⠈⠱", "⠈⡱", "⢀⡱", "⢄⡱", "⢄⡱", "⢆⡱", "⢎⡱", "⢎⡰", "⢎⡠", "⢎⡀", "⢎⠁", "⠎⠁", "⠊⠁"}

	i := 0
	for {
		select {
		case <-done:
			// These Printf's are to clear any existing spinner before stopping the spinner.
			fmt.Printf("\033[2K")
			fmt.Println()
			fmt.Printf("\033[1A")
			return
		default:
			// Print the current spinner frame or character
			fmt.Printf("\r%s", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
