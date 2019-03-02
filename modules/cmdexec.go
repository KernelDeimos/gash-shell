package modules

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"

	"github.com/KernelDeimos/gash-shell/console"
	"github.com/KernelDeimos/gash-shell/utilerr"
	"github.com/sirupsen/logrus"
)

// CommandExecutor_ChainMail chains a sequence of command executors together in
// two ways. First, it will run the command on each executor. If StopOnAction
// is true, the command will stop propagating in this first way as soon as an
// executor reports that some action was performed (i.e. the command was found).
// Second, it will set the Delegate() function in the environment input to each
// command executor such that Delegate() will pass data to the next executor in
// the chain. It makes like a triangle shape in my head, if that helps.
type CommandExecutor_ChainMail struct {
	Executors    []console.CommandExecutorI
	StopOnAction bool
	StopOnError  bool
}

func (ce CommandExecutor_ChainMail) Executor(
	args []interface{}, env console.Environment,
) (bool, error) {
	// TODO: should be error list instead of just last error
	var lastError error
	anyActionPerformed := false
	for i, exe := range ce.Executors {
		newEnv := env
		if i+1 < len(ce.Executors) {
			copyChainMail := ce
			copyChainMail.Executors = ce.Executors[i+1:]
			newEnv.Delegate = copyChainMail.Executor
		}

		actionPerformed, err := exe(args, newEnv)
		if err != nil {
			if ce.StopOnError {
				return actionPerformed, err
			}
			{
				logger := logrus.New()
				logger.SetOutput(env.Stdout)
				logger.Error(err)
				lastError = err
			}
		}
		if actionPerformed {
			anyActionPerformed = true
			if ce.StopOnAction {
				return true, err
			}
		}
	}
	return anyActionPerformed, lastError
}

func CommandExecutor_ExecPipe(
	args []interface{}, env console.Environment,
) (bool, error) {
	// TODO: report error for invalid type
	left := args[0].([]interface{})
	righ := args[1].([]interface{})

	eleft := env
	erigh := env

	var a *io.PipeReader
	var b *io.PipeWriter
	a, b = io.Pipe()

	eleft.Stdout = b
	eleft.Closer = b

	erigh.Stdin = a

	errs := utilerr.NewCompositeError(2)
	acts := []*bool{}
	for i := 0; i < 2; i++ {
		var a bool
		acts = append(acts, &a)
	}

	wg := &sync.WaitGroup{}
	start := func(
		args []interface{}, env console.Environment, err *error, act *bool,
	) {
		wg.Add(1)
		go func() {
			*act, *err = env.Delegate(args, env)
			env.Close()
			wg.Done()
		}()
	}

	start(left, eleft, errs.GetRef(0), acts[0])
	start(righ, erigh, errs.GetRef(1), acts[1])

	wg.Wait()

	var actionPerformed bool
	for _, ap := range acts {
		if *ap {
			actionPerformed = true
		}
	}

	return actionPerformed, errs.GetError()
}

func CommandExecutor_ExecOS(
	args []interface{}, env console.Environment,
) (bool, error) {
	if len(args) < 1 {
		return false, errors.New("ExecOS: blank input")
	}
	cmd := fmt.Sprint(args[0])

	strargs := []string{}
	for _, arg := range args[1:] {
		// TODO: Maybe add some kind of arg-stringer interface?
		strargs = append(strargs, fmt.Sprint(arg))
	}
	exeCmd := exec.Command(cmd, strargs...)
	exeCmd.Stdin = env.Stdin
	exeCmd.Stdout = env.Stdout
	exeCmd.Stderr = env.Stderr
	startErr := exeCmd.Start()
	if startErr != nil {
		return false, startErr
	}
	err := exeCmd.Wait()
	return true, err
}
