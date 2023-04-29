package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func print_commit_detail(timeMap map[string]int) {
	today := time.Now()

	totalday := 26*7 + int(today.Weekday())

	commitMap := make([][]string, 7)

	for i := totalday; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		day := date.Weekday()
		commitMap[int(day)] = append(commitMap[int(day)], "\033[1;30;47m"+fmt.Sprint(timeMap[date.Format(time.DateOnly)])+"\033[0m")
	}

	for i := 0; i < 7; i++ {
		if i == 1 {
			fmt.Print("Mon  ")
		} else if i == 3 {
			fmt.Print("Wed  ")
		} else if i == 5 {
			fmt.Print("Fri  ")
		} else {
			fmt.Print("     ")
		}

		for _, cnt := range commitMap[i] {
			fmt.Printf("%-17s", cnt)
		}
		fmt.Print(" \n")
	}
}

func load_get_commit(email string) map[string]int {
	gitPaths := read_setting_file()

	timeMap := generate_timeMap()

	for _, path := range gitPaths {
		repo, err := git.PlainOpen(path)
		if err != nil {
			log.Fatal("[repo]: ", err)
		}

		repoName := strings.Split(path, "/")

		fmt.Println("[repo]: loading repo:", repoName[len(repoName)-2])

		ref, err := repo.Head()

		if err != nil {
			log.Fatal("[repo]: ", err)
		}

		iterator, err := repo.Log(&git.LogOptions{From: ref.Hash()})

		if err != nil {
			log.Fatal("[repo]: ", err)
		}

		err = iterator.ForEach(func(c *object.Commit) error {
			if strings.EqualFold(c.Author.Email, email) {
				timeMap[c.Author.When.Format(time.DateOnly)] += 1
			}

			return nil
		})

		if err != nil {
			log.Fatal("[repo]: ", err)
		}
	}
	return timeMap
}

func generate_timeMap() map[string]int {
	timeMap := make(map[string]int)

	today := time.Now()

	day := today.Weekday()

	// count back for 25 weeks
	totalday := 26*7 + int(day)

	for i := 0; i < totalday; i++ {
		date := today.AddDate(0, 0, -i)
		timeMap[date.Format(time.DateOnly)] = 0
	}

	return timeMap
}
