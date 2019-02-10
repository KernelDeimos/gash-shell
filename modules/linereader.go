package modules

import (
	"strings"
)

// LineParser_BasicStringsOnly is a line parser that simply splits the input
// into tokens by spaces without parsing any specific data types. Commands
// which require numbers, arrays, and objects as arguments will be
// incompatible with this line parser.
func LineParser_BasicStringsOnly(input string) (
	args []interface{}, err error) {

	tokens := strings.Split(input, " ")
	if len(tokens) == 0 {
		//
		return []interface{}{}, nil
	}

	// In Go you can't cast []string to []interface{}, so we need to perform
	// the tedious task of copying the strings to the args list instead of
	// just doing `args = tokens[1:]`
	args = []interface{}{}
	for _, token := range tokens {
		args = append(args, token)
	}

	return args, nil
}
