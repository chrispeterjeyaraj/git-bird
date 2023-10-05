package utils

import (
	"fmt"
	"os/exec"
	"time"
)

// RunGitCommand runs a Git command with the provided arguments and returns the output.
func RunGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func Spinner(done chan bool) {
	// Define a set of spinner frames or characters
	frames := []string{"⠈⠁", "⠈⠑", "⠈⠱", "⠈⡱", "⢀⡱", "⢄⡱", "⢄⡱", "⢆⡱", "⢎⡱", "⢎⡰", "⢎⡠", "⢎⡀", "⢎⠁", "⠎⠁", "⠊⠁"}

	i := 0
	for {
		select {
		case <-done:
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
