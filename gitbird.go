package main

import (
	"flag"
	"fmt"

	gitops "github.com/chrispeterjeyaraj/git-bird/git-ops"
)

func main() {
	// Define command-line flags
	commitFlag := flag.Bool("commit", false, "Commit changes")
	pullFlag := flag.Bool("pull", false, "Pull changes from remote repository")

	// Parse command-line flags
	flag.Parse()

	// Perform Git operations based on flags
	if *commitFlag {
		fmt.Println("Performing commit operation...")
		gitops.Commit()
	}

	if *pullFlag {
		fmt.Println("Performing pull operation...")
		gitops.Pull()
	}
}
