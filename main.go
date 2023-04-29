package main

import (
	"flag"
	"fmt"
)

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@example.com", "the email to scan")
	flag.StringVar(&email, "e", "your@example.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		newGitPaths := scanNewGitPath(folder)

		// Print the results
		for _, path := range newGitPaths {
			fmt.Println("  ", path)
		}

		writeSettingFile(newGitPaths)
	}

	if email != "" {
		commitTimeMap := loadGitCommit(email)
		printCommitDetail(commitTimeMap)
	}
}
