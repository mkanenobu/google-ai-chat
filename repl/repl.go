package repl

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"github.com/mkanenobu/google-ai-chat/config"
	"github.com/nyaosorg/go-readline-ny"
	"github.com/nyaosorg/go-readline-ny/simplehistory"
	"io"
	"strings"
)

type Repl struct {
	Config       config.Config
	LocalHistory []*genai.Content
	Editor       *readline.Editor
}

func NewRepl(config config.Config) Repl {
	return Repl{Config: config, LocalHistory: make([]*genai.Content, 0)}
}

func (repl *Repl) StartRepl() {
	replHistory := simplehistory.New()

	editor := &readline.Editor{
		PromptWriter: func(w io.Writer) (int, error) {
			return io.WriteString(w, "> ")
		},
		History:        replHistory,
		HistoryCycling: false,
	}
	repl.Editor = editor

	fmt.Println("Type Ctrl-D or Ctrl-C to quit.")
	for {
		line, err := editor.ReadLine(context.Background())
		line = strings.TrimSpace(line)

		if err != nil {
			fmt.Printf("ERR=%s\n", err.Error())
			return
		}

		if len(line) <= 0 {
			continue
		}

		repl.Evaluator(line)

		replHistory.Add(line)
	}
}
