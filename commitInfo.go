package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var TRACK_BACK_WEEKS = 24

func printCommitDetail(timeMap map[string]int) {
	today := time.Now()

	totalday := TRACK_BACK_WEEKS*7 + int(today.Weekday())

	commitMap := make([][]string, 8)

	// default is white
	onFirstDateFlag := false
	for i := totalday; i >= 0; i-- {
		date := today.AddDate(0, 0, -i)
		weekday := int(date.Weekday())

		if date.Day() == 1 {
			onFirstDateFlag = true
		}

		if (weekday == 6 || i == 0) && onFirstDateFlag {
			commitMap[0] = append(commitMap[0], date.Format("Jan"))
			onFirstDateFlag = false
		} else if weekday == 6 {
			commitMap[0] = append(commitMap[0], "   ")
		}

		val := timeMap[date.Format(time.DateOnly)]

		escape := "\033[0;37;30m"
		reset := " \033[0m"
		padding := "  "
		text := fmt.Sprint(val)

		if i == 0 {
			padding = "   "
			escape = "\033[1;37;45m"
		} else if val == 0 {
			padding = "   "
			// escape color is black and text color is black
			escape = "\033[0;30;30m"
			text = " "
		} else if val > 0 && val < 5 {
			padding = "   "
			escape = "\033[1;30;47m"
		} else if val >= 5 && val < 10 {
			padding = "   "
			escape = "\033[1;30;43m"
		} else if val >= 10 {
			padding = "  "
			escape = "\033[1;30;42m"
		}
		commitMap[weekday+1] = append(commitMap[weekday+1], fmt.Sprintf("%s%s%s%s", escape, padding, text, reset))
	}
	fmt.Print("\n       ")
	for _, cnt := range commitMap[0] {
		fmt.Printf("%4s ", cnt)
	}
	fmt.Print(" \n")
	for i := 1; i < 8; i++ {
		if i == 2 {
			fmt.Print("Mon    ")
		} else if i == 4 {
			fmt.Print("Wed    ")
		} else if i == 6 {
			fmt.Print("Fri    ")
		} else {
			fmt.Print("       ")
		}

		for _, cnt := range commitMap[i] {
			fmt.Printf("%18s", cnt)
		}
		fmt.Print(" \n")
	}
	fmt.Print(" \n")
}

func loadGitCommit(email string) map[string]int {
	gitPaths := readSettingFile()

	timeMap := generateTimeMap()

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

func generateTimeMap() map[string]int {
	timeMap := make(map[string]int)

	today := time.Now()

	day := today.Weekday()

	// count back for 25 weeks
	totalday := TRACK_BACK_WEEKS*7 + int(day)

	for i := 0; i < totalday; i++ {
		date := today.AddDate(0, 0, -i)
		timeMap[date.Format(time.DateOnly)] = 0
	}

	return timeMap
}
