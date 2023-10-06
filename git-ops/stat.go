package gitops

import (
	"fmt"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

func Stat() {
	// Start the spinner in a goroutine
	done := make(chan bool)
	go utils.Spinner(done)
	// Get Status from the remote repository and current branch tracked
	fmt.Println("====================================")
	fmt.Println("           Repo Stats               ")
	fmt.Println("====================================")
	pullOutput, err := utils.RunGitCommand("remote", "show", "origin")
	if err != nil {
		fmt.Println("Error running 'git remote show origin':", err)
		return
	}
	// Stop the spinner
	done <- true
	fmt.Println("Repository and Branch details:")
	fmt.Println(pullOutput)

	done1 := make(chan bool)
	go utils.Spinner(done1)
	// Run "git status -s" and print the status
	statusOutput, err := utils.RunGitCommand("status", "--short")
	if err != nil {
		fmt.Println("Error running 'git status':", err)
		return
	}
	// Stop the spinner
	done1 <- true
	fmt.Println("====================================")
	fmt.Println("           File Changes              ")
	fmt.Println("====================================")
	fmt.Println(statusOutput)
	fmt.Println("====================================")
	fmt.Println("")

}
