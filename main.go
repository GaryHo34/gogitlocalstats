package main

import (
	"flag"
	"fmt"
)

func scan(path string) {
	fmt.Println("scan:", path)
}

func stat(email string) {
	fmt.Println("stats:", email)
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@example.com", "the email to scan")
	flag.Parse()
	if folder != "" {
		scan(folder)
	}
	stat(email)
}
