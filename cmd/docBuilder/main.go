package main

import (
	"flag"
	"fmt"
)

func main() {

	dir := flag.String("dir", ".", "Flag for the directory, where the yaml files are contained. Current dir is default")
	// yamlFile := flag.String("file", "./test-playbook/playbook.yaml", "File location of the yaml files.")
	flag.Parse()

	var files []string
	var docs []string
	files = yamlFinder(*dir)

	for i, file := range files {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!", file, "!!!!!!!!!!!!!!!!!!!!!")
		docs = append(docs, yamlExtractor(file))
		fmt.Println(docs[i])
		fmt.Println("Generating documentation. It may take a while. Please wait...")
		aiHandler(docs[i])
	}

}
