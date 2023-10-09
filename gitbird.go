package main

import (
	"flag"
	"fmt"

	gitops "github.com/chrispeterjeyaraj/git-bird/git-ops"
)

func main() {
	// Define command-line flags
	var (
		commitFlag        = flag.Bool("commit", false, "Commit changes")
		commitandpushFlag = flag.Bool("commitandpush", false, "Commit and push changes")
		pullFlag          = flag.Bool("pull", false, "Pull changes from remote repository")
		pushFlag          = flag.Bool("push", false, "Push changes to remote repository")
		statFlag          = flag.Bool("stat", false, "Show the repository, branch, and current status")
		helpFlag          = flag.Bool("help", false, "Show help documentation")
		usageFlag         = flag.Bool("usage", false, "Show help documentation")
	)

	// Parse command-line flags
	flag.Parse()

	// Display the welcome message and usage information if no flags are provided
	if flag.NFlag() == 0 {
		fmt.Println("====================================")
		fmt.Println("           Git Bird v0.1.1          ")
		fmt.Println("====================================")
		fmt.Println("Please use --usage or --help to know more about the features for Git Bird")
	}

	// Perform Git operations based on flags
	if *commitFlag || *commitandpushFlag {
		operation := "commit"
		if *commitandpushFlag {
			operation = "commitandpush"
		}
		gitops.Commit(operation)
	}

	if *pullFlag {
		gitops.Pull()
	}

	if *pushFlag {
		gitops.Push()
	}

	if *statFlag {
		gitops.Stat()
	}

	if *usageFlag || *helpFlag {
		helpdocs()
	}
}

func helpdocs() {
	fmt.Println("====================================")
	fmt.Println("           Git Bird v0.1.1          ")
	fmt.Println("====================================")
	fmt.Println("Usage:")
	fmt.Println("  -stat: Show the repository, branch and current status")
	fmt.Println("  -commit: Commit changes")
	fmt.Println("  -commitandpush: Commit and push changes to repository")
	fmt.Println("  -pull: Pull changes from remote repository")
	fmt.Println("  -push: push changes from remote repository")
	fmt.Println("  -help: Show this help documentation")
}
