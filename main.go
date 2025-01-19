// Set the GOOGLE_API_KEY env var to your API key taken from ai.google.dev
package main

import (
	"context"
	"fmt"
	"log"
	"os"
  "strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("GOOGLE_API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

  var sb strings.Builder
  for _, filename := range os.Args[1:] {
    fmt.Printf("Read %s file\n", filename)
    json, err := os.Open(filename)
    if err != nil {
      sb.WriteString("\n")
      sb.WriteString(json)
    }
    os.Close()
  }
  prompt := fmt.Sprintf("Generate a JSON Schema based on these bunch of JSON: %s\n", sb.String())
  
  answer, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
