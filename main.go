package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
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

func getDotFilePath() string {

	dotFile := ".visual-git"
	return dotFile
}

func openReadSettingFile(filePath string) []string {
	fmt.Println(filePath)
	_, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			// other error
			panic(err)
		}
	}
	f, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}
	return lines
}
func writeInSetting(newPath []string) {
	filePath := getDotFilePath()
	existing := openReadSettingFile(filePath)

	for _, line := range newPath {
		flag := false
		for _, oldLine := range existing {
			flag = flag || (line == oldLine)
		}
		if flag {
			continue
		}
		existing = append(existing, line)
	}

	content := strings.Join(existing, "\n")
	os.WriteFile(filePath, []byte(content), 0755)
}

func printCommitDetail() {
	filePath := getDotFilePath()
	existing := openReadSettingFile(filePath)
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

func scan(path string) {
	// TODO: recursive scan for .git folders
	var gitFolders []string
	scanFolder(path, &gitFolders)
	// TODO: Write path to the .visual-git file
	writeInSetting(gitFolders)
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
	printCommitDetail()
	stat(email)
}
