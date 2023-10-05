package main

import (
	"flag"
	"fmt"

	gitops "github.com/chrispeterjeyaraj/git-bird/git-ops"
)

func main() {
	// Define command-line flags
	commitFlag := flag.Bool("commit", false, "Commit changes")
	commitandpushFlag := flag.Bool("commitandpush", false, "Commit and push changes")
	pullFlag := flag.Bool("pull", false, "Pull changes from remote repository")
	pushFlag := flag.Bool("push", false, "Push changes to remote repository")
	helpFlag := flag.Bool("help", false, "Show help documentation")

	// Parse command-line flags
	flag.Parse()

	// Show help documentation if requested
	if flag.NFlag() == 0 || *helpFlag {
		helpdocs()
		return
	}

	// Perform Git operations based on flags
	if *commitFlag {
		gitops.Commit("commit")
	}

	if *commitandpushFlag {
		gitops.Commit("commit&push")
	}

	if *pullFlag {
		gitops.Pull()
	}
	if *pushFlag {
		gitops.Push()
	}
}

func helpdocs() {
	fmt.Println("====================================")
	fmt.Println("           Git Bird v0.1            ")
	fmt.Println("====================================")
	fmt.Println("Usage:")
	fmt.Println("  -stat: Show the repository, branch and current status")
	fmt.Println("  -commit: Commit changes")
	fmt.Println("  -commitandpush: Commit and push changes to repository")
	fmt.Println("  -pull: Pull changes from remote repository")
	fmt.Println("  -push: push changes from remote repository")
	fmt.Println("  -help: Show this help documentation")
}
