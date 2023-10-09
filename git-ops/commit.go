package gitops

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

// Commit performs a Git commit with the provided message.
func Commit(operation string) {
	// Run "git status -s" and print the status
	statusOutput, err := utils.RunGitBash("-c", "git status --short | column -t")
	if err != nil {
		fmt.Println("Error running 'git status':", err)
		return
	}
	fmt.Println("====================================")
	fmt.Println("           File Changes              ")
	fmt.Println("====================================")
	fmt.Println(statusOutput)
	fmt.Println("====================================")
	fmt.Println("")

	// Ask the user for input to choose what to commit
	fmt.Print("Do you want to commit all the changes? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	if input == "yes" || input == "" {
		// Commit the entire root folder
		_, err = utils.RunGitCommand("add", ".")
		if err != nil {
			fmt.Println("Error running 'git add .':", err)
			return
		}
		fmt.Println("")
		fmt.Println("✓ Added files for commit")
		fmt.Println("")
	} else if input == "no" {
		// Ask for the specific files or folders to commit
		fmt.Print("Enter the paths to the files or folders to commit (comma-separated): ")
		scanner.Scan()
		paths := strings.Split(scanner.Text(), ",")
		for _, path := range paths {
			// Trim leading and trailing spaces
			path = strings.TrimSpace(path)

			// Commit the specific file or folder
			_, err = utils.RunGitCommand("add", path)
			if err != nil {
				fmt.Printf("Error running 'git add %s': %v\n", path, err)
				return
			}
			fmt.Printf("✓ Added %s for commit\n", path)
		}
		fmt.Println("")
	} else {
		fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		return
	}

	// Get user input for the commit message
	fmt.Print("Enter the commit message: ")
	scanner.Scan()
	commitMessage := scanner.Text()

	// Commit with the provided message
	_, err = utils.RunGitCommand("commit", "-m", commitMessage)
	if err != nil {
		fmt.Println("Error running 'git commit':", err)
		return
	}

	if operation == "commitandpush" {
		// Start the spinner in a goroutine
		done := make(chan bool)
		go utils.Spinner(done)
		fmt.Println("")
		fmt.Println("\nPushing changes to repository...")
		fmt.Println("")

		// Push the changes
		pushOutput, err := utils.RunGitCommand("push")
		if err != nil {
			fmt.Println("Error running 'git push':", err)
			return
		}
		// Stop the spinner
		done <- true
		fmt.Printf("✓ Changes Pushed %s", pushOutput)
	}
}
