package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanFolderHelpler(path string, level int) []string {
	path = strings.TrimSuffix(path, "/")

	// get the file struct
	f, err := os.Open(path)

	// handle the err if any
	if err != nil {
		log.Fatal(err)
	}

	// File has a field name:
	files, err := f.ReadDir(0)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("found:")

	for _, file := range files {
		name := "dir"
		if !file.IsDir() {
			name = "file"
		}
		fmt.Printf("%10s: %s\n", name, file.Name())
	}

	return make([]string, 0)
}

func scanFolder(path string) []string {
	return scanFolderHelpler(path, 0)
}

func scan(path string) {
	// TODO: recursive scan for .git folders
	gitFolders := scanFolder(path)

	// TODO: Write path to the .visual-git file
	// writeInSetting(gitFolders)

	// Print the results
	fmt.Println("Found: ")
	for _, path := range gitFolders {
		fmt.Println(path)
	}
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
