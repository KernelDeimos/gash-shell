package modules

import (
	"fmt"
	"github.com/KernelDeimos/gash-shell/console"
)

func PromptWriter_HardcodedBashLike(ctx console.ConsoleContext) string {
	promptChar := "$"
	if ctx.Is("super") {
		promptChar = "#"
	}

	return fmt.Sprintf("%s%s ",
		ctx.Getstring("cwd"),
		promptChar,
	)
}
