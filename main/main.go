package main

import (
	"github.com/chzyer/readline"
	log "github.com/sirupsen/logrus"

	"github.com/KernelDeimos/gash-shell/console"
	"github.com/KernelDeimos/gash-shell/modules"
)

const (
	Version = "v0.0.0"
)

func main() {

	// Create a line reader using Chzyer's readline (a port of GNU readline)
	lineReader, err := readline.NewEx(&readline.Config{
		Prompt:      "error setting prompt", // default prompt
		HistoryFile: "/tmp/gash_tmp_history",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Display the name of this shell and the version
	log.Info("GASh: Go Again Shell - " + Version)
	log.SetLevel(log.DebugLevel)

	cc := console.Console{
		LineReader: lineReader,
		LineParser: modules.LineParser_BasicStringsOnly,
		Logger:     log.StandardLogger(),
	}

	cc.DoREPL()

}

//
