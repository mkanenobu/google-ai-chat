package main

import (
	"github.com/mkanenobu/google-ai-chat/config"
	"github.com/mkanenobu/google-ai-chat/repl"
)

func main() {
	conf := config.NewConfig()
	rep := repl.NewRepl(conf)
	rep.StartRepl()
}
