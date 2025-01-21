// Set the GOOGLE_API_KEY env var to your API key taken from ai.google.dev
package main

import (
	"context"
	"fmt"
	"log"
	"os"
  "strings"
	"io/ioutil"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main.go <folder>")
	}
	dirName := os.Args[1]
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY env var not set")
	}

	ctx := context.Background()

	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	
	tplFile, err := os.Open("./jsonprompt.tpl")
	if err != nil {
		log.Fatal(err)
	}
	tpl,err := ioutil.ReadAll(tplFile)
	if err != nil {
		log.Fatal(err)
	}
	defer tplFile.Close()

	var sb strings.Builder

	for _, file := range files {
		filePath := dirName + "/" + file.Name()
		fmt.Println(filePath)

		jsonFile, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		json, err := ioutil.ReadAll(jsonFile)
    if err != nil {
			log.Fatal(err)
    }
    sb.WriteString(string(json))
    sb.WriteString("\n")
    defer jsonFile.Close()
  }
	prompt := strings.Replace(string(tpl), "{{jsons}}", sb.String(), 1)

  answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
