// gitops/commit.go
package gitops

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

// Commit performs a Git commit with the provided message.
func Commit() {
	// Run "git status -s" and print the status
	statusOutput, err := utils.RunGitCommand("status", "-s")
	if err != nil {
		fmt.Println("Error running 'git status':", err)
		return
	}
	fmt.Println("=========================")
	fmt.Println("		Changed files log: 		")
	fmt.Println("=========================")
	fmt.Println(statusOutput)
	fmt.Println("=========================")

	// Ask the user for input to choose what to commit
	fmt.Print("Do you want to commit the entire root folder? (yes/no): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.ToLower(scanner.Text())

	if input == "yes" {
		// Commit the entire root folder
		addOutput, err := utils.RunGitCommand("add", ".")
		if err != nil {
			fmt.Println("Error running 'git add .':", err)
			return
		}
		fmt.Println("Git Add:")
		fmt.Println(addOutput)
	} else if input == "no" {
		// Ask for the specific folder to commit
		fmt.Print("Enter the path to the specific folder to commit: ")
		scanner.Scan()
		folderPath := scanner.Text()

		// Commit the specific folder
		addOutput, err := utils.RunGitCommand("add", folderPath)
		if err != nil {
			fmt.Println("Error running 'git add "+folderPath+"':", err)
			return
		}
		fmt.Println("Git Add:")
		fmt.Println(addOutput)
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
	fmt.Println("Committed the files. Pushing changes to repository. Please wait ........")

	// Push the changes
	pushOutput, err := utils.RunGitCommand("push")
	if err != nil {
		fmt.Println("Error running 'git push':", err)
		return
	}
	fmt.Println("Changes pushed to repository:")
	fmt.Println(pushOutput)
}
