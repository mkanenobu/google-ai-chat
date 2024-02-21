package ai

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const DEFAULT_MODEL = "gemini-1.0-pro"

type GenerateContentInput struct {
	Parts  []genai.Part
	Model  *string
	ApiKey string
}

func GenerateContent(ctx context.Context, input GenerateContentInput) (*genai.GenerateContentResponse, error) {
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(input.ApiKey))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	m := DEFAULT_MODEL
	if input.Model != nil {
		m = *input.Model
	}
	model := client.GenerativeModel(m)

	return model.GenerateContent(ctx, input.Parts...)
}
