package repl

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/mkanenobu/google-ai-chat/ai"
)

func (repl *Repl) Evaluator(line string) bool {
	fmt.Printf("Evaluating: %s\n", line)

	ctx := context.Background()
	resp := ai.GenerateContent(ctx, ai.GenerateContentInput{
		Parts: []genai.Part{
			genai.Text(line),
		},
		ApiKey: repl.Config.ApiKey,
	})

	for _, candidate := range resp.Candidates {
		fmt.Printf("%s\n", candidate.Content.Parts)
	}

	repl.LocalHistory = append(repl.LocalHistory, resp)

	return true
}
