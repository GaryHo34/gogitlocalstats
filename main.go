package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func printCommitDetail() {
	existing := read_setting_file()

	timeMap := generateTimeMap()

	for _, path := range existing {
		repo, err := git.PlainOpen(path)
		des := strings.Split(path, "/")
		fmt.Println("LOG: ", des[len(des)-1])
		if err != nil {
			panic(err)
		}

		ref, err := repo.Head()
		if err != nil {
			panic(err)
		}

		iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})
		if err != nil {
			panic(err)
		}

		err = iterator.ForEach(func(c *object.Commit) error {
			if c.Author.Email == "garyho0916@gmail.com" {
				timeMap[c.Author.When.Format("2006-01-02")] += 1
			}

			return nil
		})

		if err != nil {
			panic(err)
		}
	}

	today := time.Now()
	totalday := 26*7 + int(today.Weekday())
	commitMap := make([]string, 7)
	for i := 0; i < 7; i++ {
	}
	for i := totalday; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		day := date.Weekday()
		commitMap[int(day)] += fmt.Sprint(timeMap[date.Format("2006-01-02")]) + "\t"
	}
	for i := 0; i < 7; i++ {
		fmt.Println(commitMap[i])
	}
}

// count back for 25 weeks
func generateTimeMap() map[string]int {
	timeMap := make(map[string]int)
	today := time.Now()
	day := today.Weekday()

	totalday := 26*7 + int(day)
	for i := 0; i < totalday; i++ {
		date := today.AddDate(0, 0, -i)
		timeMap[date.Format("2006-01-02")] = 0
	}
	return timeMap
}

func stat(email string) {
	fmt.Println(DOT_FILE_PATH)
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@example.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		newGitPaths := scan_new_git_path(folder)

		fmt.Println("[scan] found", len(newGitPaths), "new git folders")

		// Print the results
		for _, path := range newGitPaths {
			fmt.Println("  ", path)
		}

		write_setting_file(newGitPaths)
	}

	printCommitDetail()
	stat(email)
}
