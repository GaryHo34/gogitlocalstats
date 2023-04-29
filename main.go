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
		newGitPaths := scan_new_git_path(folder)

		// Print the results
		for _, path := range newGitPaths {
			fmt.Println("  ", path)
		}

		write_setting_file(newGitPaths)
	}

	if email != "" {
		commitTimeMap := load_get_commit(email)
		print_commit_detail(commitTimeMap)
	}
}
