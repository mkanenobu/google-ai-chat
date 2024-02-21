package repl

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/mkanenobu/google-ai-chat/ai"
	"os"
)

func (repl *Repl) Generate(line string) {
	currentPart := genai.Text(line)

	parts := make([]genai.Part, len(repl.LocalHistory)+1)
	for _, content := range repl.LocalHistory {
		parts = append(parts, content.Parts...)
	}
	parts = append(parts, currentPart)

	for _, content := range repl.LocalHistory {
		parts = append(parts, content.Parts...)
	}

	ctx := context.Background()
	resp, err := ai.GenerateContent(ctx, ai.GenerateContentInput{
		Parts: []genai.Part{
			genai.Text(line),
		},
		ApiKey: repl.Config.ApiKey,
	})

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	content := resp.Candidates[0].Content
	for _, part := range content.Parts {
		fmt.Println(part)
	}

	repl.LocalHistory = append(repl.LocalHistory, &genai.Content{
		Parts: []genai.Part{currentPart},
		Role:  "user",
	})
	repl.LocalHistory = append(repl.LocalHistory, content)
}

func (repl *Repl) Evaluator(line string) {
	switch line {
	case ".exit":
		fmt.Println("Goodbye!")
		os.Exit(0)
	case ".history":
		for _, content := range repl.LocalHistory {
			fmt.Printf("%v\n", content)
		}
	case ".clear":
		repl.LocalHistory = make([]*genai.Content, 0)
	default:
		repl.Generate(line)
	}

}
