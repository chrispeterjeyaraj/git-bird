package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chrispeterjeyaraj/git-bird/config"
	gitops "github.com/chrispeterjeyaraj/git-bird/git-ops"
)

func main() {
	// Extract the version details from config.yaml file
	config, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatal("Error loading configuration:", err)
	}
	// Define command-line flags
	var (
		commitFlag        = flag.Bool("commit", false, "Commit changes only")
		commitandpushFlag = flag.Bool("commitandpush", false, "Commit and push changes")
		pullFlag          = flag.Bool("pull", false, "Pull changes from remote repository")
		pushFlag          = flag.Bool("push", false, "Push changes to remote repository")
		statFlag          = flag.Bool("stat", false, "Show the repository, branch, and current status")
		helpFlag          = flag.Bool("help", false, "Show help documentation")
		usageFlag         = flag.Bool("usage", false, "Show help documentation")
		versionFlag       = flag.Bool("version", false, "Show version")
	)

	// Parse command-line flags
	flag.Parse()

	// Display the welcome message and usage information if no flags are provided
	if flag.NFlag() == 0 {
		fmt.Println("====================================")
		fmt.Printf("           GitBird %s          \n", config.Version)
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
		helpdocs(config.Version)
	}

	if *versionFlag {
		fmt.Printf("GitBird %s \n", config.Version)
	}
}

func helpdocs(version string) {
	fmt.Println("====================================")
	fmt.Printf("           GitBird %s          \n", version)
	fmt.Println("====================================")
	fmt.Println("Usage:")
	fmt.Println("  -stat: Show the repository, branch and current status")
	fmt.Println("  -commit: Commit changes")
	fmt.Println("  -commitandpush: Commit and push changes to repository")
	fmt.Println("  -pull: Pull changes from remote repository")
	fmt.Println("  -push: push changes from remote repository")
	fmt.Println("  -help: Show this help documentation")
	fmt.Println("  -version: Show this current stable version")
}
