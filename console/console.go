package console

import (
	"errors"
	"io"
	"os"

	"github.com/chzyer/readline"
	log "github.com/sirupsen/logrus"
)

type Console struct {
	// Interfaces
	LineReader      LineReaderI
	LineParser      LineParserI
	CommandExecutor CommandExecutorI
	PromptWriter    PromptWriterI

	// Not interfaces (will eventually become interfaces)
	Logger *log.Logger
}

func (c *Console) DoREPL() {
	ctx := ConsoleContext{
		Variables:              map[string]interface{}{},
		DefaultCommandExecutor: c.CommandExecutor,
	}

	for {
		// Default variables
		// TODO: refactor into an optional module
		{
			wd, _ := os.Getwd()
			uid := os.Getuid()
			ctx.Variables["cwd"] = wd
			ctx.Variables["super"] = uid == 0
		}

		promptText := c.PromptWriter(ctx)

		c.LineReader.SetPrompt(promptText)

		line, err := c.LineReader.Readline()

		// Handle error values returned by readline
		if err != nil {
			if err == io.EOF {
				// EOF: End of input file or Ctrl+D; quit the terminal
				c.Logger.Info("Goodbye")
				return
			} else if err == readline.ErrInterrupt {
				// Interupt: SIGINT or Ctrl+C; print instructions for quitting
				c.Logger.Info("Send EOF (Ctrl+D) to exit")
			} else {
				// Documentation for readline says there are no other error
				// values, so print "this should never happen" if that happens
				c.Logger.Error("[GASh/LineReader]", err)
			}
			continue
		}

		args, err := c.LineParser(line)
		if err != nil {
			c.Logger.Error(err)
		}

		c.Logger.WithFields(log.Fields{
			"args": args,
		}).Debug("command")

		if err != nil {
			c.Logger.Error("[GASh/LineParser]", err)
		}

		if len(args) < 1 {
			c.Logger.Debug("[GASh] Skipping blank input...")
			continue
		}

		cmdRun, cmdErr := c.CommandExecutor(args, Environment{
			Stderr:  os.Stderr,
			Stdout:  os.Stdout,
			Stdin:   os.Stdin,
			Context: ctx,
			Delegate: func([]interface{}, Environment) (bool, error) {
				return false, errors.New("[GASh] No command found")
			},
		})

		if cmdErr != nil {
			c.Logger.Error(cmdErr)
		}

		if cmdRun {
			c.Logger.Debug("[GASh] A command was run")
		}
	}

}
