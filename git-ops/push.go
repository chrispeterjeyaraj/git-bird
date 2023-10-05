// gitops/pull.go
package gitops

import (
	"fmt"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

// Pull performs a Git pull operation.
func Push() {
	// Start the spinner in a goroutine
	done := make(chan bool)
	go utils.Spinner(done)
	// Pull changes from the remote repository
	pullOutput, err := utils.RunGitCommand("push")
	if err != nil {
		fmt.Println("Error running 'git pull':", err)
		return
	}
	// Stop the spinner
	done <- true
	fmt.Println("Pushed changes to repository")
	fmt.Println(pullOutput)
}
