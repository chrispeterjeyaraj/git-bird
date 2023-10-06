package gitops

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

// Push performs a Git Push operation.
func Push() {
	// Ask the user for input to choose what to commit
	fmt.Print("Would you like to see the changes before pushing? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	if input == "yes" {
		done1 := make(chan bool)
		go utils.Spinner(done1)
		// Run "git status -s" and print the status
		statusOutput, err := utils.RunGitCommand("status", "--short", "|", "column", "-t")
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

		ContinueOrReject()

	} else if input == "" || input == "no" {
		PushAction()
	}
}

func ContinueOrReject() {
	// Ask the user for input to choose what to commit
	fmt.Print("Looks good to continue? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	if input == "yes" {
		fmt.Println("Allright! Pushing changes to repository. Please Wait...")
		fmt.Println("")
		PushAction()
	} else if input == "" || input == "no" {
		return
	}
}

func PushAction() {
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
	fmt.Println("✓ Pushed changes to repository")
	fmt.Println(pullOutput)
}
