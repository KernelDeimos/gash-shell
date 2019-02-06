package console

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
	string, []interface{}, error)
