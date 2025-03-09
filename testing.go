package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// An artificial input source.
	file, _ := os.ReadFile("./test-playbook/playbook.yaml")
	scanner := bufio.NewScanner(strings.NewReader(string(file)))
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#") || strings.Contains(scanner.Text(), "- name:") || strings.Contains(scanner.Text(), "when:") {
			fmt.Println(scanner.Text())
		}

		if strings.Contains(scanner.Text(), "{{") {

		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}

	yamlFinder(".")
}
