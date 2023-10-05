package main

import (
	"flag"
	"fmt"
	"os"

	gitops "github.com/chrispeterjeyaraj/git-bird/git-ops"
)

func main() {
	// Define command-line flags
	commitFlag := flag.Bool("commit", false, "Commit changes")
	pullFlag := flag.Bool("pull", false, "Pull changes from remote repository")
	messageFlag := flag.String("m", "", "Commit message")

	// Parse command-line flags
	flag.Parse()

	// Perform Git operations based on flags
	if *commitFlag {
		if *messageFlag == "" {
			fmt.Println("Commit message is required when using --commit.")
			os.Exit(1)
		}
		fmt.Println("Performing commit operation...")
		gitops.Commit(*messageFlag)
	}

	if *pullFlag {
		fmt.Println("Performing pull operation...")
		gitops.Pull()
	}
}
