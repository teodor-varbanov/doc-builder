package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ollama/ollama/api"
)

func aiHandler(yamlPrompt string) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	req := &api.GenerateRequest{
		Model:  "deepseek-r1:8b",
		Prompt: "Build detailed, stylized markdown documentation from the below yaml file, with a detailed summary at the top and then an overview of each task." + yamlPrompt,
		Stream: new(bool),
	}

	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		fmt.Println(resp.Response)

		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}
}
