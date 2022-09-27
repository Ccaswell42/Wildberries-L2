package pkg

import (
	"bytes"
	"io"
	"minishell/pkg/builtins"
	"os/exec"
	"strings"
)

func ParseCmd(cmdLine string) *builtins.CmdContext {
	cmdLineTrimSpace := strings.TrimSpace(cmdLine)
	arrCmdStrings := strings.Split(cmdLineTrimSpace, "|")
	var arrCmd []builtins.CmdContext
	var cmd builtins.CmdContext
	for _, val := range arrCmdStrings {
		cmdTrimSpace := strings.TrimSpace(val)
		cmdString := strings.Split(cmdTrimSpace, " ")
		cmd.Cmd = cmdString[0]
		cmd.Args = cmdString[1:]
		arrCmd = append(arrCmd, cmd)
	}
	for i := 0; i <= len(arrCmd)-2; i++ {
		arrCmd[i].Next = &arrCmd[i+1]
	}

	return &arrCmd[0]
}

func ExecuteCmd(cmd *builtins.CmdContext, stdIn io.Reader, stdOut io.Writer) error {

	cmd.StdIn = stdIn
	cmd.StdOut = stdOut
	buf := bytes.Buffer{}
	if cmd.Next != nil {
		cmd.StdOut = &buf
	}
	if FindBuiltins(*cmd) {
		err := ExecuteBuiltins(cmd)
		if err != nil {
			return err
		}
	} else {
		command := exec.Command(cmd.Cmd, cmd.Args...)
		command.Stdin = cmd.StdIn
		command.Stdout = cmd.StdOut
		err := command.Run()
		if err != nil {
			return err
		}
	}
	if cmd.Next != nil {
		return ExecuteCmd(cmd.Next, &buf, stdOut)
	}
	return nil
}
