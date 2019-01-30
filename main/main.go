package main

import (
	"github.com/chzyer/readline"
	log "github.com/sirupsen/logrus"
	"io"
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

	// TODO: Cooler prompt
	promptText := "> "

	for {
		// TODO: Prompt will be re-written here
		lineReader.SetPrompt(promptText)

		line, err := lineReader.Readline()

		// Handle error values returned by readline
		if err != nil {
			if err == io.EOF {
				// EOF: End of input file or Ctrl+D; quit the terminal
				log.Info("Goodbye")
				return
			} else if err == readline.ErrInterrupt {
				// Interupt: SIGINT or Ctrl+C; print instructions for quitting
				log.Info("Send EOF (Ctrl+D) to exit")
			} else {
				// Documentation for readline says there are no other error
				// values, so print "this should never happen" if that happens
				log.Error("this should never happen", err)
			}
			continue
		}

		// TODO: Everything
		log.Print("You typed " + line)
	}

}

//
