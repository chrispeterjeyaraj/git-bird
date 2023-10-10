package gitops

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chrispeterjeyaraj/git-bird/utils"
)

// Pull performs a Git Checkout operation.
func Checkout() {
	// Get user input for the checkout branch name
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the branch name you want to checkout from: ")
	scanner.Scan()
	checkoutBranch := scanner.Text()
	// Start the spinner in a goroutine
	done := make(chan bool)
	go utils.Spinner(done)
	// Pull changes from the remote repository
	checkoutOutput, err := utils.RunGitCommand("checkout", checkoutBranch)
	if err != nil {
		fmt.Println("Error running 'git checkout':", err)
		return
	}
	// Stop the spinner
	done <- true
	fmt.Printf("✓ Checkout branch %s from repository\n", checkoutBranch)
	fmt.Println(checkoutOutput)
}
