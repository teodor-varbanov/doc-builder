package main

import (
	"flag"
	"fmt"
)

func main() {

	dir := flag.String("dir", ".", "Flag for the directory, where the yaml files are contained. Current dir is default")
	yamlFile := flag.String("file", "./test-playbook/playbook.yaml", "File location of the yaml files.")
	flag.Parse()

	var files []string
	files = yamlFinder(*dir)

	fmt.Println(files)
	yamlExtractor(*yamlFile)
}
