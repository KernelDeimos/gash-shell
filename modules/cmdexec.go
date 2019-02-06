package modules

import (
	"fmt"
	"github.com/KernelDeimos/gash-shell/console"
	"github.com/sirupsen/logrus"
	"os/exec"
)

type CommandExecutor_Sequential struct {
	Executors    []console.CommandExecutorI
	StopOnAction bool
	StopOnError  bool
}

func (ce CommandExecutor_Sequential) Executor(
	cmd string, args []interface{}, env console.Environment,
) (bool, error) {
	// TODO: should be error list instead of just last error
	var lastError error
	anyActionPerformed := false
	for _, exe := range ce.Executors {
		actionPerformed, err := exe(cmd, args, env)
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

func CommandExecutor_ExecOS(
	cmd string, args []interface{}, env console.Environment,
) (bool, error) {
	strargs := []string{}
	for _, arg := range args {
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
