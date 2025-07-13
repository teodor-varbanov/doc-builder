package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func yamlExtractor(fileLocation string) string {

	var result []string

	file, err := os.ReadFile(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") || strings.Contains(scanner.Text(), "- name:") || strings.Contains(scanner.Text(), "when:") {
			result = append(result, scanner.Text())
		}

		if strings.Contains(scanner.Text(), "{{") {

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
	prettyResult := strings.Join(result[:], "\n")
	return prettyResult

}
