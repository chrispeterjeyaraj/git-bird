// gitops/commit.go
package gitops

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

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
	fmt.Println("====================================")
	fmt.Println("           File Changes              ")
	fmt.Println("====================================")
	fmt.Println(statusOutput)
	fmt.Println("====================================")
	fmt.Println("")

	// Ask the user for input to choose what to commit
	fmt.Print("Do you want to commit the entire root folder? (yes/no): ")
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
		fmt.Println("Added files for commit")
		fmt.Println("")
	} else if input == "no" {
		// Ask for the specific folder to commit
		fmt.Print("Enter the path to the specific folder to commit: ")
		scanner.Scan()
		folderPath := scanner.Text()

		// Commit the specific folder
		_, err = utils.RunGitCommand("add", folderPath)
		if err != nil {
			fmt.Println("Error running 'git add "+folderPath+"':", err)
			return
		}
		fmt.Println("Added files for commit")
		fmt.Println("")
	} else {
		fmt.Println("Invalid input. Please enter 'yes' or 'no'.")
		return
	}

	// Get user input for the commit message
	fmt.Print("Enter the commit message: ")
	scanner.Scan()
	commitMessage := scanner.Text()

	// Start the spinner in a goroutine
	done := make(chan bool)
	defer close(done)
	go spinner(done)

	// Commit with the provided message
	_, err = utils.RunGitCommand("commit", "-m", commitMessage)
	if err != nil {
		fmt.Println("Error running 'git commit':", err)
		return
	}

	// Push the changes
	pushOutput, err := utils.RunGitCommand("push")
	if err != nil {
		fmt.Println("Error running 'git push':", err)
		return
	}
	fmt.Println("Changes pushed to repository:")
	fmt.Println(pushOutput)
	// Wait for a moment before stopping the spinner
	time.Sleep(2 * time.Second)
}

func spinner(done chan bool) {
	// Define a set of spinner frames or characters
	frames := []string{"-", "\\", "|", "/"}

	i := 0
	for {
		select {
		case <-done:
			return
		default:
			// Print the current spinner frame or character
			fmt.Printf("\rPushing changes to repository. Please wait ........... \n%s", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
