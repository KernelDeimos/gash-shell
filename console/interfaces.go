package console

import (
	"fmt"
	"io"
)

/*
	This file contains definitions of the things the console
	needs to function.

	For most of these, there's only one actualy implementation
	of the interface, but the interfaces will make the console
	very modular so that a user can choose between many options
	in the future.
*/

// LineReaderI defines methods which a valid line reader
// must implement. A line reader must be able to read a line
// of input (or produce an error), and recieve a string which
// it may optionally display as the prompt. An error produced
// by the line reader should match what is expected from the
// chzyer readline library.
type LineReaderI interface {
	Readline() (line string, err error)
	SetPrompt(newPrompt string)
}

// LineParserI defines a valid line parser function. The
// function must recieve a string representing some input
// (given by a line reader) and produce a command name and
// an array of interface{} values.
// The line parser may produce an error instead to indicate
// invalid syntax. The error returned should be a message
// understandable by the user.
type LineParserI func(input string) (
	[]interface{}, error)

type Environment struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	Context ConsoleContext

	// "Delegate" passes a command to the next executor in the chain
	Delegate CommandExecutorI
}

type CommandExecutorI func(
	// Command and arguments
	args []interface{},
	// Environment
	env Environment,
) (
	actionPerformed bool,
	err error,
)

type ConsoleContext struct {
	Variables map[string]interface{}

	// Note: DefaultCommandExecutor should never be used as a fallback executor!
	//       More accurately, an executor that calls DefaultCommandExecutor
	//       should be aware that it exists somewhere in the executor chain
	//       and prevent infinite loops accordingly.
	DefaultCommandExecutor CommandExecutorI
}

func (api ConsoleContext) Getstring(key string) string {
	value, _ := api.Variables[key]
	return fmt.Sprint(value)
}

func (api ConsoleContext) Is(key string) bool {
	value, _ := api.Variables[key].(bool)
	return value
}

type PromptWriterI func(
	ctx ConsoleContext,
) string
