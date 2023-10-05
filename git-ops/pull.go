// gitops/pull.go
package gitops

import (
	"fmt"
)

// Pull performs a Git pull operation.
func Pull() {
	// Pull changes from the remote repository
	pullOutput, err := gitbird.RunGitCommand("pull")
	if err != nil {
		fmt.Println("Error running 'git pull':", err)
		return
	}
	fmt.Println("Git Pull:")
	fmt.Println(pullOutput)
}
