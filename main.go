package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func scanFolderHelpler(path string, folder *[]string) {
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

	var nextPath []string

	for _, file := range files {
		if file.IsDir() {
			if strings.EqualFold(file.Name(), ".git") {
				fmt.Println("[scan] found .git:", path)
				*folder = append(*folder, path)
				return
			}
			nextPath = append(nextPath, path+"/"+file.Name())
		}
	}

	for _, next := range nextPath {
		scanFolderHelpler(next, folder)
	}
}

func scanFolder(path string, folder *[]string) {
	scanFolderHelpler(path, folder)
}

func scan(path string) {
	// TODO: recursive scan for .git folders
	var gitFolders []string
	scanFolder(path, &gitFolders)

	// TODO: Write path to the .visual-git file
	// writeInSetting(gitFolders)

	// Print the results
	fmt.Println("[Result] Found:", len(gitFolders))
	for _, path := range gitFolders {
		fmt.Println("  ", path)
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
