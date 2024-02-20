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
	LocalHistory []*genai.GenerateContentResponse
}

func NewRepl(config config.Config) Repl {
	return Repl{config, make([]*genai.GenerateContentResponse, 0)}
}

func (repl *Repl) StartRepl() {
	replHistory := simplehistory.New()

	editor := &readline.Editor{
		PromptWriter: func(w io.Writer) (int, error) {
			return io.WriteString(w, "> ")
		},
		History:        replHistory,
		HistoryCycling: true,
	}

	fmt.Println("Type Ctrl-D or Ctrl-C to quit.")
	for {
		text, err := editor.ReadLine(context.Background())

		if err != nil {
			fmt.Printf("ERR=%s\n", err.Error())
			return
		}

		line := strings.TrimSpace(text)
		if len(line) <= 0 {
			continue
		}

		ch := make(chan bool)
		go (func() {
			ch <- repl.Evaluator(line)
		})()
		<-ch

		replHistory.Add(text)
	}
}
