package main

import (
	"log"
	"os"
	"strings"
)

func yamlFinder(dir string) []string {
	files, err := os.ReadDir(dir)
	var yamlFiles []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".yaml") {
			yamlFiles = append(yamlFiles, file.Name())
		}
	}
	return yamlFiles
}
