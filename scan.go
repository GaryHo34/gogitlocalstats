package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var DOT_FILE_PATH = ".visual-git"

func readSettingFile() []string {

	file, err := os.OpenFile(DOT_FILE_PATH, os.O_CREATE|os.O_RDONLY, 0755)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var dotGitPaths []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dotGitPaths = append(dotGitPaths, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal("[scan]", err)
		}
	}

	return dotGitPaths
}

func writeSettingFile(newGitPaths []string) {
	file, err := os.OpenFile(DOT_FILE_PATH, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var writeInPaths []string

	// path set
	set := make(map[string]bool, 0)

	for _, path := range readSettingFile() {
		set[path] = true
	}

	for _, path := range newGitPaths {
		_, ok := set[path]
		if !ok {
			writeInPaths = append(writeInPaths, path+"\n")
		}
	}

	fmt.Println("[scan] found", len(writeInPaths), "new git folders")

	content := strings.Join(writeInPaths, "")

	file.Write([]byte(content))
}

func scanNewGitPath(rootPath string) []string {

	var newGitPaths []string

	rootPath = strings.TrimSuffix(rootPath, "/")

	// Turn the path into absolute path
	absRootPath, err := filepath.Abs(rootPath)

	if err != nil {
		log.Fatal(err)
	}

	err = filepath.WalkDir(absRootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("[scan] prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if d.IsDir() && d.Name() == ".git" {
			fmt.Println("[scan] found .git path:", path)
			newGitPaths = append(newGitPaths, path)
			return nil
		}

		return nil
	})

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		log.Fatal(err)
	}

	return newGitPaths
}
