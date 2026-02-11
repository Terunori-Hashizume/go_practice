package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed rights
var rightsPath embed.FS

func main() {
	suffix := "_rights.txt"

	rightFiles, err := fs.ReadDir(rightsPath, "rights")
	if err != nil {
		fmt.Printf("Failed to read rights directory: %v\n", err)
		return
	}
	availableLanguages := make([]string, 0)
	fileMap := make(map[string]bool)
	for _, file := range rightFiles {
		name := file.Name()
		if len(name) > len(suffix) && name[len(name)-len(suffix):] == suffix {
			lang := name[:len(name)-len(suffix)]
			availableLanguages = append(availableLanguages, lang)
			fileMap[lang] = true
		}
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <language>")
		fmt.Printf("Available languages: %v\n", availableLanguages)
		return
	}

	language := os.Args[1]
	if !fileMap[language] {
		fmt.Printf("Alert: language '%s' not available.\n", language)
		fmt.Printf("Available languages: %v\n", availableLanguages)
		return
	}

	fileName := language + suffix
	data, err := fs.ReadFile(rightsPath, "rights/"+fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Print(string(data))
}
